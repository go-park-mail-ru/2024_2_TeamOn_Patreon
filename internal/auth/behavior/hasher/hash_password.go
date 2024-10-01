package hasher

import "golang.org/x/crypto/bcrypt"

// HashPassword хэширует пароль с солью с помощью библиотеки bcrypt
func HashPassword(password string) (string, error) {
	// Хэширование пароля с заданным уровнем сложности
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
