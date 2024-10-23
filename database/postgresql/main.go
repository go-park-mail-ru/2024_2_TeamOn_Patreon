package main

import (
	"context"
	"fmt"
	"log"

	pgx "github.com/jackc/pgx/v4"
)

func main() {
	// Параметры подключения
	conn, err := pgx.Connect(context.Background(), "postgres://admin:adminpass@localhost:5432/testdb")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	// Создание таблицы
	_, err = conn.Exec(context.Background(), `
        CREATE TABLE IF NOT EXISTS messages (
            id SERIAL PRIMARY KEY,
            text VARCHAR(255)
        );
    `)
	if err != nil {
		log.Fatalf("Unable to create table: %v\n", err)
	}

	// Добавление строки
	_, err = conn.Exec(context.Background(), "INSERT INTO messages (text) VALUES ('Hello, PostgreSQL!');")
	// Очистить таблицу messages:
	// _, err = conn.Exec(context.Background(), "TRUNCATE TABLE messages;")

	if err != nil {
		log.Fatalf("Unable to insert row: %v\n", err)
	}

	// Вывод строки
	rows, err := conn.Query(context.Background(), "SELECT * FROM messages;")
	if err != nil {
		log.Fatalf("Unable to query rows: %v\n", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var text string
		err = rows.Scan(&id, &text)
		if err != nil {
			log.Fatalf("Unable to scan row: %v\n", err)
		}
		fmt.Printf("ID: %d, Text: %s\n", id, text)
	}
}
