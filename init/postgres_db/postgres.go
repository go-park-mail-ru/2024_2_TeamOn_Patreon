package postgres

import (
	"context"

	logger "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	pgx "github.com/jackc/pgx/v4"
)

func InitPostgres() *pgx.Conn {
	op := "internal.account.postgres_db.InitPostgres"

	host := "postgres"
	conString := "postgres://admin:adminpass@" + host + ":5432/testdb"
	// Параметры подключения
	conn, err := pgx.Connect(context.Background(), conString)
	if err != nil {
		logger.StandardDebugF(op, "Unable to connect to database {%v}", err)
		return nil
	}

	// Не забыть закрыть соединение!
	// defer conn.Close(context.Background())

	logger.StandardInfoF(op, "Successful connect to PostgresDB")
	return conn
}
