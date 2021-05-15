package security

import "github.com/unionofblackbean/backend/pkg/pool"

func HashPassword(password string, salt []byte) ([]byte, error) {
	buf := pool.GetBuffer()
	defer pool.PutBuffer(buf)

	// write password to buffer
	_, err := buf.WriteString(password)
	if err != nil {
		return nil, err
	}

	// write salt to buffer
	_, err = buf.Write(salt)
	if err != nil {
		return nil, err
	}

	// hash password&salt
	hash := pool.GetSha3512Hash()
	defer pool.PutSha3512Hash(hash)
	_, err = hash.Write(buf.Bytes())
	if err != nil {
		return nil, err
	}

	return hash.Sum(nil), nil
}
