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
	op := "internal.account.service.PostUpdateInfo"

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

func (s *Service) PostUpdateBackground(ctx context.Context, userID string, backgroundFile multipart.File, fileName string) error {
	op := "internal.account.service.PostAccountUpdateBackground"

	// Удаляем старый фон, если он есть
	if err := s.deleteOldBackground(ctx, op, userID); err != nil {
		return err
	}

	// Сохраняем новый
	if err := s.saveNewBackground(ctx, op, userID, backgroundFile, fileName); err != nil {
		return err
	}
	return nil
}

func (s *Service) deleteOldBackground(ctx context.Context, op string, userID string) error {
	oldBackgroundPath, err := s.rep.BackgroundPathByID(ctx, userID)
	if err == nil {
		if err := os.Remove(oldBackgroundPath); err != nil {
			return fmt.Errorf("error delete old background file {%v} | in %v", err, op)
		}
		logger.StandardInfo(
			fmt.Sprintf("old background deleted: %s", oldBackgroundPath),
			op,
		)
	}
	return nil
}

func (s *Service) saveNewBackground(ctx context.Context, op string, userID string, backgroundFile multipart.File, fileName string) error {
	// Директория для сохранения аватаров
	backgroundDir := "static/background"

	// Получение формата загрузочного файла из его названия
	backgroundFormat := filepath.Ext(fileName)

	// Формирование ID
	backgroundID := s.rep.GenerateID()

	// Полное имя сохраняемого файла
	fileFullName := backgroundID + backgroundFormat

	// Формируем путь к файлу из папки сохранения и названия файла
	backgroundPath := filepath.Join(backgroundDir, fileFullName)

	logger.StandardInfo(
		fmt.Sprintf("generated file name: %s", fileFullName),
		op,
	)
	out, err := os.Create(backgroundPath)
	if err != nil {
		return fmt.Errorf("error creating background file | in %v", op)
	}
	defer out.Close()

	// Сохраняем файл
	if _, err := io.Copy(out, backgroundFile); err != nil {
		return fmt.Errorf("error saving background file | in %v", op)
	}

	// Обновляем информацию в БД
	if err := s.rep.UpdateBackground(ctx, userID, backgroundPath); err != nil {
		return fmt.Errorf("error updating background in DB | in %v", op)
	}
	return nil
}
