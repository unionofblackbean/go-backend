package security

import (
	"crypto/rand"
	"fmt"
)

func GeneratePasswordSalt() (salt []byte, err error) {
	// generate salt
	salt = make([]byte, 64)
	_, err = rand.Read(salt)
	if err != nil {
		err = fmt.Errorf("failed to read random bytes from random number generator -> %v", err)
	}

	return
}
