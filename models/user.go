package models

import (
	"encoding/json"
	"fmt"
	"regexp"
)

type PhoneNumber string

func (p PhoneNumber) MarshalJSON() ([]byte, error) {
    formattedNumber := fmt.Sprintf("%s-------", p[:4]) 
    return json.Marshal(formattedNumber)
}

func (p PhoneNumber) IsValid() bool {
    return regexp.MustCompile(`^\d{10}$`).MatchString(string(p)) 
}

type User struct {
    ID          int         `json:"id"`
    Username    string      `json:"username"`
    Password    string      `json:"password"`
    Email       string      `json:"email"`
    PhoneNumber PhoneNumber `json:"phone_number"`
    Address     string      `json:"address"`
}

func (u *User) IsValid() bool {
    return u.isValidEmail() && u.PhoneNumber.IsValid()
}

func (u *User) isValidEmail() bool {
    return regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`).MatchString(u.Email)
}
