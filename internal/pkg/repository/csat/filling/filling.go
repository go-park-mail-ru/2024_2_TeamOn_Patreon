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

	metaQuestion := map[string][]string{consts.QuestionType0: {consts.Question_0_1, consts.Question_0_2, consts.Question_0_3, consts.Question_0_4, consts.Question_0_5, consts.Question_0_6, consts.Question_0_7, consts.Question_0_8, consts.Question_0_9, consts.Question_0_10, consts.Question_11}, consts.QuestionType1: {consts.Question_1_1, consts.Question_1_2, consts.Question_1_3, consts.Question_1_4, consts.Question_1_5, consts.Question_1_6, consts.Question_1_7, consts.Question_1_8, consts.Question_1_9, consts.Question_1_10}, consts.QuestionType2: {consts.Question_2_1, consts.Question_2_2, consts.Question_2_3, consts.Question_2_4, consts.Question_2_5, consts.Question_2_6, consts.Question_2_7, consts.Question_2_8, consts.Question_2_9, consts.Question_2_10}}

	if err := createQuestionType(context.Background(), pool, metaQuestion); err != nil {
		log.Fatalf("Ошибка при создании типов вопросов: %v", err)
	}

	log.Printf("Создано типов %v", len(metaQuestion))

	if err := createQuestion(context.Background(), pool, metaQuestion); err != nil {
		log.Fatalf("Ошибка при создании вопросов %v", err)
	}
	log.Printf("Создано вопросов %v", metaQuestion)

	log.Printf("dbUser: %v dbHost=%v dbPort=%v dbName=%v", dbUser, dbHost, dbPort, dbName)

}

func createQuestionType(ctx context.Context, pool *pgxpool.Pool, meta map[string][]string) error {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("не удалось получить соединение из пула: %v", err)
	}
	defer conn.Release()

	batch := &pgx.Batch{}

	// Подготавливаем данные для пользователей и связанных записей
	for qType, _ := range meta {
		questionTypeID := uuid.NewV4()

		// Запрос на добавление пользователя
		batch.Queue(`
            INSERT INTO question_theme (question_theme_id, theme) 
            VALUES ($1, $2)`,
			questionTypeID, qType)
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

func createQuestion(ctx context.Context, pool *pgxpool.Pool, meta map[string][]string) error {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("не удалось получить соединение из пула: %v", err)
	}
	defer conn.Release()

	batch := &pgx.Batch{}

	// Подготавливаем данные для пользователей и связанных записей
	for qType, questions := range meta {
		for _, question := range questions {
			questionID := uuid.NewV4()

			// Запрос на добавление пользователя
			batch.Queue(`
            INSERT INTO question (question_id, question_theme_id, question) 
            VALUES ($1, (select question_theme_id from question_theme where theme= $2  limit 1), $3)`,
				questionID, qType, question)
		}

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
