package database

import (
	"G-PROJECT/models"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func NewDatabase() (*Database, error) {
	err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        os.Getenv("DB_USERNAME"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_NAME"),
    )

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    err = db.AutoMigrate(&models.User{}, &models.Image{})
    if err != nil {
        return nil, err
    }

    return &Database{DB: db}, nil
}


func (db *Database) Close() error {
	return nil
}

func (db *Database) AddUser(user *models.User) error {
	return db.DB.Create(user).Error
}

func (db *Database) GetUserByUsername(username string) (*models.User, error) {
    var user models.User
    if err := db.DB.Where("username = ?", username).First(&user).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, fmt.Errorf("user not found")
        }
        return nil, err
    }
    return &user, nil
}         


func (db *Database) GetUserByID(userID int) *models.User {
	var user models.User
	db.DB.First(&user, userID)
	return &user
}

func (db *Database) GetAllUsersWithPagination(pageNum, limit int) ([]models.User, error) {
    var users []models.User
    if err := db.DB.Offset((pageNum - 1) * limit).Limit(limit).Find(&users).Error; err != nil {
        return nil, err
    }
    return users, nil
}


func (db *Database) AddImage(image models.Image) error {
	return db.DB.Create(&image).Error
}

func (db *Database) GetAllImages() []models.Image {
	var images []models.Image
	db.DB.Find(&images)
	return images
}

func (db *Database) GetTotalUsersCount() (int, error) {
    var count int64
    if err := db.DB.Model(&models.User{}).Count(&count).Error; err != nil {
        return 0, err
    }
    return int(count), nil
}
