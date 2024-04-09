package main

import (
	"G-PROJECT/handlers"
	"G-PROJECT/middleware"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
    timeoutMiddleware := middleware.TimeoutMiddleware

    http.Handle("/login", timeoutMiddleware(http.HandlerFunc(handlers.LoginHandler)))
    http.Handle("/register", timeoutMiddleware(http.HandlerFunc(handlers.RegisterHandler)))
    http.Handle("/refresh-token", timeoutMiddleware(http.HandlerFunc(handlers.RefreshTokenHandler)))
    http.Handle("/users", timeoutMiddleware(middleware.AuthMiddleware(http.HandlerFunc(handlers.ListUsersHandler))))
    http.Handle("/upload", timeoutMiddleware(middleware.AuthMiddleware(http.HandlerFunc(handlers.UploadImageHandler))))
    http.Handle("/open-image/", timeoutMiddleware(middleware.AuthMiddleware(http.HandlerFunc(handlers.GetImageHandler))))

    server := &http.Server{
        Addr:         "localhost:8080",
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
        IdleTimeout:  15 * time.Second,
    }

    go func() {
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("Server error: %v", err)
        }
    }()

    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    log.Println("Shutting down server...")

    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    if err := server.Shutdown(ctx); err != nil {
        log.Fatalf("Server shutdown error: %v", err)
    }
    log.Println("Server gracefully stopped")
}
