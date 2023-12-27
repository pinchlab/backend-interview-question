package utils

import (
	"errors"
	"math/rand"
)

// UploadImage uploads image
func UploadImage(file []byte) (string, error) {
	success := rand.Intn(2) == 1
	if !success {
		return "", errors.New("upload image failed")
	}

	return "https://image.com/1.png", nil
}
