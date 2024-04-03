package handlers

import (
	"G-PROJECT/utils"
	"encoding/json"
	"net/http"
	"time"
)

func RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
    expiredToken := r.Header.Get("Authorization")

    user, err := utils.VerifyToken(expiredToken)
    if err != nil {
        http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
        return
    }

    newToken, err := utils.GenerateToken(user)
    if err != nil {
        http.Error(w, "Failed to generate new token", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Authorization", newToken)

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "message": "Token refreshed successfully",
        "expiration": time.Now().Add(utils.TokenExpiration).Unix(),
    })
}
