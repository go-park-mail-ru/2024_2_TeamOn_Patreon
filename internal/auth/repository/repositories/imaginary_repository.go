package repositories

import (
	"errors"
	imModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/repository/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	bModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	"github.com/gofrs/uuid"
	"sync"
)

// ImaginaryRepository реализует интерфейс AuthRepository
type ImaginaryRepository struct {
	users map[bModels.UserID]*imModels.User
	mu    sync.RWMutex
}

// New создает новый экземпляр ImaginaryRepository.
func New() *ImaginaryRepository {
	return &ImaginaryRepository{
		users: make(map[bModels.UserID]*imModels.User),
	}
}

// SaveUser сохраняет пользователя в базу данных.
func (r *ImaginaryRepository) SaveUser(username string, role string, passwordHash string) (*bModels.User, error) {
	// создание нового пользователя
	user := imModels.User{
		UserID:       r.generateID(),
		Username:     username,
		Role:         role,
		PasswordHash: passwordHash,
	}

	r.mu.Lock()
	// сохранение пользователя в бд
	r.users[user.UserID] = &user
	r.mu.Unlock()

	// мапим в бизнес модель user
	bUser := imModels.MapImUserToBUser(user)
	return &bUser, nil
}

// UserExists проверяет, существует ли пользователь с указанным именем.
func (r *ImaginaryRepository) UserExists(username string) (bool, error) {
	r.mu.RLock()
	for _, user := range r.users {

		if user.Username == username {
			return true, nil
		}
	}
	r.mu.RUnlock()
	return false, nil
}

// GetUserByID получает пользователя по его ID.
func (r *ImaginaryRepository) GetUserByID(userID bModels.UserID) (*bModels.User, error) {
	r.mu.RLock()
	imUser := r.users[bModels.UserID(userID)]
	r.mu.RUnlock()

	if imUser == nil {
		return nil, global.ErrUserNotFound
	}

	bUser := imModels.MapImUserToBUser(*imUser)
	return &bUser, nil
}

// GetPasswordHashByID получает хэш пароля пользователя по его ID.
func (r *ImaginaryRepository) GetPasswordHashByID(userID bModels.UserID) (string, error) {
	r.mu.RLock()
	imUser := r.users[userID]
	r.mu.RUnlock()

	if imUser == nil {
		return "", errors.New("user not found")
	}
	return imUser.PasswordHash, nil
}

func (r *ImaginaryRepository) generateID() bModels.UserID {
	id, _ := uuid.NewV4()

	return bModels.UserID(id.String())
}

// GetUserByUsername возвращает пользователя по имени.
func (r *ImaginaryRepository) GetUserByUsername(username string) (*bModels.User, error) {
	var imUser *imModels.User

	r.mu.RLock()
	for _, user := range r.users {

		if user.Username == username {
			imUser = user
			break
		}
	}
	r.mu.RUnlock()

	if imUser == nil {
		return nil, global.ErrUserNotFound
	}

	bUser := imModels.MapImUserToBUser(*imUser)
	return &bUser, nil
}
