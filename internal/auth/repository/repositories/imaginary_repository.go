package repositories

import (
	"errors"
	bSession "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/behavior/session"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/behavior/uuid"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/config"
	repModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/repository/models"
	bModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/common/business/models"
	"sync"
)

// ImaginaryRepository реализует интерфейс AuthRepository
type ImaginaryRepository struct {
	users map[repModels.UserID]*repModels.User
	mu    sync.RWMutex

	sessions map[repModels.UserID]*repModels.Session
}

// New создает новый экземпляр ImaginaryRepository.
func New() *ImaginaryRepository {
	return &ImaginaryRepository{
		users:    make(map[repModels.UserID]*repModels.User),
		sessions: make(map[repModels.UserID]*repModels.Session),
	}
}

// SaveUser сохраняет пользователя в базу данных.
func (r *ImaginaryRepository) SaveUser(username string, role int, passwordHash string) (*bModels.User, error) {
	// создание нового пользователя
	user := repModels.User{
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
	bUser := repModels.MapImUserToBUser(user)
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
func (r *ImaginaryRepository) GetUserByID(userID string) (*bModels.User, error) {
	r.mu.RLock()
	imUser := r.users[repModels.UserID(userID)]
	r.mu.RUnlock()

	if imUser == nil {
		return nil, config.ErrUserNotFound
	}

	bUser := repModels.MapImUserToBUser(*imUser)
	return &bUser, nil
}

// GetPasswordHashByID получает хэш пароля пользователя по его ID.
func (r *ImaginaryRepository) GetPasswordHashByID(userID string) (string, error) {
	r.mu.RLock()
	imUser := r.users[repModels.UserID(userID)]
	r.mu.RUnlock()

	if imUser == nil {
		return "", errors.New("user not found")
	}
	return imUser.PasswordHash, nil
}

func (r *ImaginaryRepository) generateID() repModels.UserID {
	newID := uuid.GenerateUUID()
	return repModels.UserID(newID)
}

// GetUserByUsername возвращает пользователя по имени.
func (r *ImaginaryRepository) GetUserByUsername(username string) (*bModels.User, error) {
	var imUser *repModels.User

	r.mu.RLock()
	for _, user := range r.users {

		if user.Username == username {
			imUser = user
			break
		}
	}
	r.mu.RUnlock()

	if imUser == nil {
		return nil, config.ErrUserNotFound
	}

	bUser := repModels.MapImUserToBUser(*imUser)
	return &bUser, nil
}

func (r *ImaginaryRepository) RemoveUserByID(userID string) error {
	r.mu.Lock()
	delete(r.users, repModels.UserID(userID))
	r.mu.Unlock()
	return nil
}

func (r *ImaginaryRepository) SaveSession(session bSession.SessionModel) (string, error) {
	repSession := repModels.MapBehaviorSessionToRepositorySession(session)
	r.mu.Lock()
	r.sessions[repSession.UserID] = repSession
	r.mu.Unlock()

	return session.SessionID, nil
}

func (r *ImaginaryRepository) DeleteSession(session bSession.SessionModel) error {
	repSession := repModels.MapBehaviorSessionToRepositorySession(session)
	r.mu.Lock()
	delete(r.sessions, repSession.UserID)
	r.mu.Unlock()
	return nil
}

func (r *ImaginaryRepository) CheckSession(session bSession.SessionModel) (bool, error) {
	repSession := repModels.MapBehaviorSessionToRepositorySession(session)
	r.mu.RLock()
	_, ok := r.sessions[repSession.UserID]
	r.mu.RUnlock()
	if !ok {
		return false, nil
	}
	return true, nil
}
