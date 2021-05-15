package security

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePasswordSalt(t *testing.T) {
	salt, err := GeneratePasswordSalt()
	assert.Nil(t, err)
	assert.NotNil(t, salt)
}
