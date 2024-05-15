package main

import (
	"G-PROJECT/database"
	"G-PROJECT/handlers"
	"G-PROJECT/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	dbInstance, err := database.NewDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer dbInstance.Close() 

	r := gin.Default()

	r.POST("/login", middleware.TimeoutMiddleware(), middleware.ErrorHandlerMiddleware(), handlers.LoginHandler)
	r.POST("/register", middleware.TimeoutMiddleware(), middleware.ErrorHandlerMiddleware(), handlers.RegisterHandler)
	r.POST("/refresh-token", middleware.TimeoutMiddleware(), middleware.ErrorHandlerMiddleware(), handlers.RefreshTokenHandler)
	r.GET("/users", middleware.TimeoutMiddleware(), middleware.AuthMiddleware(), middleware.ErrorHandlerMiddleware(), handlers.ListUsersHandler)
	r.POST("/upload", middleware.TimeoutMiddleware(), middleware.AuthMiddleware(), middleware.ErrorHandlerMiddleware(), handlers.UploadImageHandler)
	r.GET("/open-image/:filename", middleware.TimeoutMiddleware(), middleware.AuthMiddleware(), middleware.ErrorHandlerMiddleware(),handlers.GetImageHandler)

	
	r.Run(":8080")
}
