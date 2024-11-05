package models

import "database/sql"

// Repository модель User
type User struct {
	UserID   string         `db:"user_id"`
	Username string         `db:"username"`
	Email    sql.NullString `db:"email"`
	Role     string         `db:"role_default_name"`
}
