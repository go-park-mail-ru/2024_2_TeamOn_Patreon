package service

import (
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/content/service/interfaces"
)

type Behavior struct {
	rep interfaces.ContentRepository
}

func New(repository interfaces.ContentRepository) *Behavior {
	return &Behavior{repository}
}
