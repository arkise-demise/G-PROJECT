package handlers

import (
	"G-PROJECT/db"
	"G-PROJECT/models"
	"G-PROJECT/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"time"
)

func init() {
	dbInstance = db.NewDatabase()
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// time.Sleep(10*time.Second)

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		ErrorResponse(w , UNABLE_TO_READ,err.Error())
		return
	}

	storedUser := dbInstance.GetUserByUsername(user.Username)
	if storedUser == nil || storedUser.Password != user.Password {
		ErrorResponse(w, UNAUTHORIZED, "unauthorized user!")
		return
	}

	tokenString, err := utils.GenerateToken(*storedUser)
	if err != nil {
		ErrorResponse(w, UNABLE_TO_SAVE, err.Error())
		return
	}

	fmt.Printf("User '%s' successfully logged in!\n", storedUser.Username)


	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  time.Now().Add(3 * time.Minute),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})
}



func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	
	// time.Sleep(10*time.Second)

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		ErrorResponse(w, UNABLE_TO_READ, err.Error())
		return
	}

	if dbInstance.GetUserByUsername(user.Username) != nil {
		ErrorResponse(w, UNABLE_TO_SAVE, "Username already exists")
		return
	}

	if !isValidUser(user) {
		ErrorResponse(w, UNABLE_TO_SAVE, "Invalid user data")
		return
	}

	dbInstance.AddUser(user)

	json.NewEncoder(w).Encode(user)
}

func isValidUser(user models.User) bool {
	if !isValidEmail(user.Email) {
		return false
	}
	if !isValidPhoneNumber(user.PhoneNumber) {
		return false
	}
	if !isValidAddress(user.Address) {
		return false
	}
	return true
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
	
	
func RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	   // time.Sleep(10*time.Second)

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		ErrorResponse(w, UNABLE_TO_READ, err.Error())
		return
	}

	storedUser := dbInstance.GetUserByUsername(user.Username)
	
	tokenString, err := utils.GenerateToken(*storedUser)
	if err != nil {
		ErrorResponse(w, UNABLE_TO_SAVE, err.Error())
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  time.Now().Add(3 * time.Minute),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})
}