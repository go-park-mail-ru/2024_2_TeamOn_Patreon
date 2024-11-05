package postgresql

import (
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/repository/postgres"
	_ "github.com/jackc/pgx/v5"
)

// AuthRepository - имплементирует интерфейс репозитория, используемый
// в бизнес-логике постов (авторизации), используя postgresql
type AuthRepository struct {
	db postgres.DbPool
}

func NewAuthRepository(db postgres.DbPool) *AuthRepository {
	return &AuthRepository{db}
}
