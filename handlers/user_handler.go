package handlers

import (
	"G-PROJECT/db"
	"G-PROJECT/utils"
	"encoding/json"
	"net/http"
)

var dbInstance *db.Database

func init() {
    dbInstance = db.NewDatabase()
}

func ListUsersHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    if ctx.Err() != nil {
		http.Error(w, "Request timed out", http.StatusRequestTimeout)
		return
	}
    tokenCookie, err := r.Cookie("token")
    if err != nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    _, err = utils.VerifyToken(tokenCookie.Value)
    if err != nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    users := dbInstance.GetAllUsers()
    json.NewEncoder(w).Encode(users)
}
