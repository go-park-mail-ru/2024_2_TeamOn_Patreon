package service

import (
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
