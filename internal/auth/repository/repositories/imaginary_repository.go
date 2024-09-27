package repositories

import (
	"errors"
	imModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/tree/polina-auth/internal/auth/repository/models"
	bModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/tree/polina-auth/internal/common/buisness/models"
)

// ImaginaryRepository реализует интерфейс AuthRepository
type ImaginaryRepository struct {
	users  map[imModels.UserID]*imModels.User
	lastID imModels.UserID
}

// New создает новый экземпляр ImaginaryRepository.
func New() *ImaginaryRepository {
	return &ImaginaryRepository{
		users:  make(map[imModels.UserID]*imModels.User),
		lastID: 1,
	}
}

// SaveUser сохраняет пользователя в базу данных.
func (r *ImaginaryRepository) SaveUser(username string, role int, passwordHash string) (*bModels.User, error) {
	// создание нового пользователя
	user := imModels.User{
		UserID:       r.generateID(),
		Username:     username,
		Role:         role,
		PasswordHash: passwordHash,
	}

	// сохранение пользователя в бд
	r.users[user.UserID] = &user

	// мапим в бизнес модель user
	bUser := imModels.MapImUserToBUser(user)
	return &bUser, nil
}

// UserExists проверяет, существует ли пользователь с указанным именем.
func (r *ImaginaryRepository) UserExists(username string) (bool, error) {
	for _, user := range r.users {

		if user.Username == username {
			return true, nil
		}
	}
	return false, nil
}

// GetUserByID получает пользователя по его ID.
func (r *ImaginaryRepository) GetUserByID(userID int) (*bModels.User, error) {
	imUser := r.users[imModels.UserID(userID)]

	if imUser == nil {
		return nil, errors.New("user not found")
	}

	bUser := imModels.MapImUserToBUser(*imUser)
	return &bUser, nil
}

// GetPasswordHashByID получает хэш пароля пользователя по его ID.
func (r *ImaginaryRepository) GetPasswordHashByID(userID int) (string, error) {
	imUser := r.users[imModels.UserID(userID)]

	if imUser == nil {
		return "", errors.New("user not found")
	}
	return imUser.PasswordHash, nil
}

func (r *ImaginaryRepository) generateID() imModels.UserID {
	r.lastID++
	return r.lastID
}

// GetUserByUsername возвращает пользователя по имени.
func (r *ImaginaryRepository) GetUserByUsername(username string) (*bModels.User, error) {
	var imUser *imModels.User
	for _, user := range r.users {

		if user.Username == username {
			imUser = user
			break
		}
	}
	if imUser == nil {
		return nil, errors.New("user not found")
	}

	bUser := imModels.MapImUserToBUser(*imUser)
	return &bUser, nil
}
