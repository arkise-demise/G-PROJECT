package models

import "errors"

type Image struct {
    ID     int    `json:"id"`
    Name   string `json:"name"`
    Path   string `json:"path"`
    UserID int    `json:"userId"` 
}

func (image *Image) Validate() error {
    if image.Name == "" {
        return errors.New("image name can't be empty")
    }
    if image.UserID <= 0 {
        return errors.New("invalid user ID")
    }
    return nil
}
