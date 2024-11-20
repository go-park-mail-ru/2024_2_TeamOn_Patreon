package service

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
	"github.com/pkg/errors"
)

func (b *Behavior) DeleteMedia(ctx context.Context, userID, postID string, mediaIDs []string) error {
	op := "service.behavior.DeleteMedia"

	// Валидация userID
	if ok := utils.IsValidUUIDv4(userID); !ok {
		logger.StandardWarnF(ctx, op, "userID=%v is not uuid", userID)
		return errors.Wrap(global.ErrIsInvalidUUID, op)
	}

	// Валидация postID
	if ok := utils.IsValidUUIDv4(postID); !ok {
		logger.StandardWarnF(ctx, op, "postID=%v is not uuid", postID)
		return errors.Wrap(global.ErrIsInvalidUUID, op)
	}

	// Валидация mediaID
	for _, mediaID := range mediaIDs {
		if !utils.IsValidUUIDv4(mediaID) {
			logger.StandardWarnF(ctx, op, "mediaID=%v is not uuid", mediaID)
			return errors.Wrap(global.ErrIsInvalidUUID, op)
		}
	}

	// Проверяем, что пользователь является автором поста
	isAuthor, err := b.isUserAuthorOfPost(ctx, postID, userID)
	if err != nil {
		return errors.Wrap(err, op)
	}
	if !isAuthor {
		return errors.Wrap(global.ErrNotEnoughRights, op)
	}

	// Удаляем файлы
	if err := b.deleteFile(ctx, postID, mediaIDs); err != nil {
		return errors.Wrap(err, op)
	}

	return nil
}

func (b *Behavior) deleteFile(ctx context.Context, postID string, mediaIDs []string) error {
	op := "content.service.behavior.deleteFile"

	for _, mediaID := range mediaIDs {
		logger.StandardDebugF(ctx, op, "want to delete file: %s", mediaID)

		if err := b.rep.DeleteFile(ctx, postID, mediaID); err != nil {
			return errors.Wrap(err, op)
		}

		logger.StandardInfo(ctx, fmt.Sprintf("successful delete file: %s", mediaID), op)
	}

	return nil
}
