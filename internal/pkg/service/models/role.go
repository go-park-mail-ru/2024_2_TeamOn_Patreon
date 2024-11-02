package models

// Role - alias для string
// Определяем тип роли
type Role string

// Именованные константы для ролей
const (
	Reader Role = "Reader" // Читатель
	Author Role = "Author" // Автор
)
