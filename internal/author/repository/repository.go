package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	repModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/author/repository/models"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

// Поле структуры - pool соединений с БД
type Postgres struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Postgres {
	return &Postgres{db: db}
}

func (p *Postgres) AuthorByID(ctx context.Context, authorID string) (*repModels.Author, error) {
	op := "internal.author.repository.AuthorByID"

	// SQL-запрос для получения username, info
	query := `
		SELECT 
			p.username,	pg.info
		FROM 
			page pg
		JOIN 
			people p ON pg.user_id = p.user_id
		WHERE 
			pg.user_id = $1;
	`

	var author repModels.Author

	if err := p.db.QueryRow(ctx, query, authorID).Scan(&author.Username, &author.Info); err != nil {
		if err == sql.ErrNoRows {
			// Если автор не найден, возвращаем nil без ошибки
			logger.StandardInfoF(
				ctx,
				"author with authorID='%v' not found", authorID,
				op)
			return nil, nil
		}
		return nil, errors.Wrap(err, op)
	}

	// SQL-запрос для получения followers
	queryFollowers := `
		SELECT 
			COUNT(DISTINCT s.user_id) AS followers
		FROM 
			subscription s
		JOIN 
			custom_subscription cs ON s.custom_subscription_id = cs.custom_subscription_id
		WHERE 
			cs.author_id = $1;
	`

	if err := p.db.QueryRow(ctx, queryFollowers, authorID).Scan(&author.Followers); err != nil {
		if err == sql.ErrNoRows {
			// Если подписчики не найдены, проставляем ноль
			author.Followers = 0
		}
		logger.StandardDebugF(ctx, op, "get followers error: {%v}", err)
		return nil, errors.Wrap(err, op)
	}

	// Возвращаем данные автора
	return &author, nil
}

func (p *Postgres) UserIsSubscribe(ctx context.Context, authorID, userID string) (bool, error) {
	op := "internal.account.repository.UserIsSubscribe"
	logger.StandardDebugF(ctx, op, "wants to check relations userID=%v and authorID=%v", userID, authorID)

	query := `
		SELECT 
			EXISTS (
				SELECT 1
				FROM subscription s
				JOIN custom_subscription cs ON s.custom_subscription_id = cs.custom_subscription_id
				WHERE s.user_id = $1 AND cs.author_id = $2
			) AS is_subscribed;
	`
	var subscribeStatus bool
	p.db.QueryRow(ctx, query, userID, authorID).Scan(&subscribeStatus)

	return subscribeStatus, nil
}

func (p *Postgres) SubscriptionsByID(ctx context.Context, authorID string) ([]repModels.Subscription, error) {
	op := "internal.account.repository.SubscriptionsByID"

	// SQL-запрос для получения данных о подписках
	query := `
		SELECT 
			cs.author_id, p.username
		FROM
			subscription s
		INNER JOIN
			custom_subscription cs ON s.custom_subscription_id = cs.custom_subscription_id
		INNER JOIN
			people p ON cs.author_id = p.user_id
		WHERE
			s.user_id = $1;
	`

	rows, err := p.db.Query(ctx, query, authorID)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}
	defer rows.Close()

	logger.StandardDebugF(ctx, op, "wants to form an map of subscriptions for author with authorID %v", authorID)
	var subscriptions []repModels.Subscription
	for rows.Next() {
		var subscription repModels.Subscription
		if err := rows.Scan(&subscription.AuthorID, &subscription.AuthorName); err != nil {
			return nil, errors.Wrap(err, op)
		}
		subscriptions = append(subscriptions, subscription)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, op)
	}

	// Возвращаем данные о подписках
	return subscriptions, nil
}

func (p *Postgres) UpdateInfo(ctx context.Context, authorID string, info string) error {
	op := "internal.account.repository.UpdateInfo"

	// SQL-запрос для обновления инфо
	query := `
		UPDATE page 
		SET info = $1 
		WHERE user_id = $2;
	`

	// Выполняем запрос
	_, err := p.db.Exec(ctx, query, info, authorID)
	if err != nil {
		return errors.Wrap(err, op)
	}

	// Возвращаем nil, если обновление прошло успешно
	return nil
}

