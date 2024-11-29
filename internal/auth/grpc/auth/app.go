package auth

import (
	"fmt"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/grpc/auth/interfaces"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	authv1 "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/protos/gen/go/pushart.auth.v1"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
)

type App struct {
	gRPCServer *grpc.Server
	port       int // Порт на котором будет работать grpc сервер
}

func New(auth interfaces.VerifyToken, port int) *App {
	op := "auth.controller.grpc.app.New"
	// создать gRPCServer и подключить к нему интерсепторы
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor( // обертка над инспекторами для запуска поочередно
			// инспектор для обработки паники
			recovery.UnaryServerInterceptor(getRecoveryOpts()...)),
	)
	logger.StandardDebugF(nil, op, "created new grpc server")

	// регистрация у сервера наш gRPC-сервис Auth
	authv1.RegisterAuthServer(grpcServer, &ServerAPI{auth: auth})
	logger.StandardDebugF(nil, op, "registred our server api")

	return &App{
		gRPCServer: grpcServer,
		port:       port,
	}
}

func getRecoveryOpts() []recovery.Option {
	recoveryOpts := []recovery.Option{
		recovery.WithRecoveryHandler(func(p interface{}) (err error) {
			// Логируем информацию о панике
			logger.StandardDebugF(nil, "intersceptopr", "Recovered from panic=%v", p)

			//  "internal error",  не хотим делиться внутренностями
			return status.Errorf(codes.Internal, "internal error")
		}),
	}
	return recoveryOpts
}

// MustRun runs gRPC server and panics if any error occurs.
func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

// Run runs gRPC server.
func (a *App) Run() error {
	const op = "auth.controller.grpc.app.Run"

	// Создаём listener, который будет слушить TCP-сообщения, адресованные
	// Нашему gRPC-серверу
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		err = errors.Wrap(err, op)
		return errors.Wrap(err, "net.listen")
	}

	logger.StandardInfoF(nil, op, "grpc server started addr=%v", l.Addr().String())

	// Запускаем обработчик gRPC-сообщений
	if err := a.gRPCServer.Serve(l); err != nil {
		err = errors.Wrap(err, op)
		return errors.Wrap(err, "started grpc")
	}
	logger.StandardDebugF(nil, op, "started grpc app serve")

	return nil
}
