package postgres

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/config"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"

	logger "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	pgxpool "github.com/jackc/pgx/v5/pgxpool"
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

func InitPostgres(ctx context.Context) *pgxpool.Pool {
	op := "internal.account.postgres_db.InitPostgres"

	cfg, err := initConfig()
	if err != nil {
		panic(errors.Wrap(err, op))
	}

	logger.StandardDebugF(ctx, op, "Connecting do db cfg='%v'", cfg)

	connString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBName, cfg.DBPassword, cfg.DBSSLMode)

	// Параметры подключения
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		logger.StandardDebugF(ctx, op, "Unable to connect to database {%v}", err)
		panic(errors.Wrap(err, op))
	}

	logger.StandardInfoF(ctx, op, "Successful connect to PostgresDB")
	return pool
}
