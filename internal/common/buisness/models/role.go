package models

// Role - alias для int
// Определяем тип роли
type Role int

// Именованные константы для ролей
const (
	Reader Role = iota + 1 // Читатель
	Author                 // Автор
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
