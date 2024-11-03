// Бизнес-логика сервиса Account

package service

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	cModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/account/controller/models"
	interfaces "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/account/service/interfaces"
	logger "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	rep interfaces.AccountRepository
}

func New(repository interfaces.AccountRepository) *Service {
	return &Service{repository}
}

// GetAccDataByID - получение данных аккаунта по userID
func (s *Service) GetAccDataByID(ctx context.Context, userID string) (cModels.Account, error) {
	op := "internal.account.service.GetAccDataByID"

	// получаем данные пользователя в формате service model
	user, err := s.rep.UserByID(ctx, userID)
	if err != nil {
		logger.StandardDebugF(op, "fail get user: {%v}", err)
		return cModels.Account{}, err
	}

	logger.StandardInfo(
		fmt.Sprintf("successful get user=%v with userID='%v'", user.Username, user.UserID),
		op)

	// по хорошему здесь должен быть маппер
	accountData := cModels.Account{
		Username: user.Username,
		Email:    user.Email.String,
		Role:     user.Role,
		// Subscriptions:
	}
	return accountData, nil
}

// GetAvatarByID - получение аватарки пользователя по userID
func (s *Service) GetAvatarByID(ctx context.Context, userID string) ([]byte, error) {
	op := "internal.account.service.GetAvatarByID"
	defaultAvatarPath := "static/avatar/default.jpg"

	// ОБращаемся в репозиторий для получения пути до аватара
	avatarPath, err := s.rep.AvatarPathByID(ctx, userID)
	if err != nil {
		logger.StandardDebugF(op, "fail get avatarPath: {%v}", err)
		// Если не получилось достать аватар, то возвращаем дефолтный
		avatarPath = defaultAvatarPath
	}

	// По указанному пути открываем файл аватара
	avatar, err := os.ReadFile(avatarPath)
	if err != nil {
		logger.StandardDebugF(op, "file with URL {%v} not found {%v}", avatarPath, err)
		return nil, err
	}

	logger.StandardInfo(
		fmt.Sprintln("successful get avatar file"),
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
	if err := s.deleteOldAvatar(ctx, op, userID); err != nil {
		return err
	}

	// Сохраняем новый
	if err := s.saveNewAvatar(ctx, op, userID, avatarFile, fileName); err != nil {
		return err
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
		if err := s.rep.UpdatePassword(ctx, userID, email); err != nil {
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

func (s *Service) deleteOldAvatar(ctx context.Context, op string, userID string) error {
	oldAvatarPath, err := s.rep.AvatarPathByID(ctx, userID)
	if err == nil {
		if err := os.Remove(oldAvatarPath); err != nil {
			return fmt.Errorf("error delete old avatar file {%v} | in %v", err, op)
		}
		logger.StandardInfo(
			fmt.Sprintf("old avatar deleted: %s", oldAvatarPath),
			op,
		)
	}
	return nil
}

func (s *Service) saveNewAvatar(ctx context.Context, op string, userID string, avatarFile multipart.File, fileName string) error {
	// Директория для сохранения аватаров
	avatarDir := "static/avatar"

	// Получение формата загрузочного файла из его названия
	avatarFormat := filepath.Ext(fileName)

	// Формирование ID
	avatarID := s.rep.GenerateID()

	// Полное имя сохраняемого файла
	fileFullName := avatarID + avatarFormat

	// Формируем путь к файлу из папки сохранения и названия файла
	avatarPath := filepath.Join(avatarDir, fileFullName)

	logger.StandardInfo(
		fmt.Sprintf("generated file name: %s", fileFullName),
		op,
	)
	out, err := os.Create(avatarPath)
	if err != nil {
		return fmt.Errorf("error creating avatar file | in %v", op)
	}
	defer out.Close()

	// Сохраняем файл
	if _, err := io.Copy(out, avatarFile); err != nil {
		return fmt.Errorf("error saving avatar file | in %v", op)
	}

	// Обновляем информацию в БД
	if err := s.rep.UpdateAvatar(ctx, userID, avatarID, avatarPath); err != nil {
		return fmt.Errorf("error updating avatar in DB | in %v", op)
	}
	return nil
}