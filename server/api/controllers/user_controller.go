package controllers

import (
	db "github.com/cr1m5onk1ng/nala_platform_app/db/sqlc"
	"github.com/cr1m5onk1ng/nala_platform_app/validation"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handlers) GetUser(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "An error has occured: " + err.Error(),
			"data":    nil,
		})
	}
	user, err := h.Repo.GetUser(ctx.Context(), id.String())
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "resource with the given ID was not found",
			"data":    nil,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": nil,
		"data":    user,
	})
}

func (h *Handlers) GetAllUsers(ctx *fiber.Ctx) error {

	users, err := h.Repo.GetAllUsers(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
			"data":    nil,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": nil,
		"data":    users,
	})
}

func (h *Handlers) CreateUser(ctx *fiber.Ctx) error {
	user, err := validation.CheckUserDataValidtyAndAuthorization(ctx, &db.User{})
	if err != nil {
		return err
	}

	args := db.CreateUserParams{
		ID:             user.ID,
		Username:       user.Username,
		NativeLanguage: user.NativeLanguage,
		Role:           user.Role,
	}

	insertedUser, err := h.Repo.CreateUser(ctx.Context(), args)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
			"data":    nil,
		})
	}

	return ctx.JSON(fiber.Map{
		"error":   false,
		"message": nil,
		"data":    insertedUser,
	})

}

func (h *Handlers) UpdateUserLanguage(ctx *fiber.Ctx) error {

	user, err := validation.CheckUserDataValidtyAndAuthorization(ctx, &db.User{})
	if err != nil {
		return err
	}

	_, err = h.Repo.GetUser(ctx.Context(), user.ID)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "user with specified ID not found",
			"data":    nil,
		})
	}

	args := db.UpdateUserLanguageParams{
		ID:             user.ID,
		NativeLanguage: user.NativeLanguage,
	}

	updatedUser, err := h.Repo.UpdateUserLanguageTrans(ctx.Context(), args)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
			"data":    nil,
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":   false,
		"message": nil,
		"data":    updatedUser,
	})

}

func (h *Handlers) UpdateUserRole(ctx *fiber.Ctx) error {

	user, err := validation.CheckUserDataValidtyAndAuthorization(ctx, &db.User{})
	if err != nil {
		return err
	}

	_, err = h.Repo.GetUser(ctx.Context(), user.ID)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "user with specified ID not found",
			"data":    nil,
		})
	}

	args := db.UpdateUserRoleParams{
		ID:   user.ID,
		Role: user.Role,
	}

	updatedUser, err := h.Repo.UpdateUserRoleTrans(ctx.Context(), args)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
			"data":    nil,
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":   false,
		"message": nil,
		"data":    updatedUser,
	})

}
