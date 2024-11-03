package postgres

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBName     string
	DBPassword string
	DBSSLMode  string
}

var (
	//host     = "127.0.0.1"
	host     = "postgres"
	port     = "5432"
	user     = "admin"
	password = "adminpass"
	dbname   = "testdb"

	dbSslMode = "disable"
)

func initConfig() (Config, error) {
	cfg := Config{
		DBHost:     host,
		DBPort:     port,
		DBUser:     user,
		DBName:     dbname,
		DBPassword: password,
		DBSSLMode:  dbSslMode,
	}

	return cfg, nil

}

// InitPostgresDB - возвращает pool для создания запросов.
// Надо использовать в main и передавать во всю бизнес-логику
// Использовать стоит ограниченное число пулов (наверно, ок и один на весь демон)
func InitPostgresDB(ctx context.Context) *pgxpool.Pool {
	op := "internal.pkg.repository.postgres.init.InitPostgresDB"

	cfg, err := initConfig()
	if err != nil {
		panic(errors.Wrap(err, op))
	}

	logger.StandardDebugF(op, "Connecting do db cfg='%v'")

	connString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBName, cfg.DBPassword, cfg.DBSSLMode)

	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		panic(errors.Wrap(err, op))
	}

	logger.StandardInfoF(op, "Successfully connected to PostgresDB pool=%v", pool)
	return pool
}
