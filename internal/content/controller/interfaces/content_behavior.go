package interfaces

import (
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/pkg/models"
	models2 "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/service/models"
)

// ContentBehavior интерфейс с которым взаимодействует уровень controller
// Т.е. те, методы, которые нужны от service на уровне controller
type ContentBehavior interface {
	CreatePost(userId string, title string, content string, layer int) (string, error)

	UpdatePost(userId string, post models.Post) error

	LikePost(userId, postId string) (int, error)

	DeletePost(userId, postId string) error

	// feed
	GetFeedSubscriptionGet(userId string, opt *models2.FeedOpt) ([]models.Post, error)
	GetPopularPostsForAnon(opt *models2.FeedOpt) ([]models.Post, error)
	GetPopularPostsForUser(userId string, opt *models2.FeedOpt) ([]models.Post, error)
}
