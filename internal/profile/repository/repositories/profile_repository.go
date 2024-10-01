package repositories

import (
	"errors"
	// Модель репозитория взаимодействует с БД напрямую

	busModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/common/buisness/models"
	models "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/common/buisness/models"
	repModel "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/profile/repository/models"
)

var profilesInstance *Profiles

// Profiles реализует интерфейс ProfilesRepository
type Profiles struct {
	profiles map[repModel.UserID]*repModel.Profile
}

// New создает новый экземпляр Profiles.
func New() *Profiles {
	profilesInstance = &Profiles{
		profiles: make(map[repModel.UserID]*repModel.Profile),
	}
	return profilesInstance
}

func Get() *Profiles {
	if profilesInstance == nil {
		profilesInstance = New()
	}
	return profilesInstance
}

// SaveProfile сохраняет профиль в базу данных
func (r *Profiles) SaveProfile(userID int, username string, role busModels.Role) (*repModel.Profile, error) {
	// создание нового профиля

	profile := repModel.Profile{
		UserID:        repModel.UserID(userID),
		Username:      username,
		Email:         "",
		AvatarUrl:     "",
		Status:        "Создан, чтоб творить! Новая глава, новые вибрации!",
		Role:          RoleToString(role),
		Followers:     0,
		Subscriptions: 0,
		PostsAmount:   0,
	}

	// сохранение профиля в бд
	r.profiles[profile.UserID] = &profile

	return &profile, nil
}

// UserExists проверяет, существует ли пользователь с указанным ID
func (r *Profiles) UserExist(userID int) (bool, error) {
	for _, profile := range r.profiles {

		if profile.UserID == repModel.UserID(userID) {
			return true, nil
		}
	}
	return false, nil
}

// GetProfileByID получает профиль по ID пользователя
func (r *Profiles) GetProfileByID(userID int) (*repModel.Profile, error) {
	key := repModel.UserID(userID)
	foundProfile, ok := r.profiles[key]

	if !ok {
		return nil, errors.New("profile not found")
	}

	if foundProfile == nil {
		return nil, errors.New("profile is nil")
	}

	return foundProfile, nil
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
