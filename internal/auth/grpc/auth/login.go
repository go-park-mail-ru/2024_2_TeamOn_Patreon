package auth

import (
	"context"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	authv1 "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/protos/gen/go/pushart.auth.v1"

	// rpc errors
	"google.golang.org/grpc/codes" // коды ошибок
	"google.golang.org/grpc/status"
)

// Реализация grpc метода
// Верификации токена

// Login - rpc метод
func (s *ServerAPI) Login(
	ctx context.Context,
	in *authv1.LoginRequest,
) (*authv1.LoginResponse, error) {
	op := "auth.controller.grpc.auth.login"

	logger.StandardDebugF(ctx, op, "got in=%v", in)

	isLogged, userID, err := s.auth.VerifyToken(ctx, in.Token)
	if err != nil {
		logger.StandardDebugF(ctx, op, "From verify token err=%v", err)
		return nil, status.Error(codes.Aborted, "not verify")
	}
	out := &authv1.LoginResponse{
		IsLogin: isLogged,
		UserId:  userID,
	}

	return out, nil
}
