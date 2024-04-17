package handlers

import (
	"encoding/json"
	"net/http"
)

type ErrorType string
type contextKey string



const (
    UNABLE_TO_SAVE         ErrorType = "UNABLE_TO_SAVE"
    UNABLE_TO_FIND_RESOURCE ErrorType = "UNABLE_TO_FIND_RESOURCE"
    UNABLE_TO_READ         ErrorType = "UNABLE_TO_READ"
    UNAUTHORIZED           ErrorType = "UNAUTHORIZED"
)

var ErrorMap = map[ErrorType]int{
    UNABLE_TO_SAVE:         http.StatusInternalServerError,
    UNABLE_TO_FIND_RESOURCE: http.StatusNotFound,
    UNABLE_TO_READ:         http.StatusInternalServerError,
    UNAUTHORIZED:           http.StatusUnauthorized,
}

func ErrorResponse(w http.ResponseWriter, errType ErrorType, message string) {
    statusCode, exists := ErrorMap[errType]
    if !exists {
        statusCode = http.StatusInternalServerError
    }

    w.WriteHeader(statusCode)

    errorResponse := struct {
        Error string `json:"error"`
    }{
        Error: message,
    }

    json.NewEncoder(w).Encode(errorResponse)
}

