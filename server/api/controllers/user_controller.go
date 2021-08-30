package controllers

import (
	"database/sql"
	"fmt"

	db "github.com/cr1m5onk1ng/nala_platform_app/db/sqlc"
	"github.com/cr1m5onk1ng/nala_platform_app/domain"
	"github.com/cr1m5onk1ng/nala_platform_app/util"
	"github.com/cr1m5onk1ng/nala_platform_app/validation"
	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
)

// We don't want to return sensible data like password in the response
// so we create a customized response
type userResponse struct {
	Username       string `json:"username"`
	Email          string `json:"email"`
	NativeLanguage string `json:"native_language"`
}

type loginUserResponse struct {
	AccessToken string `json:"access_token"`
}

func (h *Handlers) checkUserPermission(ctx *fiber.Ctx, id string) (db.User, error) {
	payload := ctx.Locals(util.AuthPayloadKey).(*validation.Payload)
	user, err := h.Repo.GetUser(ctx.Context(), id)
	if err != nil {
		return user, err
	}
	if payload.Username != user.Username {
		return user, validation.ErrInvalidToken
	}
	return user, nil
}

func handleUserAuthError(ctx *fiber.Ctx, err error) error {
	if err.Error() == validation.ErrInvalidToken.Error() {
		return SendFailureResponse(ctx, fiber.StatusUnauthorized, err.Error())
	}
	return SendFailureResponse(ctx, fiber.StatusInternalServerError, err.Error())
}

func (h *Handlers) LoginUser(ctx *fiber.Ctx) error {
	loginReq, err := validation.CheckLoginDataValidity(ctx, &domain.LoginRequest{})
	if err != nil {
		return SendFailureResponse(ctx, fiber.StatusBadRequest, err.Error())
	}
	user, err := h.Repo.GetUserByEmail(ctx.Context(), loginReq.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return SendFailureResponse(ctx, fiber.StatusNotFound, err.Error())
		}
		return SendFailureResponse(ctx, fiber.StatusInternalServerError, err.Error())
	}
	err = util.CheckPassword(loginReq.Password, user.HashedPassword)
	if err != nil {
		return SendFailureResponse(ctx, fiber.StatusUnauthorized, err.Error())
	}
	fmt.Printf("Token Duration: %s", h.Config.PASETO_TOKEN_DURATION)
	token, err := h.TokenManager.CreateToken(user.Username, user.Email, h.Config.PASETO_TOKEN_DURATION)
	if err != nil {
		return SendFailureResponse(ctx, fiber.StatusInternalServerError, err.Error())
	}
	response := loginUserResponse{
		AccessToken: token,
	}
	return SendSuccessResponse(ctx, fiber.StatusOK, response)
}

func (h *Handlers) GetUser(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	id, err := util.ParseRequestParam(idParam)
	if err != nil {
		return SendFailureResponse(ctx, fiber.StatusInternalServerError, err.Error())
	}
	user, err := h.checkUserPermission(ctx, id)
	if err != nil {
		return handleUserAuthError(ctx, err)
	}
	return SendSuccessResponse(ctx, fiber.StatusOK, user)
}

func (h *Handlers) GetAllUsers(ctx *fiber.Ctx) error {
	users, err := h.Repo.GetAllUsers(ctx.Context())
	if err != nil {
		return SendFailureResponse(ctx, fiber.StatusInternalServerError, err.Error())
	}
	return SendSuccessResponse(ctx, fiber.StatusOK, users)
}

