package dao

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/unionofblackbean/go-backend/services/auth/api/entities"
)

func CreateUser(user *entities.User) (err error) {
	err = pool.Validate()
	if err != nil {
		return
	}

	err = pool.Exec(
		"INSERT INTO users (uuid, password_hash, password_salt) VALUES ($1, $2, $3);",
		pgtype.UUID{
			Bytes:  user.UUID,
			Status: pgtype.Present,
		},
		user.PasswordHash,
		user.PasswordSalt,
	)
	return
}

func GetUser(uuid uuid.UUID) (*entities.User, error) {
	err := pool.Validate()
	if err != nil {
		return nil, err
	}

	var rawUUID pgtype.UUID
	var rawPasswordHash string
	var rawPasswordSalt string
	err = pool.QueryRow(
		"SELECT uuid, password_hash, password_salt FROM users WHERE uuid=$1;",
		pgtype.UUID{
			Bytes:  uuid,
			Status: pgtype.Present,
		},
	).Scan(&rawUUID, &rawPasswordHash, &rawPasswordSalt)
	if err != nil {
		return nil, err
	}

	user := entities.User{
		UUID:         rawUUID.Bytes,
		PasswordHash: rawPasswordHash,
		PasswordSalt: rawPasswordSalt,
	}
	return &user, nil
}

func GetAllUsersUUID() ([]string, error) {
	err := pool.Validate()
	if err != nil {
		return nil, err
	}

	rows, err := pool.Query(
		"SELECT uuid FROM users;",
	)
	if err != nil {
		return nil, err
	}

	var uuids []string
	for rows.Next() {
		var rawUserUUID pgtype.UUID
		err = rows.Scan(&rawUserUUID)
		if err != nil {
			return nil, err
		}

		userUUID, err := uuid.FromBytes(rawUserUUID.Bytes[:])
		if err != nil {
			return nil, fmt.Errorf("failed to process UUID obtained from database -> %v", err)
		}

		uuids = append(uuids, userUUID.String())
	}

	return uuids, nil
}

func UpdateUser(user *entities.User) (err error) {
	err = pool.Validate()
	if err != nil {
		return
	}

	err = pool.Exec(
		"UPDATE users SET password_hash=$1, password_salt=$2 WHERE uuid=$3;",
		user.PasswordHash,
		user.PasswordSalt,
		pgtype.UUID{
			Bytes:  user.UUID,
			Status: pgtype.Present,
		},
	)
	return
}

func DeleteUser(uuid uuid.UUID) (err error) {
	err = pool.Validate()
	if err != nil {
		return
	}

	err = pool.Exec(
		"DELETE FROM users WHERE uuid=$1;",
		pgtype.UUID{
			Bytes:  uuid,
			Status: pgtype.Present,
		},
	)
	return
}

func IsExistsUser(uuid uuid.UUID) (exists bool, err error) {
	err = pool.Validate()
	if err != nil {
		return
	}

	err = pool.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM users WHERE uuid=$1);",
		pgtype.UUID{
			Bytes:  uuid,
			Status: pgtype.Present,
		},
	).Scan(&exists)

	return
}
