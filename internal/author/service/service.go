// Бизнес-логика сервиса Account

package service

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	cModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/author/controller/models"
	interfaces "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/author/service/interfaces"
	logger "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/pkg/errors"
)

type Service struct {
	rep interfaces.AuthorRepository
}

func New(repository interfaces.AuthorRepository) *Service {
	return &Service{repository}
}

func (s *Service) GetAuthorDataByID(ctx context.Context, authorID string) (cModels.Author, error) {
	op := "internal.author.service.GetAuthorDataByID"

	// получаем данные автора в формате service model
	author, err := s.rep.AuthorByID(ctx, authorID)
	if err != nil {
		logger.StandardDebugF(op, "fail get author: {%v}", err)
		return cModels.Author{}, err
	}

	logger.StandardInfo(
		fmt.Sprintf("successful get author=%v with authorID='%v'", author.Username, authorID),
		op)

	// по хорошему здесь должен быть маппер
	accountData := cModels.Author{
		Username:  author.Username,
		Info:      author.Info.String,
		Followers: author.Followers,
		// Subscriptions:
	}
	return accountData, nil
}

func (s *Service) PostUpdateInfo(ctx context.Context, authorID string, info string) error {
	op := "internal.author.service.PostUpdateInfo"

	if info != "" {
		if err := s.rep.UpdateInfo(ctx, authorID, info); err != nil {
			return err
		}
		logger.StandardInfo(
			fmt.Sprintf("successful update info: %v", info),
			op)
	}
	return nil
}

func (s *Service) GetAuthorPayments(ctx context.Context, authorID string) (cModels.Payments, error) {
	op := "internal.author.service.GetAuthorPayments"

	authorPayments := cModels.Payments{}

	// получаем сумму выплат int
	payments, err := s.rep.Payments(ctx, authorID)
	if err != nil {
		logger.StandardDebugF(op, "fail get payments: {%v}", err)
		return authorPayments, err
	}

	logger.StandardInfo(
		fmt.Sprintf("successful get payments=%v for authorID='%v'", payments, authorID),
		op)

	authorPayments = cModels.Payments{
		Amount: payments,
	}

	return authorPayments, nil
}

func (s *Service) GetBackgroundByID(ctx context.Context, authorID string) ([]byte, error) {
	op := "internal.author.service.GetBackgroundByID"

	// ОБращаемся в репозиторий для получения пути до фона
	logger.StandardDebugF(op, "want to find backgroundPATH for userID = %v", authorID)

	backgroundPath, err := s.rep.BackgroundPathByID(ctx, authorID)
	if err != nil {
		// Если не получилось найти путь фона -> 404
		return nil, errors.Wrap(err, op)
	}

	// По найденному пути открываем файл фона
	logger.StandardDebugF(op, "want to read file with path: %v", backgroundPath)
	avatar, err := os.ReadFile(backgroundPath)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	logger.StandardInfo(
		fmt.Sprintf("successful get background file with path %v for author with authorID %v", backgroundPath, authorID),
		op)

	return avatar, nil
}

func (s *Service) PostUpdateBackground(ctx context.Context, authorID string, backgroundFile multipart.File, fileName string) error {
	op := "internal.author.service.PostAccountUpdateBackground"

	// Удаляем старый фон, если он есть
	logger.StandardDebugF(op, "want to delete old background file")
	if err := s.deleteOldBackground(ctx, authorID); err != nil {
		return errors.Wrap(err, op)
	}

	// Сохраняем новый
	logger.StandardDebugF(op, "want to save new background file")
	if err := s.saveNewBackground(ctx, authorID, backgroundFile, fileName); err != nil {
		return errors.Wrap(err, op)
	}
	return nil
}

func (s *Service) PostTip(ctx context.Context, userID, authorID string, cost int, message string) error {
	op := "internal.author.service.PostTip"

	logger.StandardDebugF(op, "want to save new tip")

	if err := s.rep.NewTip(ctx, userID, authorID, cost, message); err != nil {
		return errors.Wrap(err, op)
	}
	logger.StandardInfo(
		fmt.Sprintf("successful send tip: cost=%v, message=%v from user=%v to author=%v", cost, message, userID, authorID),
		op)

	return nil
}

func (s *Service) deleteOldBackground(ctx context.Context, authorID string) error {
	op := "internal.author.service.deleteOldBackground"

	// Получаем путь до старого фона
	logger.StandardDebugF(op, "want to get path to old background for authorID %v", authorID)
	oldBackgroundPath, err := s.rep.BackgroundPathByID(ctx, authorID)

	if err != nil {
		logger.StandardInfo(
			fmt.Sprintf("old background doesn`t exist for author with authorID %s", authorID),
			op,
		)
		return nil
	}

	logger.StandardDebugF(op, "want to delete old background with path %v", oldBackgroundPath)
	if err := os.Remove(oldBackgroundPath); err != nil {
		return errors.Wrap(err, op)
	}
	logger.StandardInfo(
		fmt.Sprintf("successful delete old background for authorID %s", authorID),
		op,
	)
	return nil
}

func (s *Service) saveNewBackground(ctx context.Context, authorID string, backgroundFile multipart.File, fileName string) error {
	op := "internal.author.service.saveNewAvatar"

	// Директория для сохранения фона
	backgroundDir := "./static/background"

	// Получение формата загрузочного файла из его названия
	backgroundFormat := filepath.Ext(fileName)

	// Формирование ID
	backgroundID := s.rep.GenerateID()

	// Полное имя сохраняемого файла
	fileFullName := backgroundID + backgroundFormat

	// Формируем путь к файлу из папки сохранения и названия файла
	backgroundPath := filepath.Join(backgroundDir, fileFullName)

	logger.StandardDebugF(op, "want to save new file with path %v", backgroundPath)
	out, err := os.Create(backgroundPath)
	if err != nil {
		return fmt.Errorf(op, err)
	}
	defer out.Close()

	// Сохраняем файл
	logger.StandardDebugF(op, "want to copy new background to path %v", backgroundPath)
	if _, err := io.Copy(out, backgroundFile); err != nil {
		return fmt.Errorf(op, err)
	}

	// Обновляем информацию в БД
	logger.StandardDebugF(op, "want to save background URL %v in DB", backgroundPath)
	if err := s.rep.UpdateBackground(ctx, authorID, backgroundPath); err != nil {
		return errors.Wrap(err, op)
	}

	logger.StandardInfo(
		fmt.Sprintf("successful save new background with backgroundPath %v for authorID %v", backgroundPath, authorID),
		op,
	)
	return nil
}
