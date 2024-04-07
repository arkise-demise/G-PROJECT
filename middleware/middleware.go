package middleware

import (
	"G-PROJECT/utils"
	"context"
	"net/http"
	"time"
)

type contextKey string

const (
	requestIDKey contextKey = "requestID"
	userIDKey    contextKey = "userID"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestTimeout := 3 * time.Second
		ctx, cancel := context.WithTimeout(r.Context(), requestTimeout)
		defer cancel()

		requestID := utils.GenerateUUID()
		ctx = context.WithValue(ctx, requestIDKey, requestID)

		userID := extractUserIDFromRequest(r)
		if userID != "" {
			ctx = context.WithValue(ctx, userIDKey, userID)
		}

		r = r.WithContext(ctx)
		

		time.Sleep(7*time.Second)
		next.ServeHTTP(w, r)
	})
}

func extractUserIDFromRequest(r *http.Request) string {
	userID := r.Header.Get("UserID")
	return userID
}
