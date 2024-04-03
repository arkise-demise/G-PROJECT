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
	"time"
)

func init() {
    dbInstance = db.NewDatabase()
}

const imageUploadPath = "./images/"

func UploadImageHandler(w http.ResponseWriter, r *http.Request) {
    middleware.AuthMiddleware(http.HandlerFunc(uploadImageHandler)).ServeHTTP(w, r)
}

func uploadImageHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    tokenCookie, err := r.Cookie("token")
    if err != nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    _, err = utils.VerifyToken(tokenCookie.Value)
    if err != nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    err = r.ParseMultipartForm(10 << 20) 
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    file, _, err := r.FormFile("image")
    if err != nil {
        http.Error(w, "No image provided", http.StatusBadRequest)
        return
    }
    defer file.Close()

    err = os.MkdirAll(imageUploadPath, os.ModePerm)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    filename := filepath.Join(imageUploadPath, utils.GenerateUUID()+".jpg")

    newFile, err := os.Create(filename)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer newFile.Close()

    _, err = io.Copy(newFile, file)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{"url": filename})
}

func GetImageHandler(w http.ResponseWriter, r *http.Request) {
    middleware.AuthMiddleware(http.HandlerFunc(getImageHandler)).ServeHTTP(w, r)
}

func getImageHandler(w http.ResponseWriter, r *http.Request) {
    filename := r.URL.Query().Get("images")

    file, err := os.Open(filename)
    if err != nil {
        http.Error(w, "Image not found", http.StatusNotFound)
        return
    }
    defer file.Close()

    http.ServeContent(w, r, filename, time.Time{}, file)
}
