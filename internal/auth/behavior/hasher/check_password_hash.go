package hasher

import "golang.org/x/crypto/bcrypt"

func CheckPasswordHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
