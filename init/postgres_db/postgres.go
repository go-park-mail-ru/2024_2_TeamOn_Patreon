package postgres

import (
	"context"

	logger "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	pgxpool "github.com/jackc/pgx/v4/pgxpool"
)

func InitPostgres() *pgxpool.Pool {
	op := "internal.account.postgres_db.InitPostgres"

	host := "postgres"
	conString := "postgres://admin:adminpass@" + host + ":5432/testdb"
	// Параметры подключения
	pool, err := pgxpool.Connect(context.Background(), conString)
	if err != nil {
		logger.StandardDebugF(op, "Unable to connect to database {%v}", err)
		return nil
	}

	logger.StandardInfoF(op, "Successful connect to PostgresDB")
	return pool
}
