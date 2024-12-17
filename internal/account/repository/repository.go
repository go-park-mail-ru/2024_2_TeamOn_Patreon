package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"path/filepath"

	repModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/account/repository/models"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/static"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
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
			// Если пользователь не найден
			logger.StandardInfoF(
				ctx,
				"user with userID='%v' not found", userID,
				op)
			return nil, errors.Wrap(global.ErrUserNotFound, op)
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

	logger.StandardDebugF(ctx, op, "wants to form an map of subscriptions for user with userID %v", userID)
	var subscriptions []repModels.Subscription

	// Если строки есть, обрабатываем их
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
		if errors.Is(err, sql.ErrNoRows) {
			// Если пользователь не загрузил аватарку, возвращаем путь к дефолтной аватарке
			logger.StandardInfo(
				ctx,
				fmt.Sprintf("user %v has not uploaded an avatar yet", userID),
				op,
			)
			return p.getPathForDefaultAvatar(), nil
		}
		logger.StandardDebugF(ctx, op, "error querying avatar {%v}", err)
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

func (p *Postgres) GetPasswordHashByID(ctx context.Context, userID string) (string, error) {
	op := "internal.account.repository.GetPasswordHashByID"

	query := `
		SELECT hash_password
		FROM people
		WHERE user_id = $1
	`
	rows, err := p.db.Query(ctx, query, userID)
	if err != nil {
		return "", errors.Wrap(err, op)
	}

	defer rows.Close()

	var (
		hash string
	)

	for rows.Next() {
		if err = rows.Scan(&hash); err != nil {
			return "", errors.Wrap(err, op)
		}
		logger.StandardDebugF(ctx, op, "GetPasswordHashByID found hash: %s", hash)
		return hash, nil
	}
	return "", errors.Wrap(global.ErrUserNotFound, op)
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

// Сейчас удаляется только запись в БД. Файл не удаляется (на РК2 разрешили пока не удалять)
func (p *Postgres) DeleteAvatar(ctx context.Context, userID string) error {
	op := "internal.account.repository.DeleteAvatar"

	// Путь до старой аватарки
	oldAvatarPath, err := p.AvatarPathByID(ctx, userID)

	// Если старой аватарки не было (пользователь еще не загрузил)
	if oldAvatarPath == p.getPathForDefaultAvatar() {
		logger.StandardInfo(
			ctx,
			fmt.Sprintf("nothing to delete: user %v has not uploaded an avatar yet", userID),
			op,
		)
		return nil
	}
	if err != nil {
		return errors.Wrap(err, op)
	}

	// Запрос на удаление записи о старой аватарке
	deleteQuery := `
		DELETE FROM avatar
		WHERE user_id = $1
	`

	logger.StandardDebugF(ctx, op, "want to delete record with old avatar")

	if _, err := p.db.Exec(ctx, deleteQuery, userID); err != nil {
		return errors.Wrap(err, op)
	}

	// Удаление файла из файловой системы
	// logger.StandardDebugF(ctx, op, "want to delete old avatar file")
	// if err := os.Remove(avatarPath); err != nil {
	// 	return errors.Wrap(err, op)
	// }

	return nil
}

func (p *Postgres) UpdateAvatar(ctx context.Context, userID string, file []byte, fileExtension string) error {
	op := "internal.account.repository.UpdateAvatar"

	// Директория для сохранения аватаров
	fileDir := "./static/avatar"

	// Формирование ID
	fileID := utils.GenerateUUID()

	// Формируем путь к файлу
	filePath := static.CreateFilePath(fileDir, fileID, fileExtension)

	// Сохраняем файл в хранилище
	logger.StandardDebugF(ctx, op, "want to save new file with path %v", filePath)
	err := static.SaveFile(file, filePath)
	if err != nil {
		return errors.Wrap(err, op)
	}

	// Запрос на создание новой записи о новой аватарке
	query := `
		INSERT INTO avatar (avatar_id, user_id, avatar_url)
		VALUES ($1, $2, $3)
		ON CONFLICT (avatar_id) DO UPDATE 
		SET user_id = EXCLUDED.user_id, avatar_url = EXCLUDED.avatar_url
	`
	// Выполняем запрос
	if _, err := p.db.Exec(ctx, query, fileID, userID, filePath); err != nil {
		return errors.Wrap(err, op)
	}

	// Возвращаем nil, если обновление прошло успешно
	return nil
}

func (p *Postgres) GenerateID() string {
	id := uuid.NewV4()

	return id.String()
}

func (p *Postgres) IsReader(ctx context.Context, userID string) (bool, error) {
	op := "internal.account.repository.IsReader"

	query := `
		SELECT r.role_default_name
		FROM people p
		JOIN role r ON p.role_id = r.role_id
		WHERE p.user_id = $1;
	`

	var roleName string
	err := p.db.QueryRow(ctx, query, userID).Scan(&roleName)
	if err != nil {
		return false, errors.Wrap(err, op)
	}

	return roleName == "Reader", nil

}

func (p *Postgres) UpdateRoleToAuthor(ctx context.Context, userID string) error {
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
		ctx,
		fmt.Sprintf("successful change role for userID: %s", userID),
		op,
	)

	// Возвращаем nil, если изменение прошло успешно
	return nil
}

func (p *Postgres) InitPage(ctx context.Context, userID string) error {
	op := "internal.account.repository.InitPage"

	pageID := p.GenerateID()
	backgroundURL := p.getPathForDefaultBackground()

	query := `
		INSERT INTO page (page_id, user_id, info, background_picture_url)
		VALUES ($1, $2, NULL, $3);
	`

	if _, err := p.db.Exec(ctx, query, pageID, userID, backgroundURL); err != nil {
		errors.Wrap(err, op)
	}

	// Возвращаем nil, если создание прошло успешно
	return nil
}

func (p *Postgres) getPathForDefaultAvatar() string {
	defaultDir := "./static/avatar"
	defaultName := "default.jpg"
	defaultPath := filepath.Join(defaultDir, defaultName)
	return defaultPath
}

func (p *Postgres) getPathForDefaultBackground() string {
	defaultDir := "./static/background"
	defaultName := "default.jpg"
	defaultPath := filepath.Join(defaultDir, defaultName)
	return defaultPath
}
