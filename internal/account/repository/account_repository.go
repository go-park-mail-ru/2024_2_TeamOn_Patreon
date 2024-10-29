package repositories

import (
	"context"
	"errors"
	"fmt"
	"log"

	// Модель репозитория взаимодействует с БД напрямую
	bModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/account/service/models"
	b2Models "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	models "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	pgx "github.com/jackc/pgx/v4"
)

// Поле структуры - соединение с БД
type Postgres struct {
	db *pgx.Conn
}

func New(db *pgx.Conn) *Postgres {
	return &Postgres{db: db}
}

// SaveAccount сохраняет профиль в базу данных
func (p *Postgres) SaveAccount(ctx context.Context, userID string, username string, role b2Models.Role) (*bModels.Account, error) {
	// создание нового профиля
	account := bModels.Account{
		UserID:        bModels.UserID(userID),
		Username:      username,
		Email:         "", // Если нужно, можно передать email как параметр
		Role:          RoleToString(role),
		Subscriptions: nil,
	}

	// SQL-запрос для вставки нового пользователя в таблицу users
	query := `
		INSERT INTO users (user_id, username, email, role_id)
		VALUES ($1, $2, $3, (SELECT role_id FROM roles WHERE default_role_name = $4))
		ON CONFLICT (user_id) DO UPDATE SET
			username = EXCLUDED.username,
			email = EXCLUDED.email,
			role_id = (SELECT role_id FROM roles WHERE default_role_name = EXCLUDED.role_id)
	`

	// Выполнение SQL-запроса
	_, err := p.db.Exec(ctx, query, userID, username, account.Email, account.Role)
	if err != nil {
		log.Println("Ошибка при выполнении запроса:", err)
		return nil, err
	}

	// Возвращаем сохраненный аккаунт
	return &account, nil
}

// AccountExist проверяет, существует ли пользователь с указанным ID
func (p *Postgres) AccountExist(ctx context.Context, userID string) (bool, error) {
	// SQL-запрос для проверки существования аккаунта
	query := `
		SELECT COUNT(*) 
		FROM users 
		WHERE user_id = $1
	`

	var count int
	err := p.db.QueryRow(ctx, query, userID).Scan(&count)

	if err != nil {
		fmt.Println("Ошибка при выполнении запроса:", err)
		return false, err
	}

	// Если count больше 0, аккаунт существует
	return count > 0, nil
}

// FindByID получает профиль по ID пользователя
func (p *Postgres) FindByID(ctx context.Context, userID string) (*bModels.Account, error) {
	// Проверка на пустой userID
	if userID == "" {
		return nil, errors.New("userID не может быть пустым")
	}

	// SQL-запрос для получения данных пользователя
	query := `
		SELECT u.user_id, u.username, u.email, p.default_role_name
		FROM users u
		JOIN roles p ON u.role_id = p.role_id
		WHERE u.user_id = $1
	`

	var account bModels.Account
	err := p.db.QueryRow(ctx, query, userID).Scan(
		&account.UserID,
		&account.Username,
		&account.Email,
		&account.Role,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("пользователь не найден")
		}
		log.Println("Ошибка при выполнении запроса:", err)
		return nil, err
	}

	// Возвращаем найденный аккаунт
	return &account, nil
}

// RoleToString
// Функция для отображения роли в виде строки
func RoleToString(role models.Role) string {
	switch role {
	case models.Reader:
		return "Reader"
	case models.Author:
		return "Author"
	default:
		return "Unknown"
	}
}
