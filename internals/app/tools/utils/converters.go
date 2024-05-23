package utils

import (
	"log"
	"strconv"
)

// ConvertStringToUint converts a string to an unsigned integer.
//
// It takes a string as a parameter and returns an unsigned integer and an error.
func StringToUint(str string) (uint, error) {
	idUint, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		log.Printf("@ConvertStringToUint: Error converting string to uint: %v", err)
		return 0, err
	}
	return uint(idUint), nil
}