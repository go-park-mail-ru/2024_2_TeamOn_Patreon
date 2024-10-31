package models

// Role - alias для int
// Определяем тип роли
type Role string

// Именованные константы для ролей
const (
	Reader Role = "Reader" // Читатель
	Author Role = "Author" // Автор
	Anon   Role = "Anon"   // Инкогнито
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

func StringToRole(role string) Role {
	switch role {
	case "Reader":
		return Reader
	case "Author":
		return Author
	default:
		return ""
	}
}
