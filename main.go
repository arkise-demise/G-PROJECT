package main

import (
	"G-PROJECT/handlers"
	"G-PROJECT/middleware"
	"net/http"
)

func main() {
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.Handle("/users", middleware.AuthMiddleware(http.HandlerFunc(handlers.ListUsersHandler)))
	http.Handle("/upload", middleware.AuthMiddleware(http.HandlerFunc(handlers.UploadImageHandler)))
	http.Handle("/image/", middleware.AuthMiddleware(http.HandlerFunc(handlers.GetImageHandler)))

	err := http.ListenAndServe("localhost:8000", nil)
	if err != nil {
		panic(err)
	}
}
