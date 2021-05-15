package security

import "crypto/rand"

func GeneratePasswordSalt() (salt []byte, err error) {
	// generate salt
	salt = make([]byte, 64)
	_, err = rand.Read(salt)
	return
}
