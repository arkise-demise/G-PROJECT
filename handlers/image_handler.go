package handlers

import (
	"G-PROJECT/db"
	"G-PROJECT/middleware"
	"G-PROJECT/utils"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
)


func init() {
    dbInstance = db.NewDatabase()
}

const (
    imageUploadPath = "./images/"
    maxUploadSize   = 32 << 20
)

func UploadImageHandler(w http.ResponseWriter, r *http.Request) {
    middleware.AuthMiddleware(http.HandlerFunc(uploadImageHandler)).ServeHTTP(w, r)
}

func uploadImageHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    if r.Context().Err() != nil {
        middleware.ErrorResponse(w, middleware.UNABLE_TO_READ, "Request timed out")
        return
    }

    tokenCookie, err := r.Cookie("token")
    if err != nil {
        middleware.ErrorResponse(w, middleware.UNAUTHORIZED, "Unauthorized")
        return
    }

    _, err = utils.VerifyToken(tokenCookie.Value)
    if err != nil {
        middleware.ErrorResponse(w, middleware.UNAUTHORIZED, "Unauthorized")
        return
    }

    r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)

    err = r.ParseMultipartForm(maxUploadSize)
    if err != nil {
        middleware.ErrorResponse(w, middleware.UNABLE_TO_READ, err.Error())
        return
    }

    file, _, err := r.FormFile("images")
    if err != nil {
        middleware.ErrorResponse(w, middleware.UNABLE_TO_READ, "No image provided")
        return
    }
    defer file.Close()

    err = os.MkdirAll(imageUploadPath, os.ModePerm)
    if err != nil {
        middleware.ErrorResponse(w, middleware.UNABLE_TO_SAVE, err.Error())
        return
    }

    filename := filepath.Join(imageUploadPath, utils.GenerateRequestID()+".jpg")

    newFile, err := os.Create(filename)
    if err != nil {
        middleware.ErrorResponse(w, middleware.UNABLE_TO_SAVE, err.Error())
        return
    }
    defer newFile.Close()

    _, err = io.Copy(newFile, file)
    if err != nil {
        middleware.ErrorResponse(w, middleware.UNABLE_TO_SAVE, err.Error())
        return
    }

    json.NewEncoder(w).Encode(map[string]string{"url": filename})
}

func GetImageHandler(w http.ResponseWriter, r *http.Request) {
    middleware.AuthMiddleware(http.HandlerFunc(getImageHandler)).ServeHTTP(w, r)
}

func getImageHandler(w http.ResponseWriter, r *http.Request) {
    filename := r.URL.Path[len("/open-image/"):]

    imagePath := filepath.Join(imageUploadPath, filename)

    file, err := os.Open(imagePath)
    if err != nil {
        middleware.ErrorResponse(w, middleware.UNABLE_TO_FIND_RESOURCE, "Image not found")
        return
    }
    defer file.Close()

    http.ServeFile(w, r, imagePath)
}
