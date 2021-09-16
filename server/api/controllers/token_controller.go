package controllers

import (
	"database/sql"
	"os"
	"strconv"
	"time"

	db "github.com/cr1m5onk1ng/nala_platform_app/db/sqlc"
	"github.com/cr1m5onk1ng/nala_platform_app/domain"
	"github.com/cr1m5onk1ng/nala_platform_app/validation"
	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
)

// GetNewAccessToken method for create a new access token.
// @Description Create a new access token.
// @Summary create a new access token
// @Tags Token
// @Accept json
// @Produce json
// @Success 200 {string} status "ok"
// @Router /v1/token/new [get]

type RefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (h *Handlers) GetNewJwtAccessToken(c *fiber.Ctx) error {
	// Generate a new Access token.
	secret := os.Getenv("JWT_SECRET_KEY")
	minutesCount, err := strconv.Atoi(os.Getenv("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"))
	if err != nil {
		return err
	}
	token, err := validation.GenerateNewAccessToken(secret, minutesCount)
	if err != nil {
		// Return status 500 and token generation error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"error":        false,
		"msg":          nil,
		"access_token": token,
	})
}

func (h *Handlers) GetNewPasetoAccessToken(ctx *fiber.Ctx) error {
	var tokenReq = &domain.RefreshTokenRequest{}
	err := ctx.BodyParser(tokenReq)
	if err != nil {
		return SendFailureResponse(ctx, fiber.StatusBadRequest, err.Error())
	}

	user, err := h.Repo.GetUserByEmail(ctx.Context(), tokenReq.Email)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			return SendFailureResponse(ctx, fiber.StatusForbidden, pqErr.Error())
		}
		return SendFailureResponse(ctx, fiber.StatusInternalServerError, err.Error())
	}

	if !user.AccessToken.Valid || user.AccessToken.String != tokenReq.Token {
		return SendFailureResponse(ctx, fiber.StatusForbidden, "invalid token")
	}

	err = h.Repo.InvalidateToken(ctx.Context(), tokenReq.Token)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			return SendFailureResponse(ctx, fiber.StatusForbidden, pqErr.Error())
		}
		return SendFailureResponse(ctx, fiber.StatusInternalServerError, err.Error())
	}

	newAccessToken, err := h.TokenManager.CreateToken(
		tokenReq.Username,
		tokenReq.Email,
		h.Config.PASETO_TOKEN_DURATION,
	)
	if err != nil {
		return SendFailureResponse(ctx, fiber.StatusInternalServerError, err.Error())
	}

	newRefreshToken, err := h.TokenManager.CreateRefreshToken(tokenReq.Username, tokenReq.Email)
	if err != nil {
		return SendFailureResponse(ctx, fiber.StatusInternalServerError, err.Error())
	}

	_, err = h.Repo.AddToken(
		ctx.Context(),
		db.AddTokenParams{Token: newRefreshToken, ExpiredAt: time.Now().Add(h.Config.PASETO_TOKEN_DURATION)},
	)
	if err != nil {
		SendFailureResponse(ctx, fiber.StatusInternalServerError, err.Error())
	}

	_, err = h.Repo.UpdateUserToken(ctx.Context(), db.UpdateUserTokenParams{
		Email:       tokenReq.Email,
		AccessToken: sql.NullString{String: newRefreshToken, Valid: true},
	})
	if err != nil {
		SendFailureResponse(ctx, fiber.StatusInternalServerError, err.Error())
	}

	return SendSuccessResponse(
		ctx,
		fiber.StatusOK,
		RefreshTokenResponse{AccessToken: newAccessToken, RefreshToken: newRefreshToken},
	)
}
