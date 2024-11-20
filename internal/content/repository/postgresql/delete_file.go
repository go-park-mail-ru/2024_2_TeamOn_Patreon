package postgresql

import (
	"context"

	"github.com/pkg/errors"
)

func (cr *ContentRepository) DeleteFile(ctx context.Context, postID, fileID string) error {
	op := "internal.content.repository.DeleteFile"

	query := `
		DELETE 
		FROM content 
		WHERE content_id = $1 AND post_id = $2
	`

	_, err := cr.db.Exec(ctx, query, fileID, postID)
	if err != nil {
		return errors.Wrap(err, op)
	}
	return nil
}
