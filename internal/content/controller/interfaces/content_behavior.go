package interfaces

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/pkg/models"
	pkgModels "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
)

// ContentBehavior интерфейс с которым взаимодействует уровень controller
// Т.е. те, методы, которые нужны от service на уровне controller
type ContentBehavior interface {
	// post

	CreatePost(ctx context.Context, userID string, title string, content string, layer int) (string, error)
	UpdatePost(ctx context.Context, userID string, postID string, title string, about string) error
	LikePost(ctx context.Context, userID, postID string) (int, error)
	DeletePost(ctx context.Context, userID, postID string) error

	// feed

	GetPopularPosts(ctx context.Context, userID string, opt *pkgModels.FeedOpt) ([]*models.Post, error)
	GetFeedSubscription(ctx context.Context, userID string, opt *pkgModels.FeedOpt) ([]*models.Post, error)
	GetAuthorPosts(ctx context.Context, userID string, authorID string, opt *pkgModels.FeedOpt) ([]*models.Post, error)

	// media

	GetFile(ctx context.Context, userID string, postID string) ([]*models.Media, error)
	UploadMedia(ctx context.Context, userID, postID string, file []byte, fileExtension, key string) error
	DeleteMedia(ctx context.Context, userID, postID string, mediaIDs []string) error

	CommentBehavior
}

type CommentBehavior interface {
	CreateComment(ctx context.Context, userID string, postID string, content string) (string, error)
}
