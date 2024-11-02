// Бизнес-логика сервиса Account

package service

import (
	"context"
	"fmt"

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

// GetAuthorDataByID - получение данных автора по authorID
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
