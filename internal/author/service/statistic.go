package service

import (
	"context"

	pkgModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/author/pkg/models"
	global "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	logger "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/pkg/errors"
)

const (
	DAY   = "day"
	MONTH = "month"
	YEAR  = "year"
)

func (s *Service) GetStatisticPosts(ctx context.Context, userID, time string) (*pkgModels.Graphic, error) {
	op := "internal.author.service.GetStatisticPosts"

	points := &pkgModels.Graphic{}

	// Пользователь является автором
	logger.StandardDebugF(ctx, op, "Want to check: user=%v is author", userID)
	isAuthor, err := s.isUserAuthor(ctx, userID)
	if err != nil {
		return points, errors.Wrap(err, op)
	}
	if !isAuthor {
		logger.StandardDebugF(ctx, op, "Fail: user is not author")
		return points, errors.Wrap(global.ErrNotEnoughRights, op)
	}

	// Валидация времени
	if time == DAY {
		points, err = s.statByDays(ctx, userID)
	} else if time == MONTH {
		points, err = s.statByMonth(ctx, userID)
	} else if time == YEAR {
		points, err = s.statByYear(ctx, userID)
	} else {
		return points, errors.Wrap(global.ErrBadTime, op)
	}
	if err != nil {
		return points, errors.Wrap(err, op)
	}

	logger.StandardDebugF(ctx, op, "Successful get stat for user=%v", userID)

	return points, nil
}

// isUserAuthor проверяет что пользователь является автором
func (s *Service) isUserAuthor(ctx context.Context, userID string) (bool, error) {
	op := "service.behavior.isUserAuthor"

	pageID, err := s.rep.GetAuthorPageID(ctx, userID)
	if err != nil {
		return false, errors.Wrap(err, op)
	}
	if pageID == "" {
		return false, nil
	}
	logger.StandardDebugF(ctx, op, "Got page=%v of author=%v", pageID, userID)
	return true, nil
}

func (s *Service) statByDays(ctx context.Context, userID string) (*pkgModels.Graphic, error) {
	op := "internal.author.service.statByDays"

	logger.StandardDebugF(ctx, op, "Want to get stat by day for user=%v", userID)

	points, err := s.rep.GetStatByDay(ctx, userID)
	if err != nil {
		return nil, errors.Wrap(global.ErrServer, op)
	}

	return points, nil
}

func (s *Service) statByMonth(ctx context.Context, userID string) (*pkgModels.Graphic, error) {
	op := "internal.author.service.statByMonth"

	logger.StandardDebugF(ctx, op, "Want to get stat by month for user=%v", userID)

	points, err := s.rep.GetStatByMonth(ctx, userID)
	if err != nil {
		return nil, errors.Wrap(global.ErrServer, op)
	}

	return points, nil
}

func (s *Service) statByYear(ctx context.Context, userID string) (*pkgModels.Graphic, error) {
	op := "internal.author.service.statByYear"

	points, err := s.rep.GetStatByYear(ctx, userID)
	if err != nil {
		return nil, errors.Wrap(global.ErrServer, op)
	}

	return points, nil
}
