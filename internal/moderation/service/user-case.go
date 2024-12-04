package service

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/moderation/pkg/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/pkg/errors"
)

func (s *Service) ComplaintPost(ctx context.Context, postID string, userID string) error {
	op := "moderation.service.ComplaintPost"

	// Проверяем может ли юзер видеть пост - в таком случае ошибка
	ok, err := s.userCanSeePost(ctx, userID, postID)
	if err != nil {
		err = errors.Wrap(err, op)
		return errors.Wrap(err, "userCanSeePost")
	}
	if !ok {
		err = errors.Wrap(global.ErrNotEnoughRights, op)
		return errors.Wrap(err, "user can't see post")
	}

	// Меняем статус поста на тот, на который пожаловались
	err = s.updatePostStatus(ctx, postID, models.Complained)
	if err != nil {
		err = errors.Wrap(err, op)
		return errors.Wrap(err, "updatePostStatus")
	}
	return nil
}
