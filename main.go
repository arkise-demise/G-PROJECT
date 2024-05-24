package main

import (
	"G-PROJECT/database"
	_ "G-PROJECT/docs" // Import the docs package
	"G-PROJECT/handlers"
	"G-PROJECT/middleware"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title Project-1 API
// @version 1.0
// @description This is a sample server for Project-1.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email arkisewdemisew84@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
    dbInstance, err := database.NewDatabase()
    if err != nil {
        log.Fatalf("Failed to initialize database: %v", err)
    }
    defer dbInstance.Close()

    r := gin.Default()
    
    // Swagger documentation route
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    r.POST("/login", middleware.TimeoutMiddleware(), middleware.ErrorHandlerMiddleware(), handlers.LoginHandler)
    r.POST("/register", middleware.TimeoutMiddleware(), middleware.ErrorHandlerMiddleware(), handlers.RegisterHandler)
    r.POST("/refresh-token", middleware.TimeoutMiddleware(), middleware.ErrorHandlerMiddleware(), handlers.RefreshTokenHandler)
    r.GET("/users", middleware.TimeoutMiddleware(), middleware.AuthMiddleware(), middleware.ErrorHandlerMiddleware(), handlers.ListUsersHandler)
    r.POST("/upload", middleware.TimeoutMiddleware(), middleware.AuthMiddleware(), middleware.ErrorHandlerMiddleware(), handlers.UploadImageHandler)
    r.GET("/open-image/:filename", middleware.TimeoutMiddleware(), middleware.AuthMiddleware(), middleware.ErrorHandlerMiddleware(), handlers.GetImageHandler)

    r.Run(":8080")
}

// http://localhost:8080/swagger/index.html