package models

// Role - alias для int
// Определяем тип роли
type Role int

// Именованные константы для ролей
const (
	Reader Role = iota + 1 // Читатель
	Author                 // Автор
)
