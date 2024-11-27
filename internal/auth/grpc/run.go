package grpc

import (
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/grpc/auth"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/grpc/auth/interfaces"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/pkg/errors"
	"os"
	"strconv"
)

func MustRun(beh interfaces.VerifyToken) {
	_, err := Run(beh)
	if err != nil {
		panic(err)
	}
}

func Run(beh interfaces.VerifyToken) (*auth.App, error) {
	op := "auth.controller.grpc.Run"

	port, err := strconv.Atoi(os.Getenv(global.EnvGRPCPort))
	if err != nil {
		err = errors.Wrap(err, op)
		return nil, errors.Wrap(err, "port is invalid")
	}

	grpcApp := auth.New(beh, port)

	logger.StandardDebugF(nil, op, "started grpc app")
	err = grpcApp.Run()
	if err != nil {
		return nil, errors.Wrap(err, op)
	}
	return grpcApp, nil
}
