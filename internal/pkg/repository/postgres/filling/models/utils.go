package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func getAuthors(jsonPath string) []*FillingAuthor {
	// Открытие файла
	file, err := os.Open(jsonPath)
	if err != nil {
		log.Fatalf("Ошибка при открытии файла: %v", err)
	}
	defer file.Close()

	// Чтение содержимого файла
	fileData, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Ошибка при чтении файла: %v", err)
	}

	// Слайс структур Author
	var authors []*FillingAuthor

	// Распаковка JSON
	err = json.Unmarshal(fileData, &authors)
	if err != nil {
		log.Fatalf("Ошибка декодирования JSON: %v", err)
	}

	return authors
}
