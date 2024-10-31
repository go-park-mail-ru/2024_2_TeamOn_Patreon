package service

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

func (b *Behavior) LikePost(ctx context.Context, userId, postId string) (int, error) {
	op := "service.behavior.LikePost"

	userIdUuid, err := uuid.FromString(userId)
	if err != nil {
		return 0, errors.Wrap(global.ErrServer, op)
	}

	postIdUuid, err := uuid.FromString(postId)
	if err != nil {
		return 0, errors.Wrap(err, op)
	}

	// Может ли пользователь видеть пост (?)
	userCanSeePost, err := b.userCanSeePost(ctx, userIdUuid, postIdUuid)
	if err != nil {
		return 0, errors.Wrap(err, op)
	}

	if !userCanSeePost {
		return 0, errors.Wrap(global.ErrNotEnoughRights, op)
	}

	// Поставлен ли уже лайк
	isLike, err := b.isLikePutPost(ctx, userIdUuid, postIdUuid)
	if err != nil {
		return 0, errors.Wrap(err, op)
	}

	if !isLike {
		// Ставим лайк
		err = b.likePost(ctx, userIdUuid, postIdUuid)
		if err != nil {
			return 0, errors.Wrap(err, op)
		}
	} else {
		// Убираем лайк
		err = b.unlikePost(ctx, userIdUuid, postIdUuid)
		if err != nil {
			return 0, errors.Wrap(err, op)
		}
	}

	countLikes, err := b.countPostLikes(ctx, postIdUuid)
	if err != nil {
		return 0, errors.Wrap(err, op)
	}

	return countLikes, nil
}

func (b *Behavior) isLikePutPost(ctx context.Context, userId, postId uuid.UUID) (bool, error) {
	op := "service.behavior.IsLikePut"
	// Поставлен ли уже лайк
	likeId, err := b.rep.GetPostLikeId(ctx, userId, postId)
	if err != nil {
		return false, errors.Wrap(err, op)
	}

	if likeId == (uuid.UUID{}) {
		return false, nil
	}
	return true, nil
}

func (b *Behavior) unlikePost(ctx context.Context, userId, postId uuid.UUID) error {
	op := "service.behavior.unlikePost"

	err := b.rep.DeleteLikePost(ctx, userId, postId)
	if err != nil {
		return errors.Wrap(err, op)
	}
	return nil
}

func (b *Behavior) likePost(ctx context.Context, userId, postId uuid.UUID) error {
	op := "service.behavior.likePost"

	err := b.rep.InsertLikePost(ctx, userId, postId)
	if err != nil {
		return errors.Wrap(err, op)
	}
	return nil
}

func (b *Behavior) countPostLikes(ctx context.Context, postId uuid.UUID) (int, error) {
	op := "service.behavior.countLikes"

	countLikes, err := b.rep.GetPostLikes(ctx, postId)
	if err != nil {
		return 0, errors.Wrap(err, op)
	}
	return countLikes, nil
}

func (b *Behavior) userCanSeePost(ctx context.Context, userId, postId uuid.UUID) (bool, error) {
	op := "service.behavior.userCanSeePost"

	authorPost, err := b.rep.GetAuthorOfPost(ctx, postId)
	if err != nil {
		return false, errors.Wrap(err, op)
	}
	if authorPost == userId {
		return true, nil
	}

	userLayer, err := b.rep.GetUserLayerOfAuthor(ctx, userId, authorPost)
	if err != nil {
		return false, errors.Wrap(err, op)
	}

	postLayer, err := b.rep.GetPostLayerBuPostId(ctx, postId)
	if err != nil {
		return false, errors.Wrap(err, op)
	}

	if userLayer >= postLayer {
		return true, nil
	}
	return false, nil

}
