package repositories

import (
	"context"
	"errors"
	imModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/repository/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	bModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	"github.com/gofrs/uuid"
	errors2 "github.com/pkg/errors"
	"sync"
)

// ImaginaryRepository реализует интерфейс AuthRepository
type ImaginaryRepository struct {
	users  map[bModels.UserID]*imModels.User
	mu     sync.RWMutex
	lastID int
}

// New создает новый экземпляр ImaginaryRepository.
func New() *ImaginaryRepository {
	return &ImaginaryRepository{
		users:  make(map[bModels.UserID]*imModels.User),
		lastID: 1,
	}
}

// SaveUserWithRole сохраняет пользователя
func (r *ImaginaryRepository) SaveUserWithRole(ctx context.Context, userId uuid.UUID, username, role, hash string) error {
	// создание нового пользователя
	user := imModels.User{
		UserID:       bModels.UserID(userId.String()),
		Username:     username,
		Role:         role,
		PasswordHash: hash,
	}

	r.mu.Lock()
	// сохранение пользователя в бд
	r.users[user.UserID] = &user
	r.mu.Unlock()

	return nil
}

// UserExists проверяет, существует ли пользователь с указанным именем.
func (r *ImaginaryRepository) UserExists(ctx context.Context, username string) (bool, error) {
	r.mu.RLock()
	for _, user := range r.users {

		if user.Username == username {
			return true, nil
		}
	}
	r.mu.RUnlock()
	return false, nil
}

func (r *ImaginaryRepository) GetUserByUserId(ctx context.Context, userId uuid.UUID) (username, role string, err error) {
	r.mu.RLock()
	imUser := r.users[bModels.UserID(userId.String())]
	r.mu.RUnlock()

	if imUser == nil {
		return "", "", global.ErrUserNotFound
	}

	return imUser.Username, imUser.Role, nil
}

// GetPasswordHashByID получает хэш пароля пользователя по его ID.
func (r *ImaginaryRepository) GetPasswordHashByID(ctx context.Context, userID uuid.UUID) (string, error) {
	r.mu.RLock()
	imUser := r.users[bModels.UserID(userID.String())]
	r.mu.RUnlock()

	if imUser == nil {
		return "", errors.New("user not found")
	}
	return imUser.PasswordHash, nil
}

func (r *ImaginaryRepository) generateID() bModels.UserID {
	r.lastID++
	return bModels.UserID(r.lastID)
}

// GetUserByUsername возвращает пользователя по имени.
// Output: userId, role, error
func (r *ImaginaryRepository) GetUserByUsername(ctx context.Context, username string) (uuid.UUID, string, error) {
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
		return uuid.Nil, "", errors2.Wrap(global.ErrUserNotFound, "")
	}

	userIdUuid, err := uuid.FromString(string(imUser.UserID))
	if err != nil {
		return uuid.Nil, "", global.ErrUserNotFound
	}
	return userIdUuid, imUser.Role, nil
}
