package interfaces

import "context"

type ModerationRepository interface {
	UserLayer
}

type UserLayer interface {
	// GetPostLayerBuPostID уровень поста по ид поста
	GetPostLayerBuPostID(ctx context.Context, postID string) (int, error)
	// GetAuthorOfPost - получение  ID автора поста
	GetAuthorOfPost(ctx context.Context, postID string) (string, error)
	// GetUserLayerOfAuthor - уровень подписки пользователя на определенном авторе
	GetUserLayerOfAuthor(ctx context.Context, userID, authorID string) (int, error)
}
