package models

// Role - alias для string
// Определяем тип роли
type Role string

// Именованные константы для ролей
const (
	Reader    Role = "Reader" // Читатель
	Author    Role = "Author" // Автор
	Moderator Role = "Moderator"
	Anon      Role = "Anon" // Инкогнито
)

func StringToRole(s string) Role {
	switch s {
	case "Reader":
		return Reader
	case "Author":
		return Author
	case "Moderator":
		return Moderator
	default:
		return Anon
	}
}
