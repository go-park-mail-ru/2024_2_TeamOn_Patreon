package service

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/moderation/pkg/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	pkgModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	"github.com/pkg/errors"
)

func (s *Service) isUserModerator(ctx context.Context, userID string) (bool, error) {
	// Достаем из БД роль
	role, err := s.rep.GetUserRole(ctx, userID)
	if err != nil {
		return false, errors.Wrap(err, "cant get user role")
	}

	// Проверяем, что роль соответствует нашей роли модератора
	return pkgModels.StringToRole(role) == pkgModels.Moderator, nil
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

	postLayer, err := s.rep.GetPostLayerByPostID(ctx, postID)
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
	opt := pkgModels.FeedOpt{Offset: offset, Limit: limit}
	opt.Validate()
	return opt.Limit, opt.Offset
}
