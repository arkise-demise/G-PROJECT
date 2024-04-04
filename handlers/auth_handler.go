package handlers

import (
	"G-PROJECT/models"
	"G-PROJECT/utils"
	"encoding/json"
	"net/http"
	"time"
)

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
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	tokenString, err := utils.GenerateToken(*storedUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if dbInstance.GetUserByUsername(user.Username) != nil {
        http.Error(w, "Username already exists", http.StatusBadRequest)
        return
    }

    dbInstance.AddUser(user)

    json.NewEncoder(w).Encode(user)
}
