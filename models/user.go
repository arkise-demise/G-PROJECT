package models

import "errors"

type User struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
    Password string `json:"password"`
}

func (user *User) Validate() error {
    if user.Username == "" {
        return errors.New("username cannot be empty")
    }
    if user.Password == "" {
        return errors.New("password cannot be empty")
    }
    return nil
}
