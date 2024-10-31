package interfaces

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/pkg/models"
	models2 "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
)

// ContentBehavior интерфейс с которым взаимодействует уровень controller
// Т.е. те, методы, которые нужны от service на уровне controller
type ContentBehavior interface {
	CreatePost(ctx context.Context, userId string, title string, content string, layer int) (string, error)

	UpdatePost(ctx context.Context, userId string, postId string, title string, about string) error

	LikePost(ctx context.Context, userId, postId string) (int, error)

	DeletePost(ctx context.Context, userId, postId string) error

	// feed
	GetPopularPosts(ctx context.Context, userId string, opt *models2.FeedOpt) ([]*models.Post, error)
	GetFeedSubscriptionGet(userId string, opt *models2.FeedOpt) ([]*models.Post, error)

	// страница автора

	GetAuthorPostsForAnon(authorId string, offset, limit int) ([]models.Post, error)
	GetAuthorPostsForUser(authorId, userId string, offset, limit int) ([]models.Post, error)
	GetAuthorPostsForMe(authorId string, offset, limit int) ([]models.Post, error)
}
