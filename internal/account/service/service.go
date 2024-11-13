// Бизнес-логика сервиса Account

package service

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"

	interfaces "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/account/service/interfaces"
	sModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/account/service/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	logger "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	rep interfaces.AccountRepository
}

func New(repository interfaces.AccountRepository) *Service {
	return &Service{repository}
}

// GetAccDataByID - получение данных аккаунта по userID
func (s *Service) GetAccDataByID(ctx context.Context, userID string) (sModels.User, error) {
	op := "internal.account.service.GetAccDataByID"

	// получаем данные пользователя из rep
	logger.StandardDebugF(ctx, op, "want to get user data for userID = %v", userID)

	user, err := s.rep.UserByID(ctx, userID)
	if err != nil {
		return sModels.User{}, errors.Wrap(err, op)
	}

	logger.StandardInfo(
		ctx,
		fmt.Sprintf("successful get user=%v, role=%v, email=%v", user.Username, user.Role, user.Email),
		op)

	accountData := sModels.User{
		Username: user.Username,
		Email:    user.Email.String,
		Role:     user.Role,
	}
	return accountData, nil
}

// GetAccSubscriptions - получение подписок аккаунта по userID
func (s *Service) GetAccSubscriptions(ctx context.Context, userID string) ([]sModels.Subscription, error) {
	op := "internal.account.service.GetAccSubscriptions"

	// получаем подписки пользователя из rep
	logger.StandardDebugF(ctx, op, "want to get subscriptions for user with userID %v", userID)
	repSubscriptions, err := s.rep.SubscriptionsByID(ctx, userID)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	// Преобразование подписок из репозитория в сервисные модели
	subscriptions := make([]sModels.Subscription, len(repSubscriptions))
	for i, repSub := range repSubscriptions {
		subscriptions[i] = sModels.Subscription{
			AuthorID:   repSub.AuthorID,
			AuthorName: repSub.AuthorName,
		}
	}

	logger.StandardInfo(
		ctx,
		fmt.Sprintf("successful get subscriptions: %v", subscriptions),
		op)
	return subscriptions, nil
}

// GetAvatarByID - получение аватарки пользователя по userID
func (s *Service) GetAvatarByID(ctx context.Context, userID string) ([]byte, error) {
	op := "internal.account.service.GetAvatarByID"

	// ОБращаемся в репозиторий для получения пути до аватара
	logger.StandardDebugF(ctx, op, "want to find avatarPATH for userID = %v", userID)

	avatarPath, err := s.rep.AvatarPathByID(ctx, userID)
	if err != nil {
		// Если не получилось найти путь аватара -> 404
		return nil, errors.Wrap(err, op)
	}

	// По найденному пути открываем файл аватара
	logger.StandardDebugF(ctx, op, "want to read file with path: %v", avatarPath)
	avatar, err := os.ReadFile(avatarPath)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	logger.StandardInfo(
		ctx,
		fmt.Sprintf("successful get avatar file with path %v for user with userID %v", avatarPath, userID),
		op)

	return avatar, nil
}

// PostUpdateAvatar - изменение аватарки аккаунта по userID
func (s *Service) PostUpdateAvatar(ctx context.Context, userID string, avatarFile multipart.File, fileName string) error {
	op := "internal.account.service.PostAccountUpdateAvatar"

	// Удаляем старый аватар, если он есть
	logger.StandardDebugF(ctx, op, "want to delete old avatar file")
	if err := s.rep.DeleteAvatar(ctx, userID); err != nil {
		return errors.Wrap(err, op)
	}

	logger.StandardInfo(
		ctx,
		fmt.Sprintf("successful delete old avatar for userID %v", userID),
		op,
	)

	// Сохраняем новый
	logger.StandardDebugF(ctx, op, "want to save new avatar file")
	if err := s.rep.UpdateAvatar(ctx, userID, avatarFile, fileName); err != nil {
		return errors.Wrap(err, op)
	}

	logger.StandardInfo(
		ctx,
		fmt.Sprintf("successful save new avatar for userID %v", userID),
		op,
	)
	return nil
}

// PostUpdateRole - изменение роли пользователя на "автор"
func (s *Service) PostUpdateRole(ctx context.Context, userID string) error {
	op := "internal.account.service.PostUpdateRole"

	// Проверяем, что пользователь еще не является "author"
	logger.StandardDebugF(ctx, op, "want to check user role")
	ok, err := s.rep.IsReader(ctx, userID)
	if err != nil {
		return errors.Wrap(err, op)
	} else if !ok {
		logger.StandardDebugF(ctx, op, "user with userID=%v already is 'Author'", userID)
		return global.ErrRoleAlreadyChanged
	}

	// Обновляем поле роль с "reader" на "author"
	if err := s.updateRole(ctx, userID); err != nil {
		return errors.Wrap(err, op)
	}

	// Заполняем сущность page для нового автора
	if err := s.initPage(ctx, userID); err != nil {
		return errors.Wrap(err, op)
	}

	logger.StandardInfo(
		ctx,
		fmt.Sprintf("successful change role for userID: %v", userID),
		op)

	return nil
}

func (s *Service) UpdateUsername(ctx context.Context, userID string, username string) error {
	op := "internal.account.service.updateUsername"

	if err := s.rep.UpdateUsername(ctx, userID, username); err != nil {
		return errors.Wrap(err, op)
	}
	logger.StandardInfo(
		ctx,
		fmt.Sprintf("successful update username: %v", username),
		op)
	return nil
}

func (s *Service) UpdatePassword(ctx context.Context, userID string, password string) error {
	op := "internal.account.service.updatePassword"

	// Хеширование пароля с заданным уровнем сложности
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.Wrap(err, op)
	}
	if err := s.rep.UpdatePassword(ctx, userID, string(hash)); err != nil {
		return errors.Wrap(err, op)
	}
	logger.StandardInfo(
		ctx,
		fmt.Sprintf("successful update password: %v", hash),
		op)

	return nil
}

func (s *Service) UpdateEmail(ctx context.Context, userID string, email string) error {
	op := "internal.account.service.updateEmail"

	if err := s.rep.UpdateEmail(ctx, userID, email); err != nil {
		return errors.Wrap(err, op)
	}
	logger.StandardInfo(
		ctx,
		fmt.Sprintf("successful update email: %v", email),
		op)

	return nil
}

func (s *Service) updateRole(ctx context.Context, userID string) error {
	op := "internal.account.service.updateRole"

	if err := s.rep.UpdateRoleToAuthor(ctx, userID); err != nil {
		return errors.Wrap(err, op)
	}
	logger.StandardInfo(ctx, "successful change role", op)
	return nil
}

func (s *Service) initPage(ctx context.Context, userID string) error {
	op := "internal.account.service.initPage"

	if err := s.rep.InitPage(ctx, userID); err != nil {
		return errors.Wrap(err, op)
	}
	logger.StandardInfo(ctx, "successful create page for user with userID: %v", userID)
	return nil
}
