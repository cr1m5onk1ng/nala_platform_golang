package validation

import (
	"fmt"

	db "github.com/cr1m5onk1ng/nala_platform_app/db/sqlc"
	"github.com/cr1m5onk1ng/nala_platform_app/domain"
	"github.com/cr1m5onk1ng/nala_platform_app/util"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func ValidatePostData(ctx *fiber.Ctx, post *domain.MappedUserPost) (*domain.MappedUserPost, error) {
	if err := ctx.BodyParser(post); err != nil {
		return nil, err
	}

	post.ID = uuid.NewString()

	return post, nil
}

func ValidatePostDataAndAuthorization(ctx *fiber.Ctx, post *domain.MappedUserPost) (*domain.MappedUserPost, error) {
	var err error
	if err = CheckUserAuthorization(ctx); err != nil {
		return nil, err
	}

	p, err := ValidatePostData(ctx, post)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func isDifficultyWithinValues(difficultyVote string) bool {
	values := []string{"BEGINNER", "MEDIUM", "ADVANCED", "NATIVE"}
	return util.CheckStringInSlice(values, difficultyVote)
}

func ValidateVoteData(ctx *fiber.Ctx, vote *db.VotePostParams) (*db.VotePostParams, error) {
	if err := ctx.BodyParser(vote); err != nil {
		return nil, err
	}
	if !isDifficultyWithinValues(vote.Difficulty) {
		return nil, fmt.Errorf("invalid difficulty string")
	}
	return vote, nil
}
