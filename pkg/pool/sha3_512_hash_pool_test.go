package pool

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/sha3"
)

func TestGetSha3512Hash(t *testing.T) {
	shaHash := GetSha3512Hash()
	assert.NotNil(t, shaHash)
}

func TestPutSha3512Hash(t *testing.T) {
	shaHash := sha3.New512()
	PutSha3512Hash(shaHash)

	hashFromPool := GetSha3512Hash()
	assert.Same(t, shaHash, hashFromPool)
}
