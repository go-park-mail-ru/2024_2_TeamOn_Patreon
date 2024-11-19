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

	// Получаем content_type_id из БД
	contentTypeID, err := cr.getDefaultTypeID(ctx, fileID, fileExtension)
	if err != nil {
		return errors.Wrap(err, op)
	}

	// Создаём запись в БД о новом файле
	query := `
		INSERT INTO content (content_id, post_id, content_type_id, content_url) 
		VALUES ($1, $2, $3, $4)
	`

	if _, err := cr.db.Exec(ctx, query, fileID, postID, contentTypeID, filePath); err != nil {
		return errors.Wrap(err, op)
	}

	return nil
}

func (cr *ContentRepository) getDefaultTypeID(ctx context.Context, fileID, fileExtension string) (string, error) {
	op := "internal.content.repository.getDefaultTypeID"

	logger.StandardDebugF(ctx, op, "want to get standard extension for file with fileID %v", fileID)
	query := `
		SELECT content_type_id 
		FROM content_type
		WHERE default_content_type_name = $1
	`

	rows, err := cr.db.Query(ctx, query, fileExtension)
	if err != nil {
		return "", errors.Wrap(err, op)
	}

	defer rows.Close()

	var contentTypeID string

	for rows.Next() {
		if err := rows.Scan(&contentTypeID); err != nil {
			logger.StandardDebugF(ctx, op, "get default type ID failed: %v", err)
			return "", errors.Wrap(err, op)
		}
	}
	logger.StandardDebugF(ctx, op, "got default type ID='%s' for extension='%s'", contentTypeID, fileExtension)
	return contentTypeID, nil
}
