package middleware

import (
	"context"
	"net/http"
	"time"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
        requestTimeout := 10* time.Second

		ctx, cancel := context.WithTimeout(r.Context(), requestTimeout)
		defer cancel()

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
