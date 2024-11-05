// Бизнес-логика сервиса Account

package service

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"

	interfaces "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/author/service/interfaces"
	sModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/author/service/models"
	logger "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/pkg/errors"
)

type Service struct {
	rep interfaces.AuthorRepository
}

func New(repository interfaces.AuthorRepository) *Service {
	return &Service{repository}
}

func (s *Service) GetAuthorDataByID(ctx context.Context, authorID string) (sModels.Author, error) {
	op := "internal.author.service.GetAuthorDataByID"

	// получаем данные автора в формате rep model
	author, err := s.rep.AuthorByID(ctx, authorID)
	if err != nil {
		return sModels.Author{}, errors.Wrap(err, op)
	}

	logger.StandardInfo(
		fmt.Sprintf("successful get author=%v with authorID='%v'", author.Username, authorID),
		op)

	authorData := sModels.Author{
		Username:  author.Username,
		Info:      author.Info.String,
		Followers: author.Followers,
	}
	return authorData, nil
}

func (s *Service) GetAuthorSubscriptions(ctx context.Context, authorID string) ([]sModels.Subscription, error) {
	op := "internal.account.service.GetAuthorSubscriptions"

	// получаем подписки пользователя из rep
	logger.StandardDebugF(op, "want to get subscriptions for author with authorID %v", authorID)
	repSubscriptions, err := s.rep.SubscriptionsByID(ctx, authorID)
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

func (s *Service) PostUpdateInfo(ctx context.Context, authorID string, info string) error {
	op := "internal.author.service.PostUpdateInfo"

	if err := s.rep.UpdateInfo(ctx, authorID, info); err != nil {
		return errors.Wrap(err, op)
	}
	logger.StandardInfo(
		fmt.Sprintf("successful update info: %v", info),
		op)

	return nil
}

func (s *Service) GetAuthorPayments(ctx context.Context, authorID string) (int, error) {
	op := "internal.author.service.GetAuthorPayments"

	// получаем сумму выплат int
	amount, err := s.rep.Payments(ctx, authorID)
	if err != nil {
		return 0, errors.Wrap(err, op)
	}

	logger.StandardInfo(
		fmt.Sprintf("successful get payments=%v for authorID='%v'", amount, authorID),
		op)

	return amount, nil
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
	if err := s.rep.DeleteBackground(ctx, authorID); err != nil {
		return errors.Wrap(err, op)
	}

	logger.StandardInfo(
		fmt.Sprintf("successful delete old background for authorID %s", authorID),
		op,
	)

	// Сохраняем новый
	logger.StandardDebugF(op, "want to save new background file")
	if err := s.rep.UpdateBackground(ctx, authorID, backgroundFile, fileName); err != nil {
		return errors.Wrap(err, op)
	}

	logger.StandardInfo(
		fmt.Sprintf("successful save new background for authorID %v", authorID),
		op,
	)

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
