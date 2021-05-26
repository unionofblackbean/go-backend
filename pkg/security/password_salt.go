package security

import (
	"crypto/rand"
	"fmt"
)

func GeneratePasswordSalt() ([]byte, error) {
	// generate salt
	salt := make([]byte, 64)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, fmt.Errorf("failed to read random bytes from random number generator -> %v", err)
	}

	return salt, nil
}
