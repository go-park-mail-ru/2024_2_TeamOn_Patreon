package models

// Model - интерфейс всех моделей фронта
type Model interface {
	Validate() (bool, error)
}
