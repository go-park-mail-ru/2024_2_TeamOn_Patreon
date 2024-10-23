package repositories

import (
	"errors"
	"time"

	// Модель репозитория взаимодействует с БД напрямую

	repModel "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/account/repository/models"
	busModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	models "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
)

var accountsInstance *Accounts

// Accounts реализует интерфейс AccountsRepository
type Accounts struct {
	accounts map[repModel.UserID]*repModel.Account
}

// New создает новый экземпляр Accounts.
func New() *Accounts {
	accountsInstance = &Accounts{
		accounts: make(map[repModel.UserID]*repModel.Account),
	}
	return accountsInstance
}

func Get() *Accounts {
	if accountsInstance == nil {
		accountsInstance = New()
	}
	return accountsInstance
}

// SaveAccount сохраняет профиль в базу данных
func (r *Accounts) SaveAccount(userID string, username string, role busModels.Role) (*repModel.Account, error) {
	// создание нового профиля

	account := repModel.Account{
		UserID:        repModel.UserID(userID),
		Username:      username,
		Email:         "",
		AvatarUrl:     "",
		Status:        "Создан, чтоб творить! Новая глава, новые вибрации!",
		Role:          RoleToString(role),
		Followers:     0,
		Subscriptions: 0,
		PostsAmount:   0,
		PostTitle:     "Здесь будет пост с монетизацией",
		PostContent:   "Здесь будет интересный контент",
		PostDate:      time.Now().Format("2006-01-02"),
	}

	// сохранение профиля в бд
	r.accounts[account.UserID] = &account

	return &account, nil
}

// UserExists проверяет, существует ли пользователь с указанным ID
func (r *Accounts) UserExist(userID string) (bool, error) {
	for _, account := range r.accounts {

		if account.UserID == repModel.UserID(userID) {
			return true, nil
		}
	}
	return false, nil
}

// GetAccountByID получает профиль по ID пользователя
func (r *Accounts) GetAccountByID(userID string) (*repModel.Account, error) {
	key := repModel.UserID(userID)
	foundAccount, ok := r.accounts[key]

	if !ok {
		return nil, errors.New("account not found")
	}

	if foundAccount == nil {
		return nil, errors.New("account is nil")
	}

	return foundAccount, nil
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
