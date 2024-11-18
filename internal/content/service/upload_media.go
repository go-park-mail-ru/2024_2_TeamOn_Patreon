package service

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
	"github.com/pkg/errors"
)

func (b *Behavior) UploadMedia(ctx context.Context, userID, postID string, file []byte, fileExtension, key string) error {
	op := "content.service.behavior.UploadMedia"

	// Валидация userID
	if ok := utils.IsValidUUIDv4(userID); !ok {
		return errors.Wrap(global.ErrIsInvalidUUID, op)
	}

	// Проверяем, что пользователь является автором поста
	isAuthor, err := b.isUserAuthorOfPost(ctx, postID, userID)
	if err != nil {
		return errors.Wrap(err, op)
	}
	if !isAuthor {
		return errors.Wrap(global.ErrNotEnoughRights, op)
	}

	// Сохраняем файл
	if err := b.saveFile(ctx, userID, postID, file, fileExtension, key); err != nil {
		return errors.Wrap(err, op)
	}

	return nil
}

func (b *Behavior) saveFile(ctx context.Context, userID, postID string, file []byte, fileExtension, key string) error {
	op := "content.service.behavior.saveFile"

	logger.StandardDebugF(ctx, op, "want to save new file: %s", key)
	if err := b.rep.SaveFile(ctx, postID, file, fileExtension); err != nil {
		return errors.Wrap(err, op)
	}

	logger.StandardInfo(
		ctx,
		fmt.Sprintf("successful save %s for userID %v", key, userID),
		op,
	)
	return nil
}
