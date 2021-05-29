package pool

import (
	"hash"
	"sync"

	"golang.org/x/crypto/sha3"
)

var sha3512HashPool sync.Pool

func GetSha3512Hash() hash.Hash {
	if rawSha3512Hash := sha3512HashPool.Get(); rawSha3512Hash != nil {
		sha3512Hash, _ := rawSha3512Hash.(hash.Hash)
		sha3512Hash.Reset()
		return sha3512Hash
	}

	return sha3.New512()
}

func PutSha3512Hash(hash hash.Hash) {
	sha3512HashPool.Put(hash)
}
