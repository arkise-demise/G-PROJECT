package handlers

import (
	"G-PROJECT/db"
	"G-PROJECT/middleware"
	"G-PROJECT/models"
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"
)

var dbInstance *db.Database

type MetaData struct {
    Page    int `json:"page"`
    PerPage int `json:"per_page"`
    Total   int `json:"total"`
}

func init() {
    dbInstance = db.NewDatabase()
}

func ListUsersHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    pageStr := r.URL.Query().Get("page")
    limitStr := r.URL.Query().Get("limit")

    page, err := strconv.Atoi(pageStr)
    if err != nil || page <= 0 {
        middleware.ErrorResponse(w, middleware.UNABLE_TO_READ, "Invalid page number")
        return
    }

    limit, err := strconv.Atoi(limitStr)
    if err != nil || limit <= 0 {
        middleware.ErrorResponse(w, middleware.UNABLE_TO_READ, "Invalid limit")
        return
    }

    users, err := dbInstance.GetAllUsersWithPagination(page, limit)
    if err != nil {
        middleware.ErrorResponse(w, middleware.UNABLE_TO_READ, "Failed to retrieve users")
        return
    }

    totalUsers := len(dbInstance.Users)

    metaData := MetaData{
        Page:    page,
        PerPage: limit,
        Total:   totalUsers,
    }

    successResponse := struct {
        MetaData MetaData        `json:"meta_data"`
        Data     []models.User   `json:"data"`
    }{
        MetaData: metaData,
        Data:     users,
    }

    json.NewEncoder(w).Encode(successResponse)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    var user models.User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        middleware.ErrorResponse(w, middleware.UNABLE_TO_READ, err.Error())
        return
    }

    if !isValidUser(user) {
        middleware.ErrorResponse(w, middleware.UNABLE_TO_READ, "Invalid user data")
        return
    }

    dbInstance.AddUser(user)

    user = removeEmptyFields(user).(models.User)

    json.NewEncoder(w).Encode(user)
}

func isValidUser(user models.User) bool {
    if !isValidEmail(user.Email) {
        return false
    }
    if !isValidPhoneNumber(user.PhoneNumber) {
        return false
    }
    if !isValidAddress(user.Address) {
        return false
    }
    return true
}

func isValidEmail(email string) bool {
    return regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString(email)
}

func isValidPhoneNumber(phoneNumber models.PhoneNumber) bool {
    return phoneNumber.IsValid()
}

func isValidAddress(address string) bool {
    return !regexp.MustCompile(`\d`).MatchString(address)
}

func removeEmptyFields(data interface{}) interface{} {
    switch value := data.(type) {
    case models.User:
        if value.ID == 0 {
            value.ID = -1
        }
        if value.Username == "" {
            value.Username = "N/A"
        }
        if value.Email == "" {
            value.Email = "N/A"
        }
        if value.PhoneNumber == "" {
            value.PhoneNumber = models.PhoneNumber("N/A")
        }
        if value.Address == "" {
            value.Address = "N/A"
        }
        return value
    default:
        return data
    }
}
