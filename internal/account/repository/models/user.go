package models

// Repository модель User
type User struct {
	UserID   string `db:"user_id"`
	Username string `db:"username"`
	Email    string `db:"email"`
	RoleID   int    `db:"role_id"`
}

// func MapRepUserToServUser(user User) sModels.User {
// 	return sModels.Account{
// 		UserID:   user.UserID,
// 		Username: user.Username,
// 		Email:    user.Email,
// 		// Role: user.Role
// 	}
// }
