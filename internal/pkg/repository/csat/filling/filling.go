package main

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/repository/csat/filling/consts"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	uuid "github.com/satori/go.uuid"
	"log"
)

func main() {
	// Укажите, сколько пользователей нужно создать
	nType := consts.QuestionTypeCount
	nQuestions := consts.QuestionForTypeCount

	// Параметры подключения к PostgreSQL
	dbHost := "postgres-csat"
	dbPort := 5432
	dbUser := "admin"
	dbPassword := "adminpass"
	dbName := "csatdb"

	log.Printf("dbUser: %v dbHost=%v dbPort=%v dbName=%v", dbUser, dbHost, dbPort, dbName)

	// Формируем строку подключения
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)

	//dbURL := "postgres://your_user:your_password@localhost:5432/your_dbname?sslmode=disable"
	pool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Ошибка создания пула подключений: %v", err)
	}
	defer pool.Close()

	if err := createQuestionType(context.Background(), pool, nType); err != nil {
		log.Fatalf("Ошибка при создании типов вопросов: %v", err)
	}

	log.Printf("Создано типов %v", nType)

	if err := createQuestion(context.Background(), pool, nQuestions*nType); err != nil {
		log.Fatalf("Ошибка при создании вопросов %v", err)
	}
	log.Printf("Создано вопросов %v", nQuestions)

	log.Printf("dbUser: %v dbHost=%v dbPort=%v dbName=%v", dbUser, dbHost, dbPort, dbName)

}

func createQuestionType(ctx context.Context, pool *pgxpool.Pool, n int) error {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("не удалось получить соединение из пула: %v", err)
	}
	defer conn.Release()

	batch := &pgx.Batch{}

	// Подготавливаем данные для пользователей и связанных записей
	for i := 0; i < n; i++ {
		questionTypeID := uuid.NewV4()
		questionType := consts.QuestionType

		// Запрос на добавление пользователя
		batch.Queue(`
            INSERT INTO question_theme (question_theme_id, theme) 
            VALUES ($1, $2)`,
			questionTypeID, questionType)

	}

	// Выполнение батч-запроса
	br := conn.Conn().SendBatch(ctx, batch)
	defer br.Close()

	// Проверка результатов выполнения батча
	for i := 0; i < batch.Len(); i++ {
		_, err := br.Exec()
		if err != nil {
			return fmt.Errorf("ошибка выполнения батч-запроса на шаге %d: %v", i+1, err)
		}
	}

	return nil
}

func createQuestion(ctx context.Context, pool *pgxpool.Pool, n int) error {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("не удалось получить соединение из пула: %v", err)
	}
	defer conn.Release()

	batch := &pgx.Batch{}

	// Подготавливаем данные для пользователей и связанных записей
	for i := 0; i < n; i++ {
		questionID := uuid.NewV4()
		questionType := consts.QuestionType
		question := consts.Question

		// Запрос на добавление пользователя
		batch.Queue(`
            INSERT INTO question (question_id, question_theme_id, question) 
            VALUES ($1, (select question_theme_id from question_theme where theme= $2  limit 1), $3)`,
			questionID, questionType, question)

	}

	// Выполнение батч-запроса
	br := conn.Conn().SendBatch(ctx, batch)
	defer br.Close()

	// Проверка результатов выполнения батча
	for i := 0; i < batch.Len(); i++ {
		_, err := br.Exec()
		if err != nil {
			return fmt.Errorf("ошибка выполнения батч-запроса на шаге %d: %v", i+1, err)
		}
	}

	return nil
}
