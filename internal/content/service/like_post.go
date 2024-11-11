package service

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
	"github.com/pkg/errors"
)

func (b *Behavior) LikePost(ctx context.Context, userID, postID string) (int, error) {
	op := "service.behavior.LikePost"

	if ok := utils.IsValidUUIDv4(userID); !ok {
		return 0, errors.Wrap(global.ErrIsInvalidUUID, op)
	}

	if ok := utils.IsValidUUIDv4(postID); !ok {
		return 0, errors.Wrap(global.ErrIsInvalidUUID, op)
	}

	// Может ли пользователь видеть пост (?)
	userCanSeePost, err := b.userCanSeePost(ctx, userID, postID)
	if err != nil {
		return 0, errors.Wrap(err, op)
	}

	if !userCanSeePost {
		return 0, errors.Wrap(global.ErrNotEnoughRights, op)
	}

	// Поставлен ли уже лайк
	isLike, err := b.isLikePutPost(ctx, userID, postID)
	if err != nil {
		return 0, errors.Wrap(err, op)
	}

	if !isLike {
		// Ставим лайк
		err = b.likePost(ctx, userID, postID)
		if err != nil {
			return 0, errors.Wrap(err, op)
		}
	} else {
		// Убираем лайк
		err = b.unlikePost(ctx, userID, postID)
		if err != nil {
			return 0, errors.Wrap(err, op)
		}
	}

	countLikes, err := b.countPostLikes(ctx, postID)
	if err != nil {
		return 0, errors.Wrap(err, op)
	}

	return countLikes, nil
}

func (b *Behavior) isLikePutPost(ctx context.Context, userID, postID string) (bool, error) {
	op := "service.behavior.IsLikePut"

	// Поставлен ли уже лайк
	likeID, err := b.rep.GetPostLikeID(ctx, userID, postID)
	if err != nil {
		return false, errors.Wrap(err, op)
	}

	if likeID == "" {
		return false, nil
	}
	return true, nil
}

func (b *Behavior) unlikePost(ctx context.Context, userID, postID string) error {
	op := "service.behavior.unlikePost"

	err := b.rep.DeleteLikePost(ctx, userID, postID)
	if err != nil {
		return errors.Wrap(err, op)
	}
	return nil
}

func (b *Behavior) likePost(ctx context.Context, userID, postID string) error {
	op := "service.behavior.likePost"

	err := b.rep.InsertLikePost(ctx, userID, postID)
	if err != nil {
		return errors.Wrap(err, op)
	}
	return nil
}

func (b *Behavior) countPostLikes(ctx context.Context, postID string) (int, error) {
	op := "service.behavior.countLikes"

	countLikes, err := b.rep.GetPostLikes(ctx, postID)
	if err != nil {
		return 0, errors.Wrap(err, op)
	}
	return countLikes, nil
}

func (b *Behavior) userCanSeePost(ctx context.Context, userID, postID string) (bool, error) {
	op := "service.behavior.userCanSeePost"

	authorPost, err := b.rep.GetAuthorOfPost(ctx, postID)
	if err != nil {
		return false, errors.Wrap(err, op)
	}
	if authorPost == userID {
		return true, nil
	}

	userLayer, err := b.rep.GetUserLayerOfAuthor(ctx, userID, authorPost)
	if err != nil {
		return false, errors.Wrap(err, op)
	}

	postLayer, err := b.rep.GetPostLayerBuPostID(ctx, postID)
	if err != nil {
		return false, errors.Wrap(err, op)
	}

	if userLayer >= postLayer {
		return true, nil
	}
	return false, nil

}
