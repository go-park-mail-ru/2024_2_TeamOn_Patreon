package models

// User бизнес-модель пользователя
type User struct {
	UserID   UserID
	Username string
	Role     Role
}
