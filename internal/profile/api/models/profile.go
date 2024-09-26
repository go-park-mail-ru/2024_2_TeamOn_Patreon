package models

const (
	AuthorStatus = "author"
	ReaderStatus = "reader"
)

// Модель профиля пользователя
type Profile struct {
	// Имя пользователя
	Username string `json:"username"`
	// Почта пользователя (если есть)
	Email string `json:"email,omitempty"`
	// Ссылка на фото профиля (если есть)
	AvatarUrl string `json:"avatar_url,omitempty"`
	// Статус: читатель или автор
	Status string
	// Количество подписчиков
	Followers int32 `json:"followers,omitempty"`
	// Количество подписок
	Subscriptions int32 `json:"subscriptions,omitempty"`
	// Количество постов
	Posts int32 `json:"posts,omitempty"`
}
