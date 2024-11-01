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

	// ОБращаемся в репозиторий для получения пути до аватара
	avatarPath, err := s.rep.AvatarPathByID(ctx, userID)
	if err != nil {
		logger.StandardDebugF(op, "fail get avatarPath: {%v}", err)
		return nil, err
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

	if err := updateUsername(s, ctx, op, userID, username); err != nil {
		return fmt.Errorf("fail update username | in %v", op)
	}
	if err := updatePassword(s, ctx, op, userID, password); err != nil {
		return fmt.Errorf("fail update password | in %v", op)
	}
	if err := updateEmail(s, ctx, op, userID, email); err != nil {
		return fmt.Errorf("fail update password | in %v", op)
	}

	return nil
}

// PostAccountUpdateAvatar - изменение аватарки аккаунта по userID
func (s *Service) PostAccountUpdateAvatar(ctx context.Context, userID string, avatarFile multipart.File, fileName string) error {
	op := "internal.account.service.PostAccountUpdateAvatar"

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
		fmt.Sprintf("new file name: %s", fileFullName),
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

func updateUsername(s *Service, ctx context.Context, op string, userID string, username string) error {
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

func updatePassword(s *Service, ctx context.Context, op string, userID string, password string) error {

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

func updateEmail(s *Service, ctx context.Context, op string, userID string, email string) error {
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
