package middleware

import (
	"G-PROJECT/utils"
	"context"
	"net/http"
)

const (
	requestIDKey contextKey = "requestID"
	userIDKey    contextKey = "userID"
	tokenCookieName = "token" 
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := utils.GenerateRequestID()
		ctx := context.WithValue(r.Context(), requestIDKey, requestID)

		tokenCookie, err := r.Cookie(tokenCookieName)
		if err != nil {
			http.Error(w, "Unauthorized: No token provided", http.StatusUnauthorized)
			return
		}

		user, err := utils.VerifyToken(tokenCookie.Value)
		if err != nil {
			http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
			return
		}

		ctx = context.WithValue(ctx, userIDKey, user.ID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
