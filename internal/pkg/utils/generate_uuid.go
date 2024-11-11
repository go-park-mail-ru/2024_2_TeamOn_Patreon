package utils

import "github.com/satori/go.uuid"

func GenerateUUID() string {
	return uuid.NewV4().String()
}
