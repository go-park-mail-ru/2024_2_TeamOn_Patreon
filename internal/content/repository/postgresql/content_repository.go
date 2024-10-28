package postgresql

import (
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/repository/postgres"
	_ "github.com/jackc/pgx/v5"
)

// ContentRepository - имплементирует интерфейс репозитория, используемый
// в бизнес-логике постов (контента), используя postgresql
type ContentRepository struct {
	db postgres.DbPool
}
