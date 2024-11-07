package service

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/pkg/errors"
)

func (s *Service) Subscribe(ctx context.Context, userID, authorID string) error {
	op := "internal.author.service.Subscribe"
	logger.StandardInfo(
		ctx,
		fmt.Sprintf("want to subscribe info: %v from %v", authorID, userID),
		op)
	isSub, err := s.rep.Subscribe(ctx, userID, authorID)
	if err != nil {
		return errors.Wrap(err, op)
	}
	if isSub {
		logger.StandardInfo(
			ctx,
			fmt.Sprintf("successful subscribe info: %v from %v", authorID, userID),
			op)
	}
	if !isSub {
		logger.StandardDebugF(ctx, op, "successful unscribe info: %v from %v", authorID, userID)
	}

	return nil
}
