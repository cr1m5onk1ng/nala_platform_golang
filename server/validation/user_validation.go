package validation

import (
	"time"

	db "github.com/cr1m5onk1ng/nala_platform_app/db/sqlc"
	"github.com/cr1m5onk1ng/nala_platform_app/domain"
	"github.com/cr1m5onk1ng/nala_platform_app/util"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CheckUserDataValidty(ctx *fiber.Ctx, user *domain.MappedUser) (*domain.MappedUser, error) {
	// Check if data is valid
	if err := ctx.BodyParser(user); err != nil {
		return nil, err
	}
	user.ID = uuid.NewString()
	hashedPass, err := util.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPass
	return user, nil
}

func CheckLoginDataValidity(ctx *fiber.Ctx, loginRequest *domain.LoginRequest) (*domain.LoginRequest, error) {
	if err := ctx.BodyParser(loginRequest); err != nil {
		return nil, err
	}
	return loginRequest, nil
}

func CheckUserAuthorization(ctx *fiber.Ctx) error {
	now := time.Now().Unix()

	claims, err := ExtractTokenMetadata(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
			"data":    nil,
		})
	}

	expires := claims.Expires

	if now > expires {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "unauthorized, token expired",
			"data":    nil,
		})
	}

	return nil
}

func CheckUserDataValidtyAndAuthorization(ctx *fiber.Ctx, user *domain.MappedUser) (*domain.MappedUser, error) {
	var err error
	var userToCheck *domain.MappedUser

	err = CheckUserAuthorization(ctx)
	if err != nil {
		return nil, err
	}
	userToCheck, err = CheckUserDataValidty(ctx, user)
	if err != nil {
		return nil, err
	}
	return userToCheck, nil
}

func CheckTargetLanguageDataValidity(ctx *fiber.Ctx, targetLang *db.Learning) (*db.Learning, error) {
	if err := ctx.BodyParser(targetLang); err != nil {
		return nil, err
	}
	return targetLang, nil
}
