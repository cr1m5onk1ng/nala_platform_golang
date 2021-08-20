package controllers

import (
	"database/sql"

	"github.com/cr1m5onk1ng/nala_platform_app/constants"
	db "github.com/cr1m5onk1ng/nala_platform_app/db/sqlc"
	"github.com/cr1m5onk1ng/nala_platform_app/validation"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handlers) GetPost(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("postId"))
	if err != nil {
		return SendFailureResponse(
			ctx,
			fiber.StatusInternalServerError,
			err.Error(),
		)
	}

	post, err := h.Repo.GetPostById(ctx.Context(), id.String())
	if err != nil {
		return SendFailureResponse(
			ctx,
			fiber.StatusNotFound,
			err.Error(),
		)
	}

	return SendSuccessResponse(
		ctx,
		fiber.StatusOK,
		post,
	)
}

func (h *Handlers) GetPostsByLanguage(ctx *fiber.Ctx) error {
	lang := ctx.Params("lang")
	langOk := validation.IsLanguageStringValid(lang)
	if !langOk {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Specified language is invalid: ",
			"data":    nil,
		})
	}

	args := db.GetPostsByLanguageParams{
		Language: lang,
		Limit:    constants.LIMIT,
		Offset:   constants.OFFSET,
	}
	posts, err := h.Repo.GetPostsByLanguage(ctx.Context(), args)

	if err != nil {
		return SendFailureResponse(
			ctx,
			fiber.StatusNotFound,
			err.Error(),
		)
	}

	return SendSuccessResponse(
		ctx,
		fiber.StatusOK,
		posts,
	)

}

func (h *Handlers) GetPostsByDifficulty(ctx *fiber.Ctx) error {
	difficulty := ctx.Params("diff")
	// Add Difficulty string validation

	nullableDifficulty := sql.NullString{
		String: difficulty,
		Valid:  true,
	}

	args := db.GetPostsByDifficultyParams{
		Difficulty: nullableDifficulty,
		Limit:      constants.LIMIT,
		Offset:     constants.OFFSET,
	}

	posts, err := h.Repo.GetPostsByDifficulty(ctx.Context(), args)
	if err != nil {
		return SendFailureResponse(
			ctx,
			fiber.StatusNotFound,
			err.Error(),
		)
	}

	return SendSuccessResponse(
		ctx,
		fiber.StatusOK,
		posts,
	)
}

func (h *Handlers) GetPostsByUser(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("usrId"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "An error has occured: " + err.Error(),
			"data":    nil,
		})
	}

	args := db.GetPostsByUserParams{
		UserID: id.String(),
		Limit:  constants.LIMIT,
		Offset: constants.OFFSET,
	}

	posts, err := h.Repo.GetPostsByUser(ctx.Context(), args)
	if err != nil {
		return SendFailureResponse(
			ctx,
			fiber.StatusNotFound,
			err.Error(),
		)
	}

	return SendSuccessResponse(
		ctx,
		fiber.StatusOK,
		posts,
	)
}

func (h *Handlers) GetPostsByCategory(ctx *fiber.Ctx) error {
	cat := ctx.Params("cat")

	args := db.GetPostsByCategoryParams{
		Category: cat,
		Limit:    constants.LIMIT,
		Offset:   constants.OFFSET,
	}

	posts, err := h.Repo.GetPostsByCategory(ctx.Context(), args)
	if err != nil {
		return SendFailureResponse(
			ctx,
			fiber.StatusNotFound,
			err.Error(),
		)
	}

	return SendSuccessResponse(
		ctx,
		fiber.StatusOK,
		posts,
	)
}

func (h *Handlers) GetPostsByMediaType(ctx *fiber.Ctx) error {
	mediaType := ctx.Params("media")
	nullableMediaType := sql.NullString{
		String: mediaType,
		Valid:  true,
	}

	args := db.GetPostsByMediaTypeParams{
		MediaType: nullableMediaType,
		Limit:     constants.LIMIT,
		Offset:    constants.OFFSET,
	}

	posts, err := h.Repo.GetPostsByMediaType(ctx.Context(), args)
	if err != nil {
		return SendFailureResponse(
			ctx,
			fiber.StatusNotFound,
			err.Error(),
		)
	}

	return SendSuccessResponse(
		ctx,
		fiber.StatusOK,
		posts,
	)
}

