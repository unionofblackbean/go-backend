package models

import (
	"github.com/google/uuid"
	"github.com/unionofblackbean/backend/pkg/encoding"
	"github.com/unionofblackbean/backend/pkg/security"
)

type User struct {
	UUID         uuid.UUID
	PasswordHash string
	PasswordSalt string
}

func NewUser(uuid uuid.UUID, password string) (*User, error) {
	passwordSalt, err := security.GeneratePasswordSalt()
	if err != nil {
		return nil, err
	}

	passwordHash, err := security.HashPassword(password, passwordSalt)
	if err != nil {
		return nil, err
	}

	return &User{
		UUID:         uuid,
		PasswordHash: encoding.Base64RawStdEncodeToString(passwordHash),
		PasswordSalt: encoding.Base64RawStdEncodeToString(passwordSalt),
	}, nil
}
