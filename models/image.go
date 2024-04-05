package models

type Image struct {
    ID     int    `json:"id"`
    Name   string `json:"name"`
    Path   string `json:"path"`
    UserID int    `json:"userId"` 
}


