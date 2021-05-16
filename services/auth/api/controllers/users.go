package controllers

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/unionofblackbean/backend/pkg/encoding"
	"github.com/unionofblackbean/backend/pkg/responses"
	"github.com/unionofblackbean/backend/pkg/security"
	"github.com/unionofblackbean/backend/services/auth/api/dao"
	"github.com/unionofblackbean/backend/services/auth/api/models"
)

func Users(ctx *fiber.Ctx) error {
	switch ctx.Method() {
	case http.MethodGet:
		users, err := dao.GetAllUsersUUID()
		if err != nil {
			return fmt.Errorf("failed to get all users' UUID -> %v", err)
		}

		responses.SendDataResponse(ctx, &users)
	case http.MethodPost:
		rawUUID := ctx.FormValue("uuid")
		rawPassword := ctx.FormValue("password")

		if rawUUID == "" || rawPassword == "" {
			return errors.New("missing form value uuid or password")
		}

		uuid, err := uuid.Parse(rawUUID)
		if err != nil {
			return fmt.Errorf("failed to parse UUID -> %v", err)
		}

		exists, err := dao.IsExistsUser(uuid)
		if err != nil {
			return fmt.Errorf("failed to check user's existence -> %v", err)
		}
		if exists {
			responses.SendResourceAlreadyExistsResponse(ctx)
			return nil
		}

		user, err := models.NewUser(uuid, rawPassword)
		if err != nil {
			return fmt.Errorf("failed to create new user object -> %v", err)
		}

		err = dao.CreateUser(user)
		if err != nil {
			return fmt.Errorf("failed to create user -> %v", err)
		}
	case http.MethodPatch:
		rawUUID := ctx.FormValue("uuid")
		oldPassword := ctx.FormValue("old_password")
		newPassword := ctx.FormValue("new_password")

		if rawUUID == "" || oldPassword == "" || newPassword == "" {
			return errors.New("missing form value uuid, old_password or new_password")
		}

		uuid, err := uuid.Parse(rawUUID)
		if err != nil {
			return fmt.Errorf("failed to parse UUID -> %v", err)
		}

		exists, err := dao.IsExistsUser(uuid)
		if err != nil {
			return fmt.Errorf("failed to check user's existence -> %v", err)
		}
		if !exists {
			responses.SendResourceNotFoundResponse(ctx)
			return nil
		}

		dbUser, err := dao.GetUser(uuid)
		if err != nil {
			return fmt.Errorf("failed to get user -> %v", err)
		}

		// password salt from database
		dbPasswordSalt, err := encoding.Base64RawStdDecodeString(dbUser.PasswordSalt)
		if err != nil {
			return fmt.Errorf("failed to base64 decode password salt from database -> %v", err)
		}

		// password hash produced from "old_password" form value + password salt from database
		passwordHash, err := security.HashPassword(
			oldPassword,
			dbPasswordSalt,
		)
		if err != nil {
			return fmt.Errorf("failed to hash inputted old password -> %v", err)
		}

		// real old password hash from database
		dbPasswordHash, err := encoding.Base64RawStdDecodeString(dbUser.PasswordHash)
		if err != nil {
			return fmt.Errorf("failed to base64 decode password hash from database -> %v", err)
		}

		// compare produced password hash and real password hash
		if bytes.Equal(passwordHash, dbPasswordHash) {
			newUser, err := models.NewUser(dbUser.UUID, newPassword)
			if err != nil {
				return fmt.Errorf("failed to create new user object -> %v", err)
			}

			err = dao.UpdateUser(newUser)
			if err != nil {
				return fmt.Errorf("failed to update user -> %v", err)
			}
		}
	case http.MethodDelete:
		rawUUID := ctx.FormValue("uuid")
		password := ctx.FormValue("password")

		if rawUUID == "" || password == "" {
			return errors.New("missing form value uuid or password")
		}

		uuid, err := uuid.Parse(rawUUID)
		if err != nil {
			return fmt.Errorf("failed to parse UUID -> %v", err)
		}

		exists, err := dao.IsExistsUser(uuid)
		if err != nil {
			return fmt.Errorf("failed to check user's existence -> %v", err)
		}
		if !exists {
			responses.SendResourceNotFoundResponse(ctx)
			return nil
		}

		dbUser, err := dao.GetUser(uuid)
		if err != nil {
			return fmt.Errorf("failed to get user from database -> %v", err)
		}

		dbPasswordSalt, err := encoding.Base64RawStdDecodeString(dbUser.PasswordSalt)
		if err != nil {
			return fmt.Errorf("failed to base64 decode password salt from database -> %v", err)
		}

		// password hash produced from "old_password" form value + password salt from database
		passwordHash, err := security.HashPassword(password, dbPasswordSalt)
		if err != nil {
			return fmt.Errorf("failed to hash inputted password -> %v", err)
		}

		// password hash from database
		dbPasswordHash, err := encoding.Base64RawStdDecodeString(dbUser.PasswordHash)
		if err != nil {
			return fmt.Errorf("failed to base64 decode password hash from database -> ")
		}

		// compare produced password hash and real password hash
		if bytes.Equal(passwordHash, dbPasswordHash) {
			err := dao.DeleteUser(uuid)
			if err != nil {
				return fmt.Errorf("failed to delete user -> %v", err)
			}
		} else {
			return errors.New("incorrect password")
		}
	default:
		responses.SendUnsupportedMethodResponse(ctx)
	}

	return nil
}