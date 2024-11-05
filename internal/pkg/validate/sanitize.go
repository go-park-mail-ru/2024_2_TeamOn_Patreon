package validate

import "github.com/microcosm-cc/bluemonday"

var sanitizer = bluemonday.UGCPolicy()

func Sanitize(value string) string {
	value = sanitizer.Sanitize(value)
	return value
}
