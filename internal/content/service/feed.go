package service

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/pkg/models"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

func (b *Behavior) getLikePostsForUser(ctx context.Context, userId uuid.UUID, posts []*models.Post) error {
	op := "internal.content.repository.feed.likePosts"

	err := b.rep.GetIsLikedForPosts(ctx, userId, posts)
	if err != nil {
		return errors.Wrap(err, op)
	}
	return nil
}
