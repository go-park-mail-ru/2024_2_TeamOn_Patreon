package validate

import "github.com/microcosm-cc/bluemonday"

var sanitizer = bluemonday.UGCPolicy()

// Sanitize - защищает от XSS
// Применяем и при получении и при отправки текста
func Sanitize(value string) string {
	value = sanitizer.Sanitize(value)
	return value
}
