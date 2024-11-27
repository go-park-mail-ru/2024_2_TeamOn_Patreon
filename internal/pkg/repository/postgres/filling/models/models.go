package models

/*
	Для дефолтного заполнения БД
*/

// Subscription Структура для подписки
type Subscription struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Cost        int    `json:"cost"`
	Layer       int    `json:"layer"`
}

// Post Структура для постов
type Post struct {
	Title         string `json:"title"`
	Info          string `json:"info"`
	MediaFilename string `json:"mediaFilename"`
	Layer         int    `json:"layer"`
}

// FillingAuthor Основная структура для каждого автора
type FillingAuthor struct {
	AuthorName         string         `json:"authorName"`
	Password           string         `json:"password"`
	CustomSubscription []Subscription `json:"customSubscription"`
	Posts              []Post         `json:"posts"`
}
