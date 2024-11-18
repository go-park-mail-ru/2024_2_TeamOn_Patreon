package postgresql

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/static"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/utils"
	"github.com/pkg/errors"
)

func (cr *ContentRepository) SaveFile(ctx context.Context, postID string, file []byte, fileExtension string) error {
	op := "internal.content.repository.SaveFile"

	// Директория для сохранения файлов
	fileDir := "./static/post"

	// Формирование ID
	fileID := utils.GenerateUUID()

	// Формируем путь к файлу
	filePath := static.CreateFilePath(fileDir, fileID, fileExtension)

	// Сохраняем файл в хранилище
	logger.StandardDebugF(ctx, op, "want to save new file with path %v", filePath)
	err := static.SaveFile(file, filePath)
	if err != nil {
		return errors.Wrap(err, op)
	}

	// // Запрос на создание новой записи о новой аватарке
	// query := `
	// 	INSERT INTO avatar (avatar_id, user_id, avatar_url)
	// 	VALUES ($1, $2, $3)
	// 	ON CONFLICT (avatar_id) DO UPDATE
	// 	SET user_id = EXCLUDED.user_id, avatar_url = EXCLUDED.avatar_url
	// `
	// Выполняем запрос
	// if _, err := cr.db.Exec(ctx, query, fileID, userID, filePath); err != nil {
	// 	return errors.Wrap(err, op)
	// }

	return nil
}
