package models

// User бизнес-модель пользователя
type User struct {
	UserID   string
	Username string
	Role     Role
}
