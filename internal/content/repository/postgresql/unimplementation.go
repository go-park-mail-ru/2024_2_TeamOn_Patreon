package postgresql

import (
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/pkg/models"
	"github.com/gofrs/uuid"
)

func (cr *ContentRepository) GetPopularPosts(offset int, limits int) ([]models.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (cr *ContentRepository) GetPopularPostsForUser(userId uuid.UUID, offset int, limits int) ([]models.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (cr *ContentRepository) GetAuthorByPost(postID uuid.UUID) (uuid.UUID, error) {
	//TODO implement me
	panic("implement me")
}

func (cr *ContentRepository) IsLikePutPost(userId uuid.UUID, postID uuid.UUID) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (cr *ContentRepository) InsertLikePost(userId uuid.UUID, postID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (cr *ContentRepository) DeleteLikePost(userId uuid.UUID, postID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (cr *ContentRepository) GetPostLikes(postID uuid.UUID) (int, error) {
	//TODO implement me
	panic("implement me")
}
