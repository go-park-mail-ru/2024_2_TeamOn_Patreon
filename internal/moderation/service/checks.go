package service

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/moderation/pkg/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	models2 "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	"github.com/pkg/errors"
)

func (s *Service) isUserModerator(ctx context.Context, userID string) (bool, error) {
	// TODO: Достаем из БД роль
	// TODO: Проверяем, что роль соответствует нашей роли модартора
	return true, nil
}

func isValidPostStatus(status string) bool {
	return models.CheckStatus(status)
}

func isValidPostFilter(filter string) bool {
	return models.CheckFilter(filter)
}

func isValidPostDecision(status string) bool {
	return models.CheckDecision(status)
}

func (s *Service) userCanSeePost(ctx context.Context, userID, postID string) (bool, error) {
	op := "service.behavior.userCanSeePost"

	authorPost, err := s.rep.GetAuthorOfPost(ctx, postID)
	if err != nil {
		return false, errors.Wrap(err, op)
	}
	if authorPost == userID {
		return true, nil
	}

	userLayer, err := s.rep.GetUserLayerOfAuthor(ctx, userID, authorPost)
	logger.StandardDebugF(ctx, op, "Get userLayer %v for Author %v for User %v", userLayer, authorPost, userID)
	if err != nil {
		return false, errors.Wrap(err, op)
	}

	postLayer, err := s.rep.GetPostLayerBuPostID(ctx, postID)
	if err != nil {
		return false, errors.Wrap(err, op)
	}
	logger.StandardDebugF(ctx, op, "Get postLayer %v for Author %v for User %v", postLayer, authorPost, userID)

	if userLayer >= postLayer {
		return true, nil
	}
	return false, nil

}

func validateLimitOffset(limit, offset int) (int, int) {
	opt := models2.FeedOpt{Offset: offset, Limit: limit}
	opt.Validate()
	return opt.Limit, opt.Offset
}
