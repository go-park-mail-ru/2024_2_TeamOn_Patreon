package auth

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/config"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	authv1 "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/protos/gen/go/pushart.auth.v1"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

func GetConn() (*grpc.ClientConn, error) {
	op := "internal.pkg.protos.client.auth.auth.GetConn"

	port := os.Getenv(global.EnvGRPCPort)
	address := config.GetEnv(global.EnvGRPCAddress, "localhost")
	addr := address + ":" + port
	logger.StandardDebugF(context.Background(), op, "try grpc connect to %v", addr)
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, errors.Wrap(err, op+" NewClient failed")
	}
	// не забыть закрыть соединение!!!
	// defer conn.Close()
	return conn, nil
}

func GetClient(conn *grpc.ClientConn) authv1.AuthClient {
	client := authv1.NewAuthClient(conn)
	return client
}
