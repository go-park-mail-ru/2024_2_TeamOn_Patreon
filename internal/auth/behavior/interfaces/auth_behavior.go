package interfaces

type AuthBehavior interface {
	// RegisterNewUser - регистрация | добавление нового пользователя, генерация для него jwt
	//RegisterNewUser(username string, password string) (bJWT.TokenString, error)
	RegisterNewUser(username string, password string) (string, error)

	// AuthoriseUser - авторизация | проверяет существует ли пользователь, верный ли пароль, генерирует jwt для него
	//AuthoriseUser(username string, password string) (bJWT.TokenString, error)
	AuthoriseUser(username string, password string) (string, error)
}
