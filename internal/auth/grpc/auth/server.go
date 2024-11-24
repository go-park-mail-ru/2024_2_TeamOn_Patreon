package auth

import (
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/auth/grpc/auth/interfaces"
	"google.golang.org/grpc"

	_ "google.golang.org/grpc"

	authv1 "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/protos/gen/go/pushart.auth.v1"
)

// ServerAPI - реализовывает функционал API
type ServerAPI struct {
	authv1.UnimplementedAuthServer // вложенная структура
	// сгенерированная protoc на основе прото-файла
	// пустая имплементация методов grpc сервиса

	// благодаря нему при повторной генерации
	// если мы добавили новый метод в protobuf
	// код будет компилироваться
	// просто при вызове
	// не имплементированных методов будет соответствующая ошибка

	auth interfaces.VerifyToken
}

// Register регистрирует serverAPI в gRPC-сервере
func Register(gRPCServer *grpc.Server, auth interfaces.VerifyToken) {
	authv1.RegisterAuthServer(gRPCServer, &ServerAPI{auth: auth})
}
