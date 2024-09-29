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
	Followers int32 `json:"followers"`
	// Количество подписок
	Subscriptions int32 `json:"subscriptions"`
	// Количество постов
	PostsAmount int32 `json:"posts_amount"`
}
