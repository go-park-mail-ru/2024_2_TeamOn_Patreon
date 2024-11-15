package utils

import (
	satori "github.com/satori/go.uuid"
)

// isValidUUIDv4 проверяет соответствие ID стандарту UUIDv4
func IsValidUUIDv4(uuid string) bool {

	u, err := satori.FromString(uuid)
	if err != nil {
		return false
	}

	// 1: UUID версии 1 (на основе времени).
	// 2: UUID версии 2 (DCE Security).
	// 3: UUID версии 3 (на основе хэша MD5).
	// 4: UUID версии 4 (случайные значения).
	// 5: UUID версии 5 (на основе хэша SHA-1).

	if u.Version() != 4 {
		return false
	}

	return true
}
