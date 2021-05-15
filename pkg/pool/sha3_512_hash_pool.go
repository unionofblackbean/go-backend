package pool

import (
	"hash"
	"sync"

	"golang.org/x/crypto/sha3"
)

var sha3512HashPool sync.Pool

func GetSha3512Hash() hash.Hash {
	if rawHash := sha3512HashPool.Get(); rawHash != nil {
		hash, _ := rawHash.(hash.Hash)
		hash.Reset()
		return hash
	}

	return sha3.New512()
}

func PutSha3512Hash(hash hash.Hash) {
	sha3512HashPool.Put(hash)
}
