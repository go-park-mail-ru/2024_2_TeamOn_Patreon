package models
)
type Reg struct {
	// Имя пользователя. допустимые символы - латинские буквы, цифры и \"-\", \"_\".
	Username string `json:"username"`
	// Пароль. должен содержать хотя бы 1 заглавную, 1 строчную латинские буквы, 1 цифру, 1 спец символ.
	Password string `json:"password"`
}
