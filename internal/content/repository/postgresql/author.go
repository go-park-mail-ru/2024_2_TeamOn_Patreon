package postgresql

import "github.com/gofrs/uuid"

// GetAuthorByPostId возвращает Id автора по одному Id его поста
func (cr *ContentRepository) GetAuthorByPostId() error {
	return nil
}

// GetAuthorInfoByAuthorId возвращает информацию об авторе по его Id
func (cr *ContentRepository) GetAuthorInfoByAuthorId() error {
	return nil
}

// GetAuthorPostsByLayer - возвращает посты автора с offset по (offset + limit)
func (cr *ContentRepository) GetAuthorPostsByLayer(authorId uuid.UUID, layerId uuid.UUID, offset int, limit int) error {
	return nil
}
