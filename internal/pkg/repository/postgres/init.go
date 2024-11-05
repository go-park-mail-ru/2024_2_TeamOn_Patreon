package postgres

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/config"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
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

const (
	//HOST     = "127.0.0.1"
	Host     = "127.0.0.1"
	Port     = "5432"
	User     = "admin"
	Password = "adminpass"
	DbName   = "testdb"

	DbSslMode = "disable"
)

func initConfig() (Config, error) {
	host := config.GetEnv(global.EnvDBHost, Host)
	port := config.GetEnv(global.EnvDBPort, Port)
	user := config.GetEnv(global.EnvDbUser, User)
	password := config.GetEnv(global.EnvDbPassword, Password)
	dbname := config.GetEnv(global.EnvDbName, DbName)
	dbSslMode := config.GetEnv(global.EnvDBSSLMode, DbSslMode)

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

	logger.StandardDebugF(op, "Connecting do db cfg='%v'", cfg)

	connString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBName, cfg.DBPassword, cfg.DBSSLMode)

	logger.StandardDebugF(op, "Connecting do db cfg='%v'", connString)

	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		panic(errors.Wrap(err, op))
	}

	logger.StandardInfoF(op, "Successfully connected to PostgresDB pool=%v", pool)
	return pool
}
