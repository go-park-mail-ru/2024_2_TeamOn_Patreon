package postgresql

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/pkg/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/pkg/errors"
)

const (
	// $1 - postID
	getMediaByPostIDSQL = `
	SELECT content_id, default_content_type_name, content_url
	FROM content
	JOIN content_type USING (content_type_id)
	WHERE post_id = $1
`
)

func (cr *ContentRepository) GetContentsByPost(ctx context.Context, postID string) ([]*models.Media, error) {
	op := "internal.content.repository.getContentsByPost"

	medias := make([]*models.Media, 0)

	rows, err := cr.db.Query(ctx, getMediaByPostIDSQL, postID)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	defer rows.Close()

	var (
		contentID   string
		contentType string
		contentURL  string
	)

	for rows.Next() {
		if err = rows.Scan(&contentID, &contentType, &contentURL); err != nil {
			return nil, errors.Wrap(err, op)
		}
		logger.StandardDebugF(ctx, op,
			"Got  media: post_id=%v title=%v authorId=%v authorUsername=%v likes=%v created_date=%v",
			contentID, contentType, contentURL)
		medias = append(medias, &models.Media{
			MediaID:   contentID,
			MediaType: contentType,
			MediaURL:  contentURL,
		})

	}

	return medias, nil
}
