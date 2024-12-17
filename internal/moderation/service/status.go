package service

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/moderation/pkg/models"
	"github.com/pkg/errors"
)

func (s *Service) updatePostStatus(ctx context.Context, postID string, status string) error {
	// Обновить статус с помощью репозитория
	err := s.rep.UpdatePostStatus(ctx, postID, status)
	if err != nil {
		return errors.Wrap(err, "update post status")
	}
	return nil
}

func (s *Service) getPostsByStatus(ctx context.Context, status string, limit, offset int) ([]*models.Post, error) {
	// Достаем из репозитория по статусу посты в отсортированном порядке, сначала старые
	posts, err := s.rep.GetPostsByStatus(ctx, status, limit, offset)
	if err != nil {
		return nil, errors.Wrap(err, "get posts by status")
	}
	return posts, nil
}

func (s *Service) getStatusPost(ctx context.Context, postID string) (string, error) {
	// Достать статус поста по ИД
	status, err := s.rep.GetStatusByPostID(ctx, postID)
	if err != nil {
		return "", errors.Wrap(err, "get status by post id")
	}
	return status, nil
}
