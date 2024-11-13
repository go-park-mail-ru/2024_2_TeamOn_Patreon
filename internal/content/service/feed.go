package service

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/pkg/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/pkg/errors"
)

func (b *Behavior) getLikePostsForUser(ctx context.Context, userID string, posts []*models.Post) error {
	op := "internal.content.repository.feed.likePosts"

	err := b.rep.GetIsLikedForPosts(ctx, userID, posts)
	logger.StandardDebugF(ctx, op, "Got likes for user='%v' for len(posts)= %v with err={%v}",
		userID, len(posts), err)
	if err != nil {
		return errors.Wrap(err, op)
	}
	return nil
}
