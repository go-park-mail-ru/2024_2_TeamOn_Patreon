package service

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/moderation/pkg/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/pkg/errors"
)

// DecisionPost проставляет новый статус по решению модератора
func (s *Service) DecisionPost(ctx context.Context, postID string, userID string, status string) error {
	op := "moderation.service.DecisionPost"

	// Проверяем является ли юзер модератором
	ok, err := s.isUserModerator(ctx, userID)
	if err != nil {
		err = errors.Wrap(err, op)
		return errors.Wrap(err, "isUserModerator")
	}
	if !ok {
		err = errors.Wrap(global.ErrNotEnoughRights, op)
		return errors.Wrap(err, "user isn't moderator")
	}

	// Проверяем у поста один из тех статусов, что должен валидировать модератор?
	oldStatus, err := s.getStatusPost(ctx, postID)
	if err != nil {
		err = errors.Wrap(err, op)
		return errors.Wrap(err, "getStatusPost")
	}

	if !isValidPostFilter(oldStatus) {
		err = errors.Wrap(global.ErrNotEnoughRights, op)
		return errors.Wrap(err, "post's ok")
	}

	// Проверяем валидный ли новый статус
	if !isValidPostDecision(status) {
		err = errors.Wrap(global.ErrNotEnoughRights, op)
		return errors.Wrap(err, "decision status's invalid")
	}

	// Меняем статус на новый
	err = s.updatePostStatus(ctx, postID, status)
	if err != nil {
		err = errors.Wrap(err, op)
		return errors.Wrap(err, "updatePostStatus")
	}
	return nil
}

// GetPostsForModeration - возвращает посты в зависимости от фильтра для модератора
// В фильтре может быть одно из допустимых значений для модерации - BLOCKED | COMPLAINED
func (s *Service) GetPostsForModeration(ctx context.Context, userID string, filter string, limit, offset int) ([]*models.Post, error) {
	op := "moderation.service.GetPostsForModeration"

	// Проверка является ли юзер модератором
	ok, err := s.isUserModerator(ctx, userID)
	if err != nil {
		err = errors.Wrap(err, op)
		return nil, errors.Wrap(err, "isUserModerator")
	}
	if !ok {
		err = errors.Wrap(global.ErrNotEnoughRights, op)
		return nil, errors.Wrap(err, "user isn't moderator")
	}

	// Проверка валидный ли фильтр
	if !isValidPostFilter(filter) {
		err = errors.Wrap(global.ErrStatusIncorrect, op)
		return nil, errors.Wrap(err, "filter is invalid")
	}

	// Валидация лимит оффсет
	limit, offset = validateLimitOffset(limit, offset)

	// Достаем посты по фильтру из репозитория
	posts, err := s.getPostsByStatus(ctx, userID, limit, offset)
	if err != nil {
		err = errors.Wrap(err, op)
		return nil, errors.Wrap(err, "getPostsByStatus")
	}

	return posts, nil
}
