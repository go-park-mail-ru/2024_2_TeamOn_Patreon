package postgresql

import (
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/repository/postgres"
	_ "github.com/jackc/pgx/v5"
)

type CustomSubscriptionRepository struct {
	db postgres.DbPool
}

func NewCustomSubscriptionRepository(db postgres.DbPool) *CustomSubscriptionRepository {
	return &CustomSubscriptionRepository{db}
}
