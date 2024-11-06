package service

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/pkg/errors"
)

func (s *Service) Subscribe(ctx context.Context, userID, authorID string) error {
	op := "internal.author.service.Subscribe"

	if err := s.rep.Subscribe(ctx, userID, authorID); err != nil {
		return errors.Wrap(err, op)
	}
	logger.StandardInfo(
		ctx,
		fmt.Sprintf("successful subscribe info: %v from %v", authorID, userID),
		op)

	return nil
}
