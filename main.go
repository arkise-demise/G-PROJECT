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
	requestTimeout := 1 * time.Minute

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/refresh-token", handlers.RefreshTokenHandler)
	http.Handle("/users", middleware.AuthMiddleware(http.HandlerFunc(handlers.ListUsersHandler)))
	http.Handle("/upload", middleware.AuthMiddleware(http.HandlerFunc(handlers.UploadImageHandler)))
	http.Handle("/open-image/", middleware.AuthMiddleware(http.HandlerFunc(handlers.GetImageHandler)))

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

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-sigCh:
		cancel()
		log.Printf("Received signal: %v", sig)
	case <-ctx.Done():
		log.Println("Context timed out")
	}

	shutdownCtx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}
}
