package handlers

import (
	"G-PROJECT/db"
	"G-PROJECT/middleware"
	"G-PROJECT/models"
	"G-PROJECT/utils"
	"encoding/json"
	"net/http"
	"time"
)

func init() {
	dbInstance = db.NewDatabase()
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	storedUser := dbInstance.GetUserByUsername(user.Username)
	if storedUser == nil || storedUser.Password != user.Password {
		middleware.ErrorResponse(w, middleware.UNAUTHORIZED, "Invalid username or password")
		return
	}

	tokenString, err := utils.GenerateToken(*storedUser)
	if err != nil {
		middleware.ErrorResponse(w, middleware.UNABLE_TO_SAVE, err.Error())
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		middleware.ErrorResponse(w, middleware.UNABLE_TO_READ, err.Error())
		return
	}

	if dbInstance.GetUserByUsername(user.Username) != nil {
		middleware.ErrorResponse(w, middleware.UNABLE_TO_SAVE, "Username already exists")
		return
	}

	dbInstance.AddUser(user)

	json.NewEncoder(w).Encode(user)
}

func RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		middleware.ErrorResponse(w, middleware.UNABLE_TO_READ, err.Error())
		return
	}

	storedUser := dbInstance.GetUserByUsername(user.Username)
	if storedUser == nil || storedUser.Password != user.Password {
		middleware.ErrorResponse(w, middleware.UNAUTHORIZED, "Invalid username or password")
		return
	}

	tokenString, err := utils.GenerateToken(*storedUser)
	if err != nil {
		middleware.ErrorResponse(w, middleware.UNABLE_TO_SAVE, err.Error())
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})
}
