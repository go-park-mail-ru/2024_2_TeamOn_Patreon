package repositories

import (
	"errors"
	// Модель репозитория взаимодействует с БД напрямую

	busModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/common/buisness/models"
	repModel "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/profile/repository/models"
)

// Profiles реализует интерфейс ProfilesRepository
type Profiles struct {
	profiles map[repModel.UserID]*repModel.Profile
}

// New создает новый экземпляр Profiles.
func New() *Profiles {
	return &Profiles{
		profiles: make(map[repModel.UserID]*repModel.Profile),
	}
}

// SaveProfile сохраняет профиль в базу данных
func (r *Profiles) SaveProfile(userID int, username string, role busModels.Role) (*repModel.Profile, error) {
	// создание нового профиля

	profile := repModel.Profile{
		UserID:        repModel.UserID(userID),
		Username:      username,
		Email:         "",
		AvatarUrl:     "",
		Role:          busModels.RoleToString(role),
		Followers:     0,
		Subscriptions: 0,
	}

	// сохранение профиля в бд
	r.profiles[profile.UserID] = &profile

	return &profile, nil
}

// UserExists проверяет, существует ли пользователь с указанным ID
func (r *Profiles) UserExists(userID int) (bool, error) {
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
	foundedProfile := r.profiles[key]

	if foundedProfile == nil {
		return nil, errors.New("failed to get user")
	}

	return foundedProfile, nil
}
