package utils

import (
	"G-PROJECT/models"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("secret_key")

func GenerateToken(user models.User) (string, error) {
    claims := jwt.MapClaims{
        "id": user.ID,
        "exp": time.Now().Add(time.Hour * 24).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        return "", err
    }
    return tokenString, nil
}

func VerifyToken(tokenString string) (models.User, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return jwtKey, nil
    })
    if err != nil {
        return models.User{}, err
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        userID := int(claims["id"].(float64))
       
        user := models.User{
            ID: userID,
        }
        return user, nil
    } else {
        return models.User{}, fmt.Errorf("invalid token")
    }
}