func (p *Postgres) Payments(ctx context.Context, authorID string) (int, error) {
	op := "internal.author.repository.Payments"

	// SQL-запрос для получения payments за донаты и подписки
	query := `
		SELECT 
			COALESCE(SUM(t.cost), 0) + COALESCE(SUM(cs.cost), 0) AS total_payments
		FROM 
			tip t
		FULL OUTER JOIN 
			subscription s ON t.author_id = s.custom_subscription_id
		FULL OUTER JOIN 
			custom_subscription cs ON s.custom_subscription_id = cs.custom_subscription_id
		WHERE 
			t.author_id = $1 OR cs.author_id = $1;
	`

	var amountPayments int

	if err := p.db.QueryRow(ctx, query, authorID).Scan(&amountPayments); err != nil {
		if err == sql.ErrNoRows {
			// Если автор не найден, возвращаем 0 без ошибки
			logger.StandardInfoF(
				ctx,
				"payments fpr authorID='%v' not found", authorID,
				op)
			return 0, nil
		}
		return 0, errors.Wrap(err, op)
	}

	return amountPayments, nil
}

func (p *Postgres) BackgroundPathByID(ctx context.Context, authorID string) (string, error) {
	op := "internal.account.repository.BackgroundPathByID"

	query := `
		SELECT background_picture_url 
		FROM page 
		WHERE user_id = $1
	`
	var backgroundPath string
	err := p.db.QueryRow(ctx, query, authorID).Scan(&backgroundPath)
	if err != nil {
		logger.StandardInfo(
			ctx,
			fmt.Sprintf("background doesn`t exist for author with authorID %s", authorID),
			op,
		)
		return "", errors.Wrap(err, op)
	}

	return backgroundPath, nil
}

func (p *Postgres) DeleteBackground(ctx context.Context, authorID string) error {
	op := "internal.account.repository.DeleteBackground"

	// Получаем путь до старого фона
	logger.StandardDebugF(ctx, op, "want to get path to old background for authorID %v", authorID)
	oldBackgroundPath, err := p.BackgroundPathByID(ctx, authorID)

	if err != nil {
		logger.StandardInfo(
			ctx,
			fmt.Sprintf("old background doesn`t exist for author with authorID %s", authorID),
			op,
		)
		return nil
	}

	logger.StandardDebugF(ctx, op, "want to delete old background with path %v", oldBackgroundPath)
	if err := os.Remove(oldBackgroundPath); err != nil {
		return errors.Wrap(err, op)
	}

	return nil
}

func (p *Postgres) UpdateBackground(ctx context.Context, authorID string, background multipart.File, fileName string) error {
	op := "internal.account.repository.UpdateBackground"

	// Директория для сохранения фона
	backgroundDir := "./static/background"

	// Получение формата загрузочного файла из его названия
	backgroundFormat := filepath.Ext(fileName)

	// Формирование ID
	backgroundID := p.GenerateID()

	// Полное имя сохраняемого файла
	fileFullName := backgroundID + backgroundFormat

	// Формируем путь к файлу из папки сохранения и названия файла
	backgroundPath := filepath.Join(backgroundDir, fileFullName)

	logger.StandardDebugF(ctx, op, "want to save new file with path %v", backgroundPath)
	out, err := os.Create(backgroundPath)
	if err != nil {
		return fmt.Errorf(op, err)
	}
	defer out.Close()

	// Сохраняем файл
	logger.StandardDebugF(ctx, op, "want to copy new background to path %v", backgroundPath)
	if _, err := io.Copy(out, background); err != nil {
		return fmt.Errorf(op, err)
	}

	// Обновляем информацию в БД
	logger.StandardDebugF(ctx, op, "want to save background URL %v in DB", backgroundPath)

	// Запрос на изменение графы "фон" для автора
	query := `
		UPDATE page 
		SET background_picture_url = $1 
		WHERE user_id = $2;
	`
	// Выполняем запрос
	if _, err := p.db.Exec(ctx, query, backgroundPath, authorID); err != nil {
		return errors.Wrap(err, op)
	}

	logger.StandardInfo(
		ctx,
		fmt.Sprintf("successful update record for authorID: %s", authorID),
		op,
	)
	// Возвращаем nil, если обновление прошло успешно
	return nil
}

func (p *Postgres) NewTip(ctx context.Context, userID, authorID string, cost int, message string) error {
	op := "internal.account.repository.NewTip"

	// Запрос на добавление записи Tip
	query := `
		INSERT INTO 
			tip (tip_id, user_id, author_id, cost, message, payed_date)
        VALUES 
			($1, $2, $3, $4, $5, $6)
	`

	tipID := p.GenerateID()
	// Выполняем запрос
	if _, err := p.db.Exec(ctx, query, tipID, userID, authorID, cost, message, time.Now()); err != nil {
		return errors.Wrap(err, op)
	}

	logger.StandardInfo(
		ctx,
		fmt.Sprintf("successful create new record for authorID: %s", authorID),
		op,
	)
	// Возвращаем nil, если обновление прошло успешно
	return nil
}

func (p *Postgres) GenerateID() string {
	id, _ := uuid.NewV4()

	return id.String()
}
