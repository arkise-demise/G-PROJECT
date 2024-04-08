package utils

import "github.com/google/uuid"

func GenerateRequestID() string {
    return uuid.New().String()
}
