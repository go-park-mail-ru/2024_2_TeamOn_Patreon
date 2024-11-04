package models

const (
	AuthorStatus = "Author"
	ReaderStatus = "Reader"
)

// Service модель User
type User struct {
	UserID   string
	Username string
	Email    string // так как может быть пустым
	Role     string
}
