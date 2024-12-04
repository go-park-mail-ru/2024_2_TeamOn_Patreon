package service

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/moderation/pkg/models"
)

func (s *Service) updatePostStatus(ctx context.Context, postID string, status string) error {
	// TODO: обновить статус с помощью репозитория
	return nil
}

func (s *Service) getPostsByStatus(ctx context.Context, status string, limit, offset int) ([]*models.Post, error) {
	// TODO: Достаем из репозитория по статусу посты в отсортированном порядке, сначала старые
	return nil, nil
}
