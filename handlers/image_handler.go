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
            
     w.Header().Set("Content-Type", "application/json")
    
     //time.Sleep(10*time.Second)
    
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
    
    file, header, err := r.FormFile("images")
    if err != nil {
        middleware.ErrorResponse(w, middleware.UNABLE_TO_READ, "No image provided")
        return
    }
    defer file.Close()
    
    var fileExt string
    
    switch header.Header.Get("Content-Type") {
    
    case "image/jpeg":
    
        fileExt = ".jpg"
        
    case "image/png":

        fileExt = ".png"

    case "image/gif":

        fileExt = ".gif"

    case "image/bmp":

        fileExt = ".bmp"

    default:

        middleware.ErrorResponse(w, middleware.UNABLE_TO_READ, "Unsupported image type")

        return
    }
    
    err = os.MkdirAll(imageUploadPath, os.ModePerm)
    if err != nil {
        middleware.ErrorResponse(w, middleware.UNABLE_TO_SAVE, err.Error())
        return
    }
    
    filename := filepath.Join(imageUploadPath, utils.GenerateRequestID()+fileExt)
    
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