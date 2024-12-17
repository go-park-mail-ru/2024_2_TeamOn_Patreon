package service

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/pkg/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	bModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/validate"
	"github.com/pkg/errors"
)

func (b *Behavior) GetComments(ctx context.Context, userID string, postID string, opt *bModels.FeedOpt) ([]*models.Comment, error) {
	op := "content.behavior.GetComments"

	// Проверить может ли пользователь видеть пост
	ok, err := b.userCanSeePost(ctx, userID, postID)
	if err != nil {
		err = errors.Wrap(err, op)
		return nil, errors.Wrap(err, "user can see post")
	}
	if !ok {
		err = errors.Wrap(global.ErrNotEnoughRights, "cannot see post")
		return nil, errors.Wrap(err, op)
	}

	opt.Validate()

	// Достать комменты
	comments, err := b.rep.GetCommentsByPostID(ctx, postID, opt.Limit, opt.Offset)
	if err != nil {
		err = errors.Wrap(err, op)
		return nil, errors.Wrap(err, "get comments")
	}

	logger.StandardDebugF(ctx, op, "Got comments = %v", comments)
	return comments, nil
}

func (b *Behavior) CreateComment(ctx context.Context, userID string, postID string, content string) (string, error) {
	op := "content.behavior.CreateComment"
	// Валидация контента
	content, err := ValidateComment(content)
	if err != nil {
		err = errors.Wrap(err, op)
		return "", errors.Wrap(err, "validate comment")
	}

	// Проверка может ли юзер видеть пост
	ok, err := b.userCanSeePost(ctx, userID, postID)
	if err != nil {
		err = errors.Wrap(err, op)
		return "", errors.Wrap(err, "user can see post")
	}
	if !ok {
		err = errors.Wrap(global.ErrNotEnoughRights, "user cann't see post")
		return "", errors.Wrap(err, op)
	}

	// Генерация UUID
	commentID := utils.GenerateUUID()

	// Создание коммента с ююдом созданным
	err = b.createComment(ctx, userID, postID, commentID, content)
	if err != nil {
		err = errors.Wrap(err, op)
		return "", errors.Wrap(err, "create comment")
	}

	// Отправляем уведомление автору поста о новом комментарии
	if err := b.sendNotificationOfComment(ctx, userID, postID, content); err != nil {
		logger.StandardDebugF(ctx, op, "failed send notification to AUTHOR about new comment: %v", err)
	}

	// Возвращаем ид коммента
	return commentID, nil
}

func (b *Behavior) UpdateComment(ctx context.Context, userID string, commentID string, content string) error {
	op := "content.behavior.UpdateComment"
	// Валидация контента
	content, err := ValidateComment(content)
	if err != nil {
		err = errors.Wrap(err, op)
		return errors.Wrap(err, "validate comment")
	}

	// Проверка автор ли юзер этого коммента
	ok, err := b.isUserAuthorOfComment(ctx, userID, commentID)
	if err != nil {
		err = errors.Wrap(err, op)
		return errors.Wrap(err, "isUserAuthorOfComment")
	}
	if !ok {
		err = errors.Wrap(global.ErrNotEnoughRights, op)
		return errors.Wrap(err, "user can not update comment")
	}

	// Изменение коммента
	err = b.updateComment(ctx, userID, commentID, content)
	if err != nil {
		err = errors.Wrap(err, op)
		return errors.Wrap(err, "updateComment")
	}
	return nil
}

func (b *Behavior) DeleteComment(ctx context.Context, userID string, commentID string) error {
	op := "content.behavior.DeleteComment"
	// Проверка автор ли юзер этого коммента
	userIsAuthor, err := b.isUserAuthorOfComment(ctx, userID, commentID)
	if err != nil {
		err = errors.Wrap(err, op)
		return errors.Wrap(err, "isUserAuthorOfComment")
	}
	if !userIsAuthor {
		err = errors.Wrap(global.ErrNotEnoughRights, op)
		return errors.Wrap(err, "user can not delete comment")
	}

	// Удаление коммента
	err = b.deleteComment(ctx, userID, commentID)
	if err != nil {
		err = errors.Wrap(err, op)
		return errors.Wrap(err, "deleteComment")
	}
	return nil
}

func ValidateComment(comment string) (string, error) {
	if comment == "" {
		return "", global.ErrCommentDoesntExist
	}
	if len(comment) > 500 {
		return "", global.ErrCommentTooLong
	}
	return validate.Sanitize(comment), nil
}

func (b *Behavior) isUserAuthorOfComment(ctx context.Context, userID string, commentID string) (bool, error) {
	op := "content.behavior.IsUserAuthorOfComment"
	// Проверка что юзер автор коммента
	authorID, err := b.rep.GetUserIDByCommentID(ctx, commentID)
	if err != nil {
		return false, errors.Wrap(err, op)
	}
	if userID != authorID {
		return false, errors.Wrap(global.ErrNotEnoughRights, op)
	}
	return true, nil
}

func (b *Behavior) createComment(ctx context.Context, userID string, postID, commentID string, content string) error {
	op := "content.behavior.createComment"

	err := b.rep.CreateComment(ctx, userID, postID, commentID, content)
	if err != nil {
		return errors.Wrap(err, op)
	}
	return nil
}

func (b *Behavior) updateComment(ctx context.Context, userID string, commentID string, content string) error {
	op := "content.behavior.updateComment"

	err := b.rep.UpdateComment(ctx, commentID, content)
	if err != nil {
		return errors.Wrap(err, op)
	}
	return nil
}

func (b *Behavior) deleteComment(ctx context.Context, userID string, commentID string) error {
	op := "content.behavior.deleteComment"

	err := b.rep.DeleteComment(ctx, commentID)
	if err != nil {
		return errors.Wrap(err, op)
	}
	return nil
}

func (b *Behavior) sendNotificationOfComment(ctx context.Context, userID, postID, content string) error {
	op := "internal.author.service.sendNotificationOfComment"

	logger.StandardDebugF(ctx, op, "want to send notification to AUTHOR about new comment")

	// Имя пользователя
	username, err := b.rep.GetUsername(ctx, userID)
	if err != nil {
		return errors.Wrap(err, op)
	}

	// ID автора поста
	authorID, err := b.rep.GetAuthorOfPost(ctx, postID)
	if err != nil {
		return errors.Wrap(err, op)
	}

	//title поста
	postTitle, err := b.rep.GetTitleOfPost(ctx, postID)
	if err != nil {
		return errors.Wrap(err, op)
	}

	message := fmt.Sprintf("Пользователь @%v под вашим постом «%v» оставил комментарий: «%v»", username, postTitle, content)

	if err := b.rep.SendNotification(ctx, message, userID, authorID); err != nil {
		return errors.Wrap(err, op)
	}

	logger.StandardDebugF(ctx, op, "Successful send notification: %v", message)
	return nil
}
