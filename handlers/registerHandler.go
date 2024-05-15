package handlers

import (
	"G-PROJECT/database"
	"G-PROJECT/middleware"
	"G-PROJECT/models"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var dbInstance *database.Database

func init() {
	var err error
	dbInstance, err = database.NewDatabase()
	if err != nil {
		panic(err)
	}
}

func RegisterHandler(c *gin.Context) {
	var userInput models.User
	if err := c.BindJSON(&userInput); err != nil {
		c.Set("error", middleware.CustomError{
			Type:    middleware.UNABLE_TO_READ,
			Message: err.Error(),
		})
		return
	}

	// Validate user details

	if !isValidUser(userInput) {
		c.Set("error", middleware.CustomError{
			Type:    middleware.UNABLE_TO_SAVE,
			Message: "Invalid user data",
		})
		return
	}
	formattedPhoneNumber, err := userInput.PhoneNumber.MarshalJSON()
	if err != nil {
		c.Set("error", middleware.CustomError{
			Type:    middleware.UNABLE_TO_SAVE,
			Message: "Failed to format phone number",
		})
		return
	}

	// Update user data with formatted phone number
	userInput.PhoneNumber = models.PhoneNumber(formattedPhoneNumber)

	var user models.User
	dbInstance.DB.Where("username = ?", userInput.Username).Find(&user)
	if user.ID != 0 {
		c.Set("error", middleware.CustomError{
			Type:    middleware.UNABLE_TO_SAVE,
			Message: "Username already exists",
		})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	if err != nil {
		c.Set("error", middleware.CustomError{
			Type:    middleware.UNABLE_TO_FIND_RESOURCE,
			Message: "Can't hash the password",
		})
		return
	}

	// Create user data
	userData := models.User{
		Username:    userInput.Username,
		Password:    string(hashedPassword),
		Email:       userInput.Email,
		PhoneNumber: userInput.PhoneNumber,
		Address:     userInput.Address,
	}

	// Add the user to the database
	if err := dbInstance.DB.Create(&userData).Error; err != nil {
		c.Set("error", middleware.CustomError{
			Type:    middleware.UNABLE_TO_SAVE,
			Message: "Failed to register user",
		})
		return
	}
	c.String(http.StatusOK,"user registered successfully!")
 
	c.IndentedJSON(http.StatusOK, gin.H{"user": userData})
}

func isValidUser(user models.User) bool {
	return isValidEmail(user.Email) && isValidPhoneNumber(user.PhoneNumber) && isValidAddress(user.Address)
}

func isValidEmail(email string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString(email)
}

func isValidPhoneNumber(phoneNumber models.PhoneNumber) bool {
	return phoneNumber.IsValid()
}

func isValidAddress(address string) bool {
	return !regexp.MustCompile(`\d`).MatchString(address)
}
