package utils

import "golang.org/x/crypto/bcrypt"

func ComparePasswords(hashedPassword string, plaintextPassword string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plaintextPassword))
    return err == nil
}
