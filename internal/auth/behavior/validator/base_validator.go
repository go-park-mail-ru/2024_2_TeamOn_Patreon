package validator

import "regexp"

type BaseValidator struct {
}

var validatorInstance *BaseValidator

func New() *BaseValidator {
	return &BaseValidator{}
}

func Get() *BaseValidator {
	if validatorInstance == nil {
		validatorInstance = New()
	}
	return validatorInstance
}

func (val *BaseValidator) isGreaterThan(field string, n int) bool {
	return len(field) > n
}

func (val *BaseValidator) isLessThan(field string, n int) bool {
	return len(field) < n
}

func (val *BaseValidator) checkPattern(field string, pattern string) bool {
	re := regexp.MustCompile(pattern)
	return re.MatchString(field)
}
