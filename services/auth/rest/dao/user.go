package dao

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/unionofblackbean/backend/services/auth/rest/models"
)

func CreateUser(user *models.User) (err error) {
	_, err = pool.Exec(context.Background(),
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
	var rawUUID pgtype.UUID
	var rawPasswordHash string
	var rawPasswordSalt string
	err = pool.QueryRow(context.Background(),
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

func GetAllUsers() ([]models.User, error) {
	rows, err := pool.Query(context.Background(),
		"SELECT * FROM users;",
	)
	if err != nil {
		return nil, err
	}

	var users []models.User
	for rows.Next() {
		var uuid pgtype.UUID
		var passwordHash string
		var passwordSalt string
		rows.Scan(
			&uuid,
			&passwordHash,
			&passwordSalt,
		)

		users = append(users, models.User{
			UUID:         uuid.Bytes,
			PasswordHash: passwordHash,
			PasswordSalt: passwordSalt,
		})
	}

	return users, nil
}

func UpdateUser(user *models.User) (err error) {
	_, err = pool.Exec(context.Background(),
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
	_, err = pool.Exec(context.Background(),
		"DELETE FROM users WHERE uuid=$1;",
		pgtype.UUID{
			Bytes:  uuid,
			Status: pgtype.Present,
		},
	)
	return
}
