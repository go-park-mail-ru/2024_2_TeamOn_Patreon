package models

// Model - интерфейс всех моделей фронта
type Model interface {
	Validate() (bool, error)
}

// check - просто структурка, которая используется в ЭТОМ пакете
// нужна для удобного хранения паттернов для регулярок и сообщений о них вместе
type check struct {
	pattern string
	msg     string
}
