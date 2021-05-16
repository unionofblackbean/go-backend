package dao

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/unionofblackbean/backend/services/auth/rest/models"
)

func CreateUser(user *models.User) (err error) {
	err = checkPool()
	if err != nil {
		return
	}

	err = pool.Exec(
		"INSERT INTO users VALUES ($1, $2, $3);",
		pgtype.UUID{
			Bytes:  user.UUID,
			Status: pgtype.Present,
		},
		user.PasswordHash,
		user.PasswordSalt,
	)
	return
}

func GetUser(uuid uuid.UUID) (user *models.User, err error) {
	err = checkPool()
	if err != nil {
		return
	}

	var rawUUID pgtype.UUID
	var rawPasswordHash string
	var rawPasswordSalt string
	err = pool.QueryRow(
		"SELECT * FROM users WHERE uuid=$1;",
		pgtype.UUID{
			Bytes:  uuid,
			Status: pgtype.Present,
		},
	).Scan(&rawUUID, &rawPasswordHash, &rawPasswordSalt)
	if err != nil {
		return
	}

	user = &models.User{
		UUID:         rawUUID.Bytes,
		PasswordHash: rawPasswordHash,
		PasswordSalt: rawPasswordSalt,
	}
	return
}

func GetAllUsers() (users []models.User, err error) {
	err = checkPool()
	if err != nil {
		return
	}

	rows, err := pool.Query(
		"SELECT * FROM users;",
	)
	if err != nil {
		return
	}

	for rows.Next() {
		var uuid pgtype.UUID
		var passwordHash string
		var passwordSalt string
		err = rows.Scan(
			&uuid,
			&passwordHash,
			&passwordSalt,
		)
		if err != nil {
			return
		}

		users = append(users, models.User{
			UUID:         uuid.Bytes,
			PasswordHash: passwordHash,
			PasswordSalt: passwordSalt,
		})
	}

	return
}

func GetAllUsersUUID() (uuids []string, err error) {
	err = checkPool()
	if err != nil {
		return
	}

	rows, err := pool.Query(
		"SELECT uuid FROM users;",
	)
	if err != nil {
		return
	}

	for rows.Next() {
		var rawUserUUID pgtype.UUID
		err = rows.Scan(&rawUserUUID)
		if err != nil {
			return
		}

		userUUID, err := uuid.FromBytes(rawUserUUID.Bytes[:])
		if err != nil {
			return nil, fmt.Errorf("failed to process UUID obtained from database -> %v", err)
		}

		uuids = append(uuids, userUUID.String())
	}

	return
}

func UpdateUser(user *models.User) (err error) {
	err = checkPool()
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
	err = checkPool()
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
