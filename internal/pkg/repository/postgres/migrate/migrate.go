package main

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/config"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/golang-migrate/migrate/v4"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	// Параметры подключения к PostgreSQL
	dbHost := config.GetEnv(global.EnvDBHost, "127.0.0.1")
	dbPort, _ := strconv.Atoi(config.GetEnv(global.EnvPort, "5432"))
	dbUser := config.GetEnv(global.EnvDbUser, "admin")
	dbPassword := config.GetEnv(global.EnvDbPassword, "adminpass")
	dbName := config.GetEnv(global.EnvDbName, "testdb")

	migrationsDir := "file://database/migrations" // Путь к папке с миграциями

	wd, _ := os.Getwd()
	log.Println("Current working directory:", wd)

	// Формируем строку подключения
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Настраиваем подключение через pgxpool
	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		log.Fatalf("Ошибка парсинга строки подключения: %v", err)
	}
	config.MaxConns = 10 // Настройка пула подключений
	config.HealthCheckPeriod = 5 * time.Second

	// Создаем пул подключений
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Ошибка создания пула подключений: %v", err)
	}
	defer pool.Close()

	// Проверка подключения
	if err := pool.Ping(context.Background()); err != nil {
		log.Fatalf("Не удалось подключиться к БД: %v", err)
	}
	fmt.Println("Подключение к БД успешно установлено")

	// Запуск миграций
	m, err := migrate.New(migrationsDir, dbURL)
	if err != nil {
		log.Fatalf("Ошибка создания migrate: %v", err)
	}

	// Применяем все доступные миграции
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Ошибка применения миграций: %v", err)
	}

	fmt.Println("Миграции успешно применены")
}
