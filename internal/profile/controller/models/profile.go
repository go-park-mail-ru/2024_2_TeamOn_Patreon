package models

const (
	AuthorStatus = "Author"
	ReaderStatus = "Reader"
)

// API Модель профиля пользователя
type Profile struct {
	// Имя пользователя
	Username string `json:"username"`
	// Почта пользователя (если есть)
	Email string `json:"email,omitempty"`
	// Ссылка на фото профиля (если есть)
	AvatarUrl string `json:"avatar_url,omitempty"`
	// Статус "О себе"
	Status string `json:"status"`
	// Роль: читатель или автор
	Role string `json:"role"`
	// Количество подписчиков
	Followers uint `json:"followers"`
	// Количество подписок
	Subscriptions uint `json:"subscriptions"`
	// Количество постов
	PostsAmount uint `json:"posts_amount"`
	// TODO: сделать ProfilePostGet()
	PostTitle   string `json:"posts_title"`
	PostContent string `json:"posts_content"`
	PostDate    string `json:"posts_date"`
}
