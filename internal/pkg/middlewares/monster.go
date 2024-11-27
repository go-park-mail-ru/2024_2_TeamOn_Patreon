package middlewares

import (
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/protos/client/auth"
	"github.com/pkg/errors"
)

// получается
// подключаем монстра
// когда закончили - monster.Close()

type Monster struct {
	client AuthClient
}

type AuthClient interface {
	VerifyToken(token string) (isLogged bool, userID string, err error)
	Close() error
}

func NewMonster() *Monster {
	op := "internal.pkg.middlewares.NewMonster"
	m := &Monster{}

	client, err := auth.NewVerifyClient()
	if err != nil {
		panic(errors.Wrap(err, op))
	}

	m.client = client
	logger.StandardDebugF(nil, op, "Successfully created new Monster with client=%v", client)
	return m
}

func (m *Monster) Close() error {
	return m.client.Close()
}
