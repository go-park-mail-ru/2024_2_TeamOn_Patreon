package service

import (
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/pkg/models"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/service/interfaces"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/gofrs/uuid"
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
	authorIdUuid, err := uuid.FromString(authorId)
	if err != nil {
		return nil, errors.Wrap(global.ErrServer, op)
	}
	_ = authorIdUuid
	return nil, err
}

func (b *Behavior) GetAuthorPostsForUser(authorId, userId string, offset, limit int) ([]models.Post, error) {
	op := "service.behavior.GetAuthorPostsForUser"
	authorIdUuid, err := uuid.FromString(authorId)
	if err != nil {
		return nil, errors.Wrap(global.ErrServer, op)
	}
	_ = authorIdUuid
	return nil, err
}

func (b *Behavior) GetAuthorPostsForMe(authorId string, offset, limit int) ([]models.Post, error) {
	op := "service.behavior.GetAuthorPostsForMe"
	authorIdUuid, err := uuid.FromString(authorId)
	if err != nil {
		return nil, errors.Wrap(global.ErrServer, op)
	}
	_ = authorIdUuid
	return nil, err
}
