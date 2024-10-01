package models

// User бизнес-модель пользователя
type User struct {
	UserID   int
	Username string
	Role     Role
}
