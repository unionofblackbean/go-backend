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
	"github.com/unionofblackbean/backend/services/auth/rest/dao"
	"github.com/unionofblackbean/backend/services/auth/rest/models"
)

func Users(ctx *fiber.Ctx) error {
	switch ctx.Method() {
	case http.MethodGet:
		rawUUID := ctx.Params("uuid")

		if rawUUID == "" {
			users, err := dao.GetAllUsers()
			if err != nil {
				return fmt.Errorf("failed to get all users -> %v", err)
			}

			responses.SendDataResponse(ctx, &users)
		} else {
			uuid, err := uuid.Parse(rawUUID)
			if err != nil {
				return fmt.Errorf("failed to parse UUID -> %v", err)
			}

			user, err := dao.GetUser(uuid)
			if err != nil {
				return fmt.Errorf("failed to get user -> %v", err)
			}

			responses.SendDataResponse(ctx, user)
		}
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

		oldUser, err := dao.GetUser(uuid)
		if err != nil {
			return fmt.Errorf("failed to get user -> %v", err)
		}

		// real old password salt from database
		oldPasswordSalt, err := encoding.Base64RawStdDecodeString(oldUser.PasswordSalt)
		if err != nil {
			return fmt.Errorf("failed to base64 decode old password salt -> %v", err)
		}
		// password hash produced from "old_password" form value + real old password salt
		passwordHash, err := security.HashPassword(
			oldPassword,
			oldPasswordSalt,
		)
		if err != nil {
			return fmt.Errorf("failed to hash inputted old password -> %v", err)
		}

		// real old password hash from database
		oldPasswordHash, err := encoding.Base64RawStdDecodeString(oldUser.PasswordHash)
		if err != nil {
			return fmt.Errorf("failed to base64 decode old password hash -> %v", err)
		}

		// compare produced password hash and real password hash
		if bytes.Equal(passwordHash, oldPasswordHash) {
			newUser, err := models.NewUser(oldUser.UUID, newPassword)
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

		user, err := dao.GetUser(uuid)
		if err != nil {
			return fmt.Errorf("failed to get user -> %v", err)
		}

		passwordSalt, err := encoding.Base64RawStdDecodeString(user.PasswordSalt)
		if err != nil {
			return fmt.Errorf("failed to base64 decode password salt -> %v", err)
		}

		// password hash produced from "old_password" form value + real old password salt
		passwordHash, err := security.HashPassword(password, passwordSalt)
		if err != nil {
			return fmt.Errorf("failed to hash password -> %v", err)
		}

		// real old password hash from database
		oldPasswordHash, err := encoding.Base64RawStdDecodeString(user.PasswordHash)
		if err != nil {
			return fmt.Errorf("failed to base64 decode old password hash -> ")
		}

		// compare produced password hash and real password hash
		if bytes.Equal(passwordHash, oldPasswordHash) {
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
