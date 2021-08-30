package controllers

import (
	"database/sql"

	"github.com/cr1m5onk1ng/nala_platform_app/constants"
	db "github.com/cr1m5onk1ng/nala_platform_app/db/sqlc"
	"github.com/cr1m5onk1ng/nala_platform_app/domain"
	"github.com/cr1m5onk1ng/nala_platform_app/util"
	"github.com/cr1m5onk1ng/nala_platform_app/validation"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

// UTILITY FUNCTIONS

func parsePostData(ctx *fiber.Ctx) (*domain.MappedUserPost, error) {
	postData, err := validation.ValidatePostData(ctx, &domain.MappedUserPost{})
	if err != nil {
		return nil, err
	}
	return postData, nil
}

func parseVoteData(ctx *fiber.Ctx) (*db.VotePostParams, error) {
	vote, err := validation.ValidateVoteData(ctx, &db.VotePostParams{})
	if err != nil {
		return nil, err
	}
	return vote, nil
}

// CONTROLLERS

func (h *Handlers) GetPost(ctx *fiber.Ctx) error {
	id, err := util.ParseRequestParam(ctx.Params("post-id"))
	if err != nil {
		return SendFailureResponse(
			ctx,
			fiber.StatusInternalServerError,
			err.Error(),
		)
	}

	post, err := h.Repo.GetPostById(ctx.Context(), id)
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
	lang, err := util.ParseRequestParam(ctx.Params("lang"))
	if err != nil {
		SendFailureResponse(ctx, fiber.StatusInternalServerError, err.Error())
	}
	langOk := validation.IsLanguageStringValid(lang)
	if !langOk {
		SendFailureResponse(
			ctx,
			fiber.StatusBadRequest,
			"Specified language is invalid",
		)
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
	difficulty, err := util.ParseRequestParam(ctx.Params("diff"))
	if err != nil {
		SendFailureResponse(ctx, fiber.StatusInternalServerError, err.Error())
	}

	args := db.GetPostsByDifficultyParams{
		Difficulty: difficulty,
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
	id, err := uuid.Parse(ctx.Params("usr-id"))
	if err != nil {
		return SendFailureResponse(ctx, fiber.StatusInternalServerError, err.Error())
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
	cat, err := util.ParseRequestParam(ctx.Params("cat"))
	if err != nil {
		return SendFailureResponse(ctx, fiber.StatusInternalServerError, err.Error())
	}

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
	mediaType, err := util.ParseRequestParam(ctx.Params("media"))
	if err != nil {
		return SendFailureResponse(ctx, fiber.StatusInternalServerError, err.Error())
	}

	args := db.GetPostsByMediaTypeParams{
		MediaType: mediaType,
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
	post, err := parsePostData(ctx)
	if err != nil {
		return err
	}

	args := db.AddPostParams{
		ID:              post.ID,
		UserID:          post.UserID,
		ResourceID:      post.ResourceID,
		PostTitle:       post.PostTitle,
		PostDescription: sql.NullString{String: post.PostDescription, Valid: true},
	}

	addedPost, err := h.Repo.AddPost(ctx.Context(), args)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			return SendFailureResponse(ctx, fiber.StatusForbidden, pqErr.Error())
		}
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

func (h *Handlers) AddPostNotSecure(ctx *fiber.Ctx) error {
	post, err := validation.ValidatePostData(ctx, &domain.MappedUserPost{})
	if err != nil {
		return err
	}

	args := db.AddPostParams{
		ID:              post.ID,
		UserID:          post.UserID,
		ResourceID:      post.ResourceID,
		PostTitle:       post.PostTitle,
		PostDescription: sql.NullString{String: post.PostDescription, Valid: true},
	}

	addedPost, err := h.Repo.AddPost(ctx.Context(), args)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			return SendFailureResponse(ctx, fiber.StatusForbidden, pqErr.Error())
		}
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
	post, err := parsePostData(ctx)
	if err != nil {
		return err
	}

	_, err = h.checkUserPermission(ctx, post.UserID)
	if err != nil {
		handleUserAuthError(ctx, err)
	}

	args := db.UpdatePostParams{
		UserID:          post.UserID,
		ResourceID:      post.ResourceID,
		PostTitle:       post.PostTitle,
		PostDescription: sql.NullString{String: post.PostDescription, Valid: true},
	}
	editedPost, err := h.Repo.UpdatePostTrans(ctx.Context(), args)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			return SendFailureResponse(ctx, fiber.StatusForbidden, pqErr.Error())
		}
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
	postId := ctx.Params("post-id")

	tags, err := h.Repo.GetPostTags(ctx.Context(), postId)

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
	postId := ctx.Params("post_id")

	votes, err := h.Repo.GetPostDifficultyVotes(ctx.Context(), postId)

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
	postId := ctx.Params("post-id")

	likes, err := h.Repo.GetPostLikes(ctx.Context(), postId)

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
	postId := ctx.Params("post-id")

	postData, err := h.Repo.GetPostById(ctx.Context(), postId)
	if err != nil {
		SendFailureResponse(
			ctx,
			fiber.StatusNotFound,
			err.Error(),
		)
	}

	_, err = h.checkUserPermission(ctx, postData.UserID)
	if err != nil {
		return handleUserAuthError(ctx, err)
	}

	err = h.Repo.RemovePost(ctx.Context(), postId)
	if err != nil {
		SendFailureResponse(
			ctx,
			fiber.StatusInternalServerError,
			err.Error(),
		)
	}

	return SendSuccessResponse(
		ctx,
		fiber.StatusNoContent,
		nil,
	)
}

func (h *Handlers) VotePost(ctx *fiber.Ctx) error {
	vote, err := parseVoteData(ctx)
	if err != nil {
		return err
	}

	args := db.VotePostParams{
		UserID:     vote.UserID,
		PostID:     vote.PostID,
		Difficulty: vote.Difficulty,
	}

	addedVote, err := h.Repo.VotePost(ctx.Context(), args)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			return SendFailureResponse(ctx, fiber.StatusForbidden, pqErr.Error())
		}
		return SendFailureResponse(
			ctx,
			fiber.StatusInternalServerError,
			err.Error(),
		)
	}

	return SendSuccessResponse(
		ctx,
		fiber.StatusOK,
		addedVote,
	)
}

func (h *Handlers) UpdateVote(ctx *fiber.Ctx) error {
	vote, err := parseVoteData(ctx)
	if err != nil {
		return err
	}

	_, err = h.checkUserPermission(ctx, vote.UserID)
	if err != nil {
		return handleUserAuthError(ctx, err)
	}

	args := db.UpdateVoteParams{
		UserID:     vote.UserID,
		PostID:     vote.PostID,
		Difficulty: vote.Difficulty,
	}
	editedVote, err := h.Repo.UpdateVote(ctx.Context(), args)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			return SendFailureResponse(ctx, fiber.StatusForbidden, pqErr.Error())
		}
		return SendFailureResponse(
			ctx,
			fiber.StatusInternalServerError,
			err.Error(),
		)
	}

	return SendSuccessResponse(
		ctx,
		fiber.StatusOK,
		editedVote,
	)
}

func (h *Handlers) RemoveVote(ctx *fiber.Ctx) error {
	vote, err := parseVoteData(ctx)
	if err != nil {
		return err
	}

	_, err = h.checkUserPermission(ctx, vote.UserID)
	if err != nil {
		return handleUserAuthError(ctx, err)
	}

	removeVoteArgs := db.RemoveVoteParams{
		UserID: vote.UserID,
		PostID: vote.PostID,
	}

	err = h.Repo.RemoveVote(ctx.Context(), removeVoteArgs)

	if err != nil {
		SendFailureResponse(
			ctx,
			fiber.StatusInternalServerError,
			err.Error(),
		)
	}

	return SendSuccessResponse(
		ctx,
		fiber.StatusNoContent,
		nil,
	)
}
