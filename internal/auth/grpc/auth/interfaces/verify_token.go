package interfaces

import "context"

// VerifyToken - интерфейс сервисного слоя для рпс
type VerifyToken interface {
	VerifyToken(
		ctx context.Context,
		token string,
	) (isLogged bool, userID string, err error)
}
