package service

import (
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/pkg/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/service/interfaces"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

type Behavior struct {
	rep interfaces.ContentRepository
}

func New(repository interfaces.ContentRepository) *Behavior {
	if repository == nil {
		panic(errors.New("repository is nil"))
	}
	return &Behavior{repository}
}

func (b *Behavior) CreatePost(authorId string, title string, content string, layer int) (string, error) {
	op := "service.behavior.CreatePost"

	logger.StandardDebugF(op, "Got authorId:=%v title:=%v content:=%v layer:=%v", authorId, title, content, layer)

	newPostId, err := uuid.NewV4()
	if err != nil {
		return "", errors.Wrap(global.ErrServer, op)
	}
	logger.StandardDebugF(op, "NewPostId=%v newPostId=%v", newPostId, newPostId)

	uuidUserId, err := uuid.FromString(authorId)
	if err != nil {
		return "", errors.Wrap(global.ErrServer, op)
	}
	logger.StandardDebugF(op, "uuidUserId=%v uuidUserId=%v", uuidUserId, newPostId)
	logger.StandardDebugF(op, "Want to insert user:=%v post=%v title=%v content=%v layer=%v", authorId, newPostId, title, content, layer)
	err = b.rep.InsertPost(uuidUserId, newPostId, title, content, layer)
	logger.StandardDebugF(op, "InsertPost=%v err=%v", newPostId, err)

	if err != nil {
		return "", errors.Wrap(global.ErrServer, op)
	}
	return newPostId.String(), nil
}

func (b *Behavior) UpdatePost(userID string, post models.Post) error {
	op := "service.behavior.UpdatePost"

	userIdUuid, err := uuid.FromString(userID)
	if err != nil {
		return errors.Wrap(global.ErrServer, op)
	}

	postIdUuid, err := uuid.FromString(post.PostId)
	if err != nil {
		return errors.Wrap(global.ErrServer, op)
	}

	authorIdUuid, err := b.rep.GetAuthorByPost(postIdUuid)
	if err != nil {
		return errors.Wrap(global.ErrServer, op)
	}

	if authorIdUuid != userIdUuid {
		logger.StandardDebugF(op, "authorIdUuid=%v userIdUuid=%v", authorIdUuid, userIdUuid)
		logger.StandardDebug(op, "НАДО ЧТОБЫ notEnoughRigths")
		return errors.Wrap(global.ErrServer, op)
	}

	err = b.rep.UpdatePost(userIdUuid, postIdUuid, post)
	if err != nil {
		return errors.Wrap(global.ErrServer, op)
	}

	return nil
}

//func (b *Behavior) DeletePost(userID string, postID string) error {
//	op := "service.behavior.DeletePost"
//	userIdUuid, err := uuid.FromString(userID)
//	if err != nil {
//		return errors.Wrap(global.ErrServer, op)
//	}
//	postIdUuid, err := uuid.FromString(postID)
//	if err != nil {
//		return errors.Wrap(global.ErrServer, op)
//	}
//	authorIdUuid, err := b.rep.GetAuthorByPost(postIdUuid)
//	if err != nil {
//		return errors.Wrap(global.ErrServer, op)
//
//	}
//
//}

func (b *Behavior) LikePost(userId, postId string) (int, error) {
	op := "service.behavior.LikePost"

	userIdUuid, err := uuid.FromString(userId)
	if err != nil {
		return 0, errors.Wrap(global.ErrServer, op)
	}

	postIdUuid, err := uuid.FromString(postId)
	if err != nil {
		return 0, errors.Wrap(global.ErrServer, op)
	}

	isLike, err := b.rep.IsLikePutPost(userIdUuid, postIdUuid)
	if err != nil {
		return 0, errors.Wrap(global.ErrServer, op)
	}
	if isLike {
		err = b.rep.InsertLikePost(userIdUuid, postIdUuid)
		if err != nil {
			return 0, errors.Wrap(global.ErrServer, op)
		}
	} else {
		err = b.rep.DeleteLikePost(userIdUuid, postIdUuid)
		if err != nil {
			return 0, errors.Wrap(global.ErrServer, op)
		}
	}

	countLikes, err := b.rep.GetPostLikes(postIdUuid)
	if err != nil {
		return 0, errors.Wrap(global.ErrServer, op)
	}

	return countLikes, nil
}

func (b *Behavior) DeletePost(postId, userId string) error {
	return nil
}
