package utils

import (
	"crypto/rand"
	"encoding/base64"
	"log"
)

// GenerateRandomToken generates a random token.
//
// Returns a string and an error.
func  GenerateRandomToken(lenght int) (string, error){
	bytes := make([]byte, lenght)

	if _, err := rand.Read(bytes); err != nil {
		log.Printf("@GenerateRandomToken Error generating random token: %v", err)
		return "", err
	}

	token := base64.URLEncoding.EncodeToString(bytes)
	return token, nil
}