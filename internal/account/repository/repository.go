package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	repModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/account/repository/models"

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

func (p *Postgres) UserByID(ctx context.Context, userID string) (*repModels.User, error) {
	op := "internal.account.repository.UserByID"

	// SQL-запрос для получения данных пользователя
	query := `
		SELECT 
			p.user_id, p.username, p.email, r.role_default_name
		FROM
			people p 
		INNER JOIN
			role r ON p.role_id = r.role_id 
		WHERE
			p.user_id = $1;
	`

	var user repModels.User
	err := p.db.QueryRow(ctx, query, userID).Scan(&user.UserID, &user.Username, &user.Email, &user.Role)

	if err != nil {
		if err == sql.ErrNoRows {
			// Если пользователь не найден, возвращаем nil без ошибки
			logger.StandardInfoF(
				"user with userID='%v' not found", userID,
				op)
			return nil, nil
		}
		return nil, errors.Wrap(err, op)
	}

	// Возвращаем данные пользователя
	return &user, nil
}

func (p *Postgres) SubscriptionsByID(ctx context.Context, userID string) ([]repModels.Subscription, error) {
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

	rows, err := p.db.Query(ctx, query, userID)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}
	defer rows.Close()

	logger.StandardDebugF(op, "wants to form an map of subscriptions for user with userID %v", userID)
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

func (p *Postgres) AvatarPathByID(ctx context.Context, userID string) (string, error) {
	op := "internal.account.repository.AvatarByID"

	query := `
		SELECT avatar_url 
		FROM avatar 
		WHERE user_id = $1
	`
	var avatarPath string
	err := p.db.QueryRow(ctx, query, userID).Scan(&avatarPath)
	if err != nil {
		return "", errors.Wrap(err, op)
	}

	return avatarPath, nil
}

func (p *Postgres) UpdateUsername(ctx context.Context, userID string, username string) error {
	op := "internal.account.repository.UpdateUsername"

	// SQL-запрос для обновления имени пользователя
	query := `
		UPDATE people 
		SET username = $1 
		WHERE user_id = $2;
	`

	// Выполняем запрос
	_, err := p.db.Exec(ctx, query, username, userID)
	if err != nil {
		errors.Wrap(err, op)
	}

	// Возвращаем nil, если обновление прошло успешно
	return nil
}

func (p *Postgres) UpdatePassword(ctx context.Context, userID string, hashPassword string) error {
	op := "internal.account.repository.UpdatePassword"

	// SQL-запрос для изменения пароля
	query := `
		UPDATE people 
		SET hash_password = $1 
		WHERE user_id = $2;
	`

	// Выполняем запрос
	_, err := p.db.Exec(ctx, query, hashPassword, userID)
	if err != nil {
		errors.Wrap(err, op)
	}

	// Возвращаем nil, если обновление прошло успешно
	return nil
}

func (p *Postgres) UpdateEmail(ctx context.Context, userID string, email string) error {
	op := "internal.account.repository.UpdateEmail"

	// SQL-запрос для изменения почты
	query := `
		UPDATE people 
		SET email = $1 
		WHERE user_id = $2;
	`

	// Выполняем запрос
	_, err := p.db.Exec(ctx, query, email, userID)
	if err != nil {
		errors.Wrap(err, op)
	}

	// Возвращаем nil, если обновление прошло успешно
	return nil
}

func (p *Postgres) DeleteAvatar(ctx context.Context, userID string) error {
	op := "internal.account.repository.DeleteAvatar"

	// Путь до старой аватарки
	avatarPath, err := p.AvatarPathByID(ctx, userID)
	if err != nil {
		logger.StandardInfo(
			fmt.Sprintf("old avatar doesn`t exist for user with userID %s", userID),
			op,
		)
		return nil
	}

	// Удаление записи о старой аватарке
	deleteQuery := `
		DELETE FROM avatar
		WHERE user_id = $1
	`

	logger.StandardDebugF(op, "want to delete record with old avatar")
	p.db.Exec(ctx, deleteQuery, userID)

	logger.StandardDebugF(op, "want to delete old avatar file")
	if err := os.Remove(avatarPath); err != nil {
		return errors.Wrap(err, op)
	}

	return nil
}

func (p *Postgres) UpdateAvatar(ctx context.Context, userID string, avatar multipart.File, fileName string) error {
	op := "internal.account.repository.UpdateAvatar"

	// Директория для сохранения аватаров
	avatarDir := "./static/avatar"

	// Получение формата загрузочного файла из его названия
	avatarFormat := filepath.Ext(fileName)

	// Формирование ID
	avatarID := p.GenerateID()

	// Полное имя сохраняемого файла
	fileFullName := avatarID + avatarFormat

	// Формируем путь к файлу из папки сохранения и названия файла
	avatarPath := filepath.Join(avatarDir, fileFullName)

	logger.StandardDebugF(op, "want to save new file with path %v", avatarPath)
	out, err := os.Create(avatarPath)
	if err != nil {
		return fmt.Errorf(op, err)
	}
	defer out.Close()

	// Сохраняем файл
	logger.StandardDebugF(op, "want to copy new avatar to path %v", avatarPath)
	if _, err := io.Copy(out, avatar); err != nil {
		return fmt.Errorf(op, err)
	}

	// Запрос на создание новой записи о новой аватарке
	query := `
		INSERT INTO avatar (avatar_id, user_id, avatar_url)
		VALUES ($1, $2, $3)
		ON CONFLICT (avatar_id) DO UPDATE 
		SET user_id = EXCLUDED.user_id, avatar_url = EXCLUDED.avatar_url
	`
	// Выполняем запрос
	if _, err := p.db.Exec(ctx, query, avatarID, userID, avatarPath); err != nil {
		return errors.Wrap(err, op)
	}

	// Возвращаем nil, если обновление прошло успешно
	return nil
}

func (p *Postgres) GenerateID() string {
	id, _ := uuid.NewV4()

	return id.String()
}

// UpdateRole меняет роль пользователя на "author"
func (p *Postgres) UpdateRole(ctx context.Context, userID string) error {
	op := "internal.account.repository.UpdateRole"

	query := `
		UPDATE people
		SET role_id = (SELECT role_id FROM role WHERE role_default_name = 'Author')
		WHERE user_id = $1
	`

	// Выполняем запрос
	if _, err := p.db.Exec(ctx, query, userID); err != nil {
		errors.Wrap(err, op)
	}

	logger.StandardInfo(
		fmt.Sprintf("successful change role for userID: %s", userID),
		op,
	)

	// Возвращаем nil, если изменение прошло успешно
	return nil
}

func (p *Postgres) InitPage(ctx context.Context, userID string) error {
	op := "internal.account.repository.InitPage"

	pageID := p.GenerateID()

	query := `
		INSERT INTO page (page_id, user_id, info, background_picture_url)
		VALUES ($1, $2, NULL, NULL);
	`

	if _, err := p.db.Exec(ctx, query, pageID, userID); err != nil {
		errors.Wrap(err, op)
	}

	// Возвращаем nil, если создание прошло успешно
	return nil
}
