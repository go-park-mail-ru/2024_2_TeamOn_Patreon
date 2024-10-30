package interfaces

import (
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/pkg/models"
)

// ContentBehavior интерфейс с которым взаимодействует уровень controller
// Т.е. те, методы, которые нужны от service на уровне controller
type ContentBehavior interface {
	CreatePost(authorId string, title string, content string, layer int) (string, error)

	GetPopularPosts(offset, limits int) ([]models.Post, error)

	UpdatePost(authorId string, post models.Post) error

	LikePost(userId, postId string) (int, error)

	DeletePost(userId, postId string) error
}
