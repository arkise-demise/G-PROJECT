package models

import (
	"encoding/json"
	"fmt"
	"regexp"
)

type PhoneNumber string

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
    return regexp.MustCompile(`^(\+2519|09)\d{8}$`).MatchString(string(p)) 
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
