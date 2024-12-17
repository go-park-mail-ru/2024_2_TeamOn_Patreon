package postgresql

import (
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/repository/postgres"
	_ "github.com/jackc/pgx/v5"
)

type ModerationRepository struct {
	db postgres.DbPool
}

func NewModerationRepository(db postgres.DbPool) *ModerationRepository {
	return &ModerationRepository{db}
}
