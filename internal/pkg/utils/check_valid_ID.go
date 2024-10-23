package utils

import "regexp"

// Функция isValidUUIDv4 проверяет соответствие ID стандарту UUIDv4
func IsValidUUIDv4(uuid string) bool {
	// Регулярное выражение для проверки формата UUID v4
	re := regexp.MustCompile(`^([0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12})$`)
	return re.MatchString(uuid)
}
