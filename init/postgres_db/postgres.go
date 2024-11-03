package postgres

import (
	"context"

	logger "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	pgx "github.com/jackc/pgx/v4"
)

func InitPostgres() *pgx.Conn {
	op := "internal.account.postgres_db.InitPostgres"
	// Параметры подключения
	conn, err := pgx.Connect(context.Background(), "postgres://admin:adminpass@localhost:5432/testdb")
	if err != nil {
		logger.StandardDebugF(op, "Unable to connect to database {%v}", err)
		return nil
	}

	// Не забыть закрыть соединение!
	// defer conn.Close(context.Background())

	logger.StandardInfoF(op, "Successful connect to PostgresDB")
	return conn
}
