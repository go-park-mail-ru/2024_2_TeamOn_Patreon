package service

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/moderation/pkg/models"
)

func (s Service) DecisionPost(ctx context.Context, postID string, userID string, status string) error {
	//TODO implement me
	panic("implement me")
}

func (s Service) GetPostsForModeration(ctx context.Context, userID string, filter string, limit, offset int) ([]*models.Post, error) {
	//TODO implement me
	panic("implement me")
}
