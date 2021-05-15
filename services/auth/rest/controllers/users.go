package controllers

import (
	"bytes"
	"errors"
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
				return err
			}

			responses.SendDataResponse(ctx, &users)
		} else {
			uuid, err := uuid.Parse(rawUUID)
			if err != nil {
				return err
			}

			user, err := dao.GetUser(uuid)
			if err != nil {
				return err
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
			return err
		}

		user, err := models.NewUser(uuid, rawPassword)
		if err != nil {
			return err
		}

		dao.CreateUser(user)
	case http.MethodPatch:
		rawUUID := ctx.FormValue("uuid")
		oldPassword := ctx.FormValue("old_password")
		newPassword := ctx.FormValue("new_password")

		if rawUUID == "" || oldPassword == "" || newPassword == "" {
			return errors.New("missing form value uuid, old_password or new_password")
		}

		uuid, err := uuid.Parse(rawUUID)
		if err != nil {
			return err
		}

		oldUser, err := dao.GetUser(uuid)
		if err != nil {
			return err
		}

		// real old password salt from database
		oldPasswordSalt, err := encoding.Base64RawStdDecodeString(oldUser.PasswordSalt)
		if err != nil {
			return err
		}
		// password hash produced from "old_password" form value + real old password salt
		passwordHash, err := security.HashPassword(
			oldPassword,
			oldPasswordSalt,
		)
		if err != nil {
			return err
		}

		// real old password hash from database
		oldPasswordHash, err := encoding.Base64RawStdDecodeString(oldUser.PasswordHash)
		if err != nil {
			return err
		}

		// compare produced password hash and real password hash
		if bytes.Equal(passwordHash, oldPasswordHash) {
			newUser, err := models.NewUser(oldUser.UUID, newPassword)
			if err != nil {
				return err
			}

			return dao.UpdateUser(newUser)
		}
	case http.MethodDelete:
		rawUUID := ctx.FormValue("uuid")
		password := ctx.FormValue("password")

		if rawUUID == "" || password == "" {
			return errors.New("missing form value uuid or password")
		}

		uuid, err := uuid.Parse(rawUUID)
		if err != nil {
			return err
		}

		user, err := dao.GetUser(uuid)
		if err != nil {
			return err
		}

		passwordSalt, err := encoding.Base64RawStdDecodeString(user.PasswordHash)
		if err != nil {
			return err
		}

		// password hash produced from "old_password" form value + real old password salt
		passwordHash, err := security.HashPassword(password, passwordSalt)
		if err != nil {
			return err
		}

		// real old password hash from database
		oldPasswordHash, err := encoding.Base64RawStdDecodeString(user.PasswordHash)
		if err != nil {
			return err
		}

		// compare produced password hash and real password hash
		if bytes.Equal(passwordHash, oldPasswordHash) {
			return dao.DeleteUser(uuid)
		} else {
			return errors.New("incorrect password")
		}
	default:
		responses.SendUnsupportedMethodResponse(ctx)
	}

	return nil
}
