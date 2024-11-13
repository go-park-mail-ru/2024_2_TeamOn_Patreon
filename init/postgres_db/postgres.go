package postgres

import (
	"context"

	logger "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	pgxpool "github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

func InitPostgres(ctx context.Context) *pgxpool.Pool {
	op := "internal.account.postgres_db.InitPostgres"

	host := "postgres"
	conString := "postgres://admin:adminpass@" + host + ":5432/testdb"
	// Параметры подключения
	pool, err := pgxpool.New(ctx, conString)
	if err != nil {
		logger.StandardDebugF(ctx, op, "Unable to connect to database {%v}", err)
		panic(errors.Wrap(err, op))
	}

	logger.StandardInfoF(ctx, op, "Successful connect to PostgresDB")
	return pool
}
