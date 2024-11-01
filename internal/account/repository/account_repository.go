package repositories

import (
	"context"
	"database/sql"
	"fmt"

	// Модель репозитория взаимодействует с БД напрямую

	sModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/account/service/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/gofrs/uuid"
	pgx "github.com/jackc/pgx/v4"
)

// Поле структуры - соединение с БД
type Postgres struct {
	db *pgx.Conn
}

func New(db *pgx.Conn) *Postgres {
	return &Postgres{db: db}
}

// UserByID получает данные пользователя по указанному ID
func (p *Postgres) UserByID(ctx context.Context, userID string) (*sModels.User, error) {
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

	// rep модель пользователя, в дальнейшем конвертировать в service модель аккаунта
	var user sModels.User
	err := p.db.QueryRow(ctx, query, userID).Scan(&user.UserID, &user.Username, &user.Email, &user.Role)

	if err != nil {
		if err == sql.ErrNoRows {
			// Если пользователь не найден, возвращаем nil без ошибки
			logger.StandardInfoF(
				"user with userID='%v' not found", userID,
				op)
			return nil, nil
		}
		logger.StandardDebugF(op, "get user error: {%v}", err)
		return nil, err
	}

	// Возвращаем данные пользователя
	return &user, nil
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
		logger.StandardDebugF(op, "get avatar error: {%v}", err)
		return "", err
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
		logger.StandardDebugF(op, "update username error: {%v}", err)
		return err
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
		logger.StandardDebugF(op, "update password error: {%v}", err)
		return err
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
		logger.StandardDebugF(op, "update email error: {%v}", err)
		return err
	}

	// Возвращаем nil, если обновление прошло успешно
	return nil
}

func (p *Postgres) UpdateAvatar(ctx context.Context, userID string, avatarID string, avatarPath string) error {
	op := "internal.account.repository.UpdateAvatar"

	// Удаление записи о старой аватарке (если она есть)
	deleteQuery := `
		DELETE FROM avatar
		WHERE user_id = $1
	`

	if _, err := p.db.Exec(ctx, deleteQuery, userID); err != nil {
		logger.StandardDebugF(op, "delete old avatar error: {%v}", err)
		return err
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
