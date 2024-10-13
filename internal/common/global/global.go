package global

import (
	"log/slog"
)

// Cookie JWT

const CookieJWT = "JWT"
const TTL = 24

// Cookie Session

const CookieSession = "Session"
const SessionTimeH = 24

// logger

const LoggerLevel = slog.LevelDebug

// From context

// BehaviorKey - ключ для получения бизнес-логики из контекста
const BehaviorKey = "behavior"

// UserKey - ключ для получения бизнес-модельки юзера из контекста
const UserKey string = "user"
