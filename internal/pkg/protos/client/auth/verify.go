package auth

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	authv1 "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/protos/gen/go/pushart.auth.v1"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type VerifyClient struct {
	grpcClient authv1.AuthClient
	conn       *grpc.ClientConn
}

func NewVerifyClient() (*VerifyClient, error) {
	op := "internal.pkg.protos.client.auth.verify.NewVerifyClient"

	conn, err := GetConn()
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	grpcClient := GetClient(conn)

	return &VerifyClient{grpcClient: grpcClient, conn: conn}, nil
}

func (ac *VerifyClient) VerifyToken(token string) (isLogged bool, userID string, err error) {
	op := "internal.pkg.protos.client.auth.verify.VerifyToken"
	logger.StandardDebugF(context.Background(), op, "Got grps request with token=%v", token)
	response, err := ac.grpcClient.Login(context.Background(), &authv1.LoginRequest{Token: token})
	if err != nil {
		err = errors.Wrap(err, op)
		return false, "", errors.Wrap(err, "call failed")
	}

	logger.StandardDebugF(context.Background(), op, "Got grps response {userID=%v, isLogged=%v}", response.UserId, response.IsLogin)
	return response.IsLogin, response.UserId, nil
}

func (ac *VerifyClient) Close() error {
	return ac.conn.Close()
}

func (ac *VerifyClient) String() string {
	return fmt.Sprintf("VarifyClient{grpcClient=%v, conn=%v}", ac.grpcClient, ac.conn)
}
