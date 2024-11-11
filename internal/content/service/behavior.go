package service

import (
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/pkg/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/service/interfaces"
	"github.com/pkg/errors"
)

type Behavior struct {
	rep interfaces.ContentRepository
}

func New(repository interfaces.ContentRepository) *Behavior {
	if repository == nil {
		panic(errors.New("repository is nil"))
	}
	return &Behavior{repository}
}

func (b *Behavior) GetAuthorPostsForAnon(authorId string, offset, limit int) ([]models.Post, error) {
	op := "service.behavior.GetAuthorPostsForAnon"
	_ = op
	return nil, nil
}

func (b *Behavior) GetAuthorPostsForUser(authorId, userId string, offset, limit int) ([]models.Post, error) {
	op := "service.behavior.GetAuthorPostsForLayer"
	_ = op
	return nil, nil
}

func (b *Behavior) GetAuthorPostsForMe(authorId string, offset, limit int) ([]models.Post, error) {
	op := "service.behavior.GetAuthorPostsForMe"
	_ = op
	return nil, nil
}
