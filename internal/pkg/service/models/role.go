package models

// Role - alias для int
// Определяем тип роли
type Role string

// Именованные константы для ролей
const (
	Reader Role = "1" // Читатель
	Author Role = "2" // Автор
)

// RoleToString
// Функция для отображения роли в виде строки
func RoleToString(role Role) string {
	switch role {
	case Reader:
		return "Reader"
	case Author:
		return "Author"
	default:
		return "Unknown"
	}
}
