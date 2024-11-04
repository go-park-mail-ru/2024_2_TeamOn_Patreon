package repositories

import (
	"context"
	"database/sql"
	"fmt"

	sModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/author/service/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Поле структуры - pool соединений с БД
type Postgres struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Postgres {
	return &Postgres{db: db}
}

func (p *Postgres) AuthorByID(ctx context.Context, authorID string) (*sModels.Author, error) {
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

	var author sModels.Author

	if err := p.db.QueryRow(ctx, query, authorID).Scan(&author.Username, &author.Info); err != nil {
		if err == sql.ErrNoRows {
			// Если автор не найден, возвращаем nil без ошибки
			logger.StandardInfoF(
				"author with authorID='%v' not found", authorID,
				op)
			return nil, nil
		}
		logger.StandardDebugF(op, "get author error: {%v}", err)
		return nil, err
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
		logger.StandardDebugF(op, "get followers error: {%v}", err)
		return nil, err
	}

	// Возвращаем данные автора
	return &author, nil
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
		logger.StandardDebugF(op, "update info error: {%v}", err)
		return err
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
				"author with authorID='%v' not found", authorID,
				op)
			return 0, nil
		}
		logger.StandardDebugF(op, "get payments error: {%v}", err)
		return 0, err
	}

	return amountPayments, nil
}

func (p *Postgres) BackgroundPathByID(ctx context.Context, userID string) (string, error) {
	op := "internal.account.repository.BackgroundPathByID"

	query := `
		SELECT background_picture_url 
		FROM page 
		WHERE user_id = $1
	`
	var backgroundPath string
	err := p.db.QueryRow(ctx, query, userID).Scan(&backgroundPath)
	if err != nil {
		logger.StandardDebugF(op, "get background error: {%v}", err)
		return "", err
	}

	return backgroundPath, nil
}

func (p *Postgres) UpdateBackground(ctx context.Context, userID string, backgroundPath string) error {
	op := "internal.account.repository.UpdateBackground"

	// Запрос на изменение графы "фон" для пользователя
	query := `
		UPDATE page 
		SET background_picture_url = $1 
		WHERE user_id = $2;
	`
	// Выполняем запрос
	if _, err := p.db.Exec(ctx, query, backgroundPath, userID); err != nil {
		logger.StandardDebugF(op, "update avatar error: {%v}", err)
		return err
	}

	logger.StandardInfo(
		fmt.Sprintf("successful update avatar for userID: %s", userID),
		op,
	)
	// Возвращаем nil, если обновление прошло успешно
	return nil
}

func (p *Postgres) GenerateID() string {
	id, _ := uuid.NewV4()

	return id.String()
}
