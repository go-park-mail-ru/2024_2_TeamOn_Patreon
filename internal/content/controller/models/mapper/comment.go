package mapper

import (
	models2 "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/controller/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/pkg/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/validate"
)

func MapCommonCommentToTransportComment(comment *models.Comment) *models2.Comment {
	content := validate.Sanitize(comment.Content)

	return &models2.Comment{
		CommentID: comment.CommentID,
		Content:   content,
		Username:  comment.Username,
		UserID:    comment.UserID,
		CreatedAt: comment.CreatedAt.String(),
	}
}

func MapCommonCommentsToControllerComments(comments []*models.Comment) []*models2.Comment {
	tComments := make([]*models2.Comment, len(comments))
	for _, comment := range comments {
		tComments = append(tComments, MapCommonCommentToTransportComment(comment))
	}
	return tComments
}