func (h *Handlers) AddPost(ctx *fiber.Ctx) error {
	post, err := validation.ValidatePostData(ctx, &db.UserPost{})
	if err != nil {
		return SendFailureResponse(
			ctx,
			fiber.StatusNotFound,
			"an error occured: "+err.Error(),
		)
	}

	args := db.AddPostParams{
		ID:              post.ID,
		UserID:          post.UserID,
		ResourceID:      post.ResourceID,
		PostTitle:       post.PostTitle,
		PostDescription: post.PostDescription,
	}
	addedPost, err := h.Repo.AddPost(ctx.Context(), args)

	if err != nil {
		return SendFailureResponse(
			ctx,
			fiber.StatusInternalServerError,
			err.Error(),
		)
	}

	return SendSuccessResponse(
		ctx,
		fiber.StatusOK,
		addedPost,
	)
}

func (h *Handlers) UpdatePost(ctx *fiber.Ctx) error {
	post, err := validation.ValidatePostData(ctx, &db.UserPost{})
	if err != nil {
		return SendFailureResponse(
			ctx,
			fiber.StatusNotFound,
			"an error occured: "+err.Error(),
		)
	}

	args := db.UpdatePostParams{
		UserID:          post.UserID,
		ResourceID:      post.ResourceID,
		PostTitle:       post.PostTitle,
		PostDescription: post.PostDescription,
	}
	editedPost, err := h.Repo.UpdatePostTrans(ctx.Context(), args)

	if err != nil {
		return SendFailureResponse(
			ctx,
			fiber.StatusInternalServerError,
			err.Error(),
		)
	}

	return SendSuccessResponse(
		ctx,
		fiber.StatusOK,
		editedPost,
	)
}

func (h *Handlers) GetPostTags(ctx *fiber.Ctx) error {
	postId, err := uuid.Parse(ctx.Params("postId"))
	if err != nil {
		SendFailureResponse(
			ctx,
			fiber.StatusInternalServerError,
			"An error has occured: "+err.Error(),
		)
	}

	tags, err := h.Repo.GetPostTags(ctx.Context(), postId.String())

	if err != nil {
		return SendFailureResponse(
			ctx,
			fiber.StatusNotFound,
			err.Error(),
		)
	}

	return SendSuccessResponse(
		ctx,
		fiber.StatusOK,
		tags,
	)

}

func (h *Handlers) GetPostDifficultyVotes(ctx *fiber.Ctx) error {
	postId, err := uuid.Parse(ctx.Params("postId"))
	if err != nil {
		SendFailureResponse(
			ctx,
			fiber.StatusInternalServerError,
			"An error has occured: "+err.Error(),
		)
	}

	votes, err := h.Repo.GetPostDifficultyVotes(ctx.Context(), postId.String())

	if err != nil {
		return SendFailureResponse(
			ctx,
			fiber.StatusNotFound,
			err.Error(),
		)
	}

	return SendSuccessResponse(
		ctx,
		fiber.StatusOK,
		votes,
	)
}

func (h *Handlers) GetPostLikes(ctx *fiber.Ctx) error {
	postId, err := uuid.Parse(ctx.Params("postId"))
	if err != nil {
		SendFailureResponse(
			ctx,
			fiber.StatusInternalServerError,
			"An error has occured: "+err.Error(),
		)
	}

	likes, err := h.Repo.GetPostLikes(ctx.Context(), postId.String())

	if err != nil {
		return SendFailureResponse(
			ctx,
			fiber.StatusNotFound,
			err.Error(),
		)
	}

	return SendSuccessResponse(
		ctx,
		fiber.StatusOK,
		likes,
	)
}

func (h *Handlers) RemovePost(ctx *fiber.Ctx) error {
	postId, err := uuid.Parse(ctx.Params("postId"))
	if err != nil {
		SendFailureResponse(
			ctx,
			fiber.StatusInternalServerError,
			"An error has occured: "+err.Error(),
		)
	}

	_, err = h.Repo.GetPostById(ctx.Context(), postId.String())
	if err != nil {
		SendFailureResponse(
			ctx,
			fiber.StatusNotFound,
			err.Error(),
		)
	}

	err = h.Repo.RemovePost(ctx.Context(), postId.String())
	if err != nil {
		SendFailureResponse(
			ctx,
			fiber.StatusInternalServerError,
			"an error has occured: "+err.Error(),
		)
	}

	return SendSuccessResponse(
		ctx,
		fiber.StatusNoContent,
		nil,
	)
}
