package db

import "G-PROJECT/models"

type Database struct {
	Users  map[string]models.User
	Images map[int]models.Image
}

func NewDatabase() *Database {
	return &Database{
		Users:  make(map[string]models.User),
		Images: make(map[int]models.Image),
	}
}

func (db *Database) AddUser(user models.User) {
	db.Users[user.Username] = user
}

func (db *Database) GetUserByUsername(username string) *models.User {
	user, ok := db.Users[username]
	if !ok {
		return nil
	}
	return &user
}

func (db *Database) GetUserByID(userID int) *models.User {
	for _, user := range db.Users {
		if user.ID == userID {
			return &user
		}
	}
	return nil
}

func (db *Database) GetAllUsers() []models.User {
	users := make([]models.User, 0, len(db.Users))
	for _, user := range db.Users {
		users = append(users, user)
	}
	return users
}

func (db *Database) AddImage(image models.Image) {
	db.Images[image.ID] = image
}

func (db *Database) GetAllImages() []models.Image {
	images := make([]models.Image, 0, len(db.Images))
	for _, image := range db.Images {
		images = append(images, image)
	}
	return images
}
