package models

// Repository модель User
type User struct {
	UserID   string `db:"user_id"`
	Username string `db:"username"`
	Email    string `db:"email"`
	RoleID   int    `db:"role_id"`
}
