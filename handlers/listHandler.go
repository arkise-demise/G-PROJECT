package handlers

import (
	"G-PROJECT/database"
	"G-PROJECT/middleware"
	"G-PROJECT/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MetaData struct {
    Page    int `json:"page"`
    PerPage int `json:"per_page"`
    Total   int `json:"total"`
}


func init() {
    var err error
    dbInstance, err = database.NewDatabase()
    if err != nil {
        panic(err)
    }
}

func ListUsersHandler(c *gin.Context) {
    page := 1
    limit := 20

    pageStr := c.Query("page")
    limitStr := c.Query("limit")

    if pageStr != "" {
        pageValue, err := strconv.Atoi(pageStr)
        if err != nil || pageValue <= 0 {
            c.Set("error", middleware.CustomError{
                Type:    middleware.UNABLE_TO_READ,
                Message: "Invalid page number",
            })
            return
        }
        page = pageValue
    }

    if limitStr != "" {
        limitValue, err := strconv.Atoi(limitStr)
        if err != nil || limitValue <= 0 {
            c.Set("error", middleware.CustomError{
                Type:    middleware.UNABLE_TO_READ,
                Message: "Invalid limit",
            })
            return
        }
        limit = limitValue
    }

    users, err := dbInstance.GetAllUsersWithPagination(page, limit)
if err != nil {
    c.Set("error", middleware.CustomError{
        Type:    middleware.UNABLE_TO_READ,
        Message: "Failed to retrieve users",
    })
    return
}

    totalUsers, err := dbInstance.GetTotalUsersCount()
    if err != nil {
        c.Set("error", middleware.CustomError{
            Type:    middleware.UNABLE_TO_READ,
            Message: "Failed to retrieve total users count",
        })
        return
    }

    metaData := MetaData{
        Page:    page,
        PerPage: limit,
        Total:   totalUsers,
    }

    successResponse := struct {
        MetaData MetaData      `json:"meta_data"`
        Data     []models.User `json:"data"`
    }{
        MetaData: metaData,
        Data:     users,
    }

    c.JSON(http.StatusOK, successResponse)
}
