package interfaces

import (
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/tree/polina-auth/internal/common/interfaces"
)

// AuthRepository включает оба интерфейса Repository и UserRepository
type AuthRepository interface {
	interfaces.Repository
	UserRepository
}
