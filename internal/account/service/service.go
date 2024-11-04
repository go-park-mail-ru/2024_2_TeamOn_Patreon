// Бизнес-логика сервиса Account

package service

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	interfaces "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/account/service/interfaces"
	sModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/account/service/models"
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
	logger.StandardDebugF(op, "want to get user data for userID = %v", userID)

	user, err := s.rep.UserByID(ctx, userID)
	if err != nil {
		return sModels.User{}, errors.Wrap(err, op)
	}

	logger.StandardInfo(
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
	logger.StandardDebugF(op, "want to get subscriptions for user with userID %v", userID)
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
		fmt.Sprintf("successful get subscriptions: %v", subscriptions),
		op)
	return subscriptions, nil
}

// GetAvatarByID - получение аватарки пользователя по userID
func (s *Service) GetAvatarByID(ctx context.Context, userID string) ([]byte, error) {
	op := "internal.account.service.GetAvatarByID"

	// ОБращаемся в репозиторий для получения пути до аватара
	logger.StandardDebugF(op, "want to find avatarPATH for userID = %v", userID)

	avatarPath, err := s.rep.AvatarPathByID(ctx, userID)
	if err != nil {
		// Если не получилось найти путь аватара -> 404
		return nil, errors.Wrap(err, op)
	}

	// По найденному пути открываем файл аватара
	logger.StandardDebugF(op, "want to read file with path: %v", avatarPath)
	avatar, err := os.ReadFile(avatarPath)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	logger.StandardInfo(
		fmt.Sprintf("successful get avatar file with path %v for user with userID %v", avatarPath, userID),
		op)

	return avatar, nil
}

// PostAccUpdateByID - изменение данных аккаунта по userID
func (s *Service) PostAccUpdateByID(ctx context.Context, userID string, username string, password string, email string) error {
	op := "internal.account.service.PostAccUpdateByID"

	if err := s.updateUsername(ctx, op, userID, username); err != nil {
		return fmt.Errorf("fail update username | in %v", op)
	}
	if err := s.updatePassword(ctx, op, userID, password); err != nil {
		return fmt.Errorf("fail update password | in %v", op)
	}
	if err := s.updateEmail(ctx, op, userID, email); err != nil {
		return fmt.Errorf("fail update password | in %v", op)
	}

	return nil
}

// PostUpdateAvatar - изменение аватарки аккаунта по userID
func (s *Service) PostUpdateAvatar(ctx context.Context, userID string, avatarFile multipart.File, fileName string) error {
	op := "internal.account.service.PostAccountUpdateAvatar"

	// Удаляем старый аватар, если он есть
	logger.StandardDebugF(op, "want to delete old avatar file")
	if err := s.deleteOldAvatar(ctx, userID); err != nil {
		return errors.Wrap(err, op)
	}

	// Сохраняем новый
	logger.StandardDebugF(op, "want to save new avatar file")
	if err := s.saveNewAvatar(ctx, userID, avatarFile, fileName); err != nil {
		return errors.Wrap(err, op)
	}
	return nil
}

// PostUpdateRole - изменение роли пользователя на "автор"
func (s *Service) PostUpdateRole(ctx context.Context, userID string) error {
	op := "internal.account.service.PostUpdateRole"

	// Обновляем поле роль с "reader" на "author"
	if err := s.updateRole(ctx, op, userID); err != nil {
		return fmt.Errorf("fail update role {%v} | in %v", err, op)
	}

	// Заполняем сущность page для нового автора
	if err := s.initPage(ctx, op, userID); err != nil {
		return fmt.Errorf("fail create page {%v} | in %v", err, op)
	}

	return nil
}

func (s *Service) updateUsername(ctx context.Context, op string, userID string, username string) error {
	if username != "" {
		if err := s.rep.UpdateUsername(ctx, userID, username); err != nil {
			return err
		}
		logger.StandardInfo(
			fmt.Sprintf("successful update username: %v", username),
			op)
	}
	return nil
}

func (s *Service) updatePassword(ctx context.Context, op string, userID string, password string) error {

	if password != "" {
		// Хеширование пароля с заданным уровнем сложности
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		if err := s.rep.UpdatePassword(ctx, userID, string(hash)); err != nil {
			return err
		}
		logger.StandardInfo(
			fmt.Sprintf("successful update password: %v", hash),
			op)
	}
	return nil
}

func (s *Service) updateEmail(ctx context.Context, op string, userID string, email string) error {
	if email != "" {
		if err := s.rep.UpdateEmail(ctx, userID, email); err != nil {
			return err
		}
		logger.StandardInfo(
			fmt.Sprintf("successful update email: %v", email),
			op)
	}
	return nil
}

func (s *Service) updateRole(ctx context.Context, op string, userID string) error {
	if err := s.rep.UpdateRole(ctx, userID); err != nil {
		return err
	}
	logger.StandardInfo("successful change role", op)
	return nil
}

func (s *Service) initPage(ctx context.Context, op string, userID string) error {
	if err := s.rep.InitPage(ctx, userID); err != nil {
		return err
	}
	logger.StandardInfo("successful create page", op)
	return nil
}

func (s *Service) deleteOldAvatar(ctx context.Context, userID string) error {
	op := "internal.account.service.deleteOldAvatar"

	// Получаем путь до старой аватарки
	logger.StandardDebugF(op, "want to get path to old avatar for userId %v", userID)
	oldAvatarPath, err := s.rep.AvatarPathByID(ctx, userID)

	if err != nil {
		logger.StandardInfo(
			fmt.Sprintf("old avatar doesn`t exist for user with userID %s", userID),
			op,
		)
		return nil
	}

	logger.StandardDebugF(op, "want to delete old avatar with path %v", oldAvatarPath)
	if err := os.Remove(oldAvatarPath); err != nil {
		return errors.Wrap(err, op)
	}
	logger.StandardInfo(
		fmt.Sprintf("successful delete old avatar fo userID %s", userID),
		op,
	)
	return nil
}

func (s *Service) saveNewAvatar(ctx context.Context, userID string, avatarFile multipart.File, fileName string) error {
	op := "internal.account.service.saveNewAvatar"

	// Директория для сохранения аватаров
	avatarDir := "./static/avatar"

	// Получение формата загрузочного файла из его названия
	avatarFormat := filepath.Ext(fileName)

	// Формирование ID
	avatarID := s.rep.GenerateID()

	// Полное имя сохраняемого файла
	fileFullName := avatarID + avatarFormat

	// Формируем путь к файлу из папки сохранения и названия файла
	avatarPath := filepath.Join(avatarDir, fileFullName)

	logger.StandardDebugF(op, "want to save new file with path %v", avatarPath)
	out, err := os.Create(avatarPath)
	if err != nil {
		return fmt.Errorf(op, err)
	}
	defer out.Close()

	// Сохраняем файл
	logger.StandardDebugF(op, "want to copy new avatar to path %v", avatarPath)
	if _, err := io.Copy(out, avatarFile); err != nil {
		return fmt.Errorf(op, err)
	}

	// Обновляем информацию в БД
	logger.StandardDebugF(op, "want to save avatar URL %v in DB", avatarPath)
	if err := s.rep.UpdateAvatar(ctx, userID, avatarID, avatarPath); err != nil {
		return errors.Wrap(err, op)
	}

	logger.StandardInfo(
		fmt.Sprintf("successful update save new avatar with avatar path %v for userID %v", avatarPath, userID),
		op,
	)
	return nil
}
