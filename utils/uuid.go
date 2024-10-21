package utils

import (
	// "fmt"

	"github.com/google/uuid"
)

func GenerateUUID() string {
	return uuid.New().String()
}
