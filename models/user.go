package models

import (
	"encoding/json"
	"fmt"
	"regexp"
)
type PhoneNumber string

type User struct {
    ID       int    `json:"id,omitempty"`
    Username string `json:"username,omitempty"`
    Password string `json:"password,omitempty"`
    Email       string      `json:"email,omitempty"`
    PhoneNumber PhoneNumber `json:"phone_number,omitempty"`
    Address     string      `json:"address,omitempty"`
    Profile     string      `json:"profile_image,omitempty"`
}


func (p PhoneNumber) MarshalJSON() ([]byte, error) {

    if len(p) > 10 {
        formattedNumber := fmt.Sprintf("2519%s", p[5:])
        return json.Marshal(formattedNumber)

    } else {
        formattedNumber := fmt.Sprintf("2519%s", p[2:])
        return json.Marshal(formattedNumber)

    }

}


func (p PhoneNumber) IsValid() bool {
    return regexp.MustCompile(`^(\+2519|09|9|2519)\d{8}$`).MatchString(string(p)) 
}

func (u *User) IsValid() bool {
    return u.isValidEmail() && u.PhoneNumber.IsValid()
}

func (u *User) isValidEmail() bool {
    return regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`).MatchString(u.Email)
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
	User    string `json:"user"`
}

type RefreshTokenRequest struct {
	Username string `json:"username"`
	Password    string `json:"password"`
}
