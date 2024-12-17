package global

import (
	"log/slog"
)

// Cookie JWT

const CookieJWT = "JWT"
const TTL = 24

// logger

const LoggerLevel = slog.LevelDebug

// From context

// BehaviorKey - ключ для получения бизнес-логики из контекста
const BehaviorKey = "service"

// UserKey - ключ для получения бизнес-модельки юзера из контекста
const UserKey string = "user"

// CtxReqId Константа для пролучения reqID из контекста
const CtxReqId = "request_id"

// ROLEs - лежат в общих моделях

// statistic
const (
	StatPosts    = "posts"
	StatPayments = "payments"
)

// ENV
const (
	EnvStatus = "STATUS"

	EnvInDocker   = "IN_DOCKER"
	EnvDBHost     = "DB_HOST"
	EnvDBPort     = "DB_PORT"
	EnvDbUser     = "DB_USER"
	EnvDbPassword = "DB_PASSWORD"
	EnvDbName     = "DB_NAME"
	EnvDBSSLMode  = "DB_SSL_MODE"

	EnvLogLevel = "LOG_LEVEL"

	EnvServiceName = "SERVICE_NAME"
	EnvJWTKey      = "JWT_KEY"
	EnvPort        = "SERVICE_PORT"

	EnvJWT      = "JWT_KEY"
	EnvTokenTTL = "TOKEN_TTL"

	// grpc

	EnvGRPCPort    = "GRPC_PORT"
	EnvGRPCTimeout = "GRPC_TIMEOUT"
	EnvGRPCAddress = "GRPC_ADDRESS"

	// API Payment Service

	EnvClientID  = "CLIENT_ID"
	EnvSecretKey = "SECRET_KEY"
)
