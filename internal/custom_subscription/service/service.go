package service

import (
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/custom_subscription/service/interfaces"
	"github.com/pkg/errors"
)

type Behavior struct {
	rep interfaces.CustomSubscriptionRepository
}

func New(repository interfaces.CustomSubscriptionRepository) *Behavior {
	if repository == nil {
		panic(errors.New("repository is nil"))
	}
	return &Behavior{repository}
}
