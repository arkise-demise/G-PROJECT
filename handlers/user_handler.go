package handlers

import (
	"G-PROJECT/db"
	"G-PROJECT/middleware"
	"G-PROJECT/models"
	"encoding/json"
	"net/http"
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

    authHandler := middleware.AuthMiddleware(http.HandlerFunc(listUsersHandler))
    authHandler.ServeHTTP(w, r)
}


func listUsersHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    page := 1
    limit := 10 

    pageStr := r.URL.Query().Get("page")
    limitStr := r.URL.Query().Get("limit")

    if pageStr != "" {
        pageValue, err := strconv.Atoi(pageStr)
        if err != nil || pageValue <= 0 {
            middleware.ErrorResponse(w, middleware.UNABLE_TO_READ, "Invalid page number")
            return
        }
        page = pageValue
    }

    if limitStr != "" {
        limitValue, err := strconv.Atoi(limitStr)
        if err != nil || limitValue <= 0 {
            middleware.ErrorResponse(w, middleware.UNABLE_TO_READ, "Invalid limit")
            return
        }
        limit = limitValue
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
