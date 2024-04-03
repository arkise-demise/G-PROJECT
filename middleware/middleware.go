package middleware

import (
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			panic(err)
		}

		tokenString := cookie.Value

		if tokenString == "" {
			w.Write([]byte("unauthorized"))
		}

		next.ServeHTTP(w, r)
	})
}
