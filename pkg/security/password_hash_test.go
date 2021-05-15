package security

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	salt, err := GeneratePasswordSalt()
	assert.Nil(t, err)

	_, err = HashPassword("test", salt)
	assert.Nil(t, err)
}