func (h *Handlers) CreateUser(ctx *fiber.Ctx) error {

	user, err := validation.CheckUserDataValidty(ctx, &domain.MappedUser{})
	if err != nil {
		return err
	}

	args := db.CreateUserParams{
		ID:             user.ID,
		Username:       user.Username,
		Email:          user.Email,
		HashedPassword: user.Password,
		NativeLanguage: user.NativeLanguage,
		Role:           sql.NullString{String: user.Role, Valid: true},
	}

	insertedUser, err := h.Repo.CreateUser(ctx.Context(), args)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			return SendFailureResponse(ctx, fiber.StatusForbidden, pqErr.Error())
		}
		return SendFailureResponse(ctx, fiber.StatusInternalServerError, err.Error())
	}

	return SendSuccessResponse(
		ctx,
		fiber.StatusCreated,
		userResponse{
			Username:       insertedUser.Username,
			Email:          insertedUser.Email,
			NativeLanguage: insertedUser.NativeLanguage,
		},
	)

}

func (h *Handlers) UpdateUserLanguage(ctx *fiber.Ctx) error {

	userId := ctx.Params("id")

	_, err := h.checkUserPermission(ctx, userId)
	if err != nil {
		return handleUserAuthError(ctx, err)
	}

	langData, err := validation.ValidateBodyData(ctx, &db.UpdateUserLanguageParams{})
	if err != nil {
		return err
	}

	lang, ok := langData.(*db.UpdateUserLanguageParams)
	if !ok {
		return SendFailureResponse(ctx, fiber.StatusInternalServerError, "failed to parse request body")
	}

	args := db.UpdateUserLanguageParams{
		ID:             userId,
		NativeLanguage: lang.NativeLanguage,
	}

	updatedUser, err := h.Repo.UpdateUserLanguageTrans(ctx.Context(), args)

	if err != nil {
		return SendFailureResponse(ctx, fiber.StatusInternalServerError, err.Error())
	}
	return SendSuccessResponse(
		ctx,
		fiber.StatusCreated,
		userResponse{
			Username:       updatedUser.Username,
			Email:          updatedUser.Email,
			NativeLanguage: updatedUser.NativeLanguage,
		},
	)

}

func (h *Handlers) UpdateUserRole(ctx *fiber.Ctx) error {

	userId := ctx.Params("id")

	_, err := h.checkUserPermission(ctx, userId)
	if err != nil {
		return handleUserAuthError(ctx, err)
	}

	roleData, err := validation.ValidateBodyData(ctx, &db.UpdateUserRoleParams{})
	if err != nil {
		return err
	}

	role, ok := roleData.(*db.UpdateUserRoleParams)
	if !ok {
		return SendFailureResponse(ctx, fiber.StatusInternalServerError, "failed to parse request body")
	}

	args := db.UpdateUserRoleParams{
		ID:   userId,
		Role: sql.NullString{String: role.Role.String, Valid: true},
	}

	updatedUser, err := h.Repo.UpdateUserRoleTrans(ctx.Context(), args)

	if err != nil {
		return SendFailureResponse(ctx, fiber.StatusInternalServerError, err.Error())
	}

	return SendSuccessResponse(
		ctx,
		fiber.StatusCreated,
		userResponse{
			Username:       updatedUser.Username,
			Email:          updatedUser.Email,
			NativeLanguage: updatedUser.NativeLanguage,
		},
	)

}

func (h *Handlers) AddUserTargetLanguage(ctx *fiber.Ctx) error {
	targetLang, err := validation.CheckTargetLanguageDataValidity(ctx, &db.Learning{})
	if err != nil {
		return SendFailureResponse(ctx, fiber.StatusBadRequest, err.Error())
	}

	/*
		_, err = h.checkUserPermission(ctx, targetLang.UserID)
		if err != nil {
			return err
		} */
	_, err = h.checkUserPermission(ctx, targetLang.UserID)
	if err != nil {
		return handleUserAuthError(ctx, err)
	}

	args := db.AddUserTargetLanguageParams{
		UserID:   targetLang.UserID,
		Language: targetLang.Language,
	}

	addedLanguage, err := h.Repo.AddUserTargetLanguage(ctx.Context(), args)
	if err != nil {
		return SendFailureResponse(ctx, fiber.StatusInternalServerError, err.Error())
	}
	return SendSuccessResponse(ctx, fiber.StatusCreated, addedLanguage)

}
