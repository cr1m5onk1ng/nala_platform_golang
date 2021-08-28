package controllers

import (
	"database/sql"
	"strconv"

	db "github.com/cr1m5onk1ng/nala_platform_app/db/sqlc"
	"github.com/cr1m5onk1ng/nala_platform_app/domain"
	"github.com/cr1m5onk1ng/nala_platform_app/validation"
	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
)

type PostDiscussioneResponse struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type PostDiscussionCommentResponse struct {
	Content string `json:"content"`
}

// UTILS
func parseDiscussionData(ctx *fiber.Ctx) (*domain.MappedPostDiscussion, error) {
	discussionData, err := validation.ValidateBodyData(ctx, &domain.MappedPostDiscussion{})
	if err != nil {
		return nil, err
	}
	discussion, ok := discussionData.(*domain.MappedPostDiscussion)

	if !ok {
		return nil, SendFailureResponse(
			ctx,
			fiber.StatusInternalServerError,
			"Could not cast input data to data model",
		)
	}
	return discussion, nil
}

func parseCommentData(ctx *fiber.Ctx) (*db.AddCommentParams, error) {
	commentData, err := validation.ValidateBodyData(ctx, &db.AddCommentParams{})
	if err != nil {
		return nil, err
	}

	comment, ok := commentData.(*db.AddCommentParams)

	if !ok {
		return nil, SendFailureResponse(
			ctx,
			fiber.StatusInternalServerError,
			"Could not cast input data to data model",
		)
	}
	return comment, nil
}

// CONTROLLERS

func (h *Handlers) CreateDiscussion(ctx *fiber.Ctx) error {
	discussionData, err := parseDiscussionData(ctx)
	if err != nil {
		return err
	}

	args := db.AddPostDiscussionParams{
		CreatorID:   discussionData.CreatorID,
		PostID:      discussionData.PostID,
		Title:       discussionData.Title,
		Description: sql.NullString{String: discussionData.Description, Valid: true},
	}

	addedDiscussion, err := h.Repo.AddPostDiscussion(ctx.Context(), args)
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
	return SendSuccessResponse(ctx, fiber.StatusOK, PostDiscussioneResponse{
		Title:       addedDiscussion.Title,
		Description: addedDiscussion.Description.String,
	},
	)
}

func (h *Handlers) RemoveDiscussion(ctx *fiber.Ctx) error {
	discussionId, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil {
		return SendFailureResponse(ctx, fiber.StatusInternalServerError, err.Error())
	}

	_, err = h.Repo.GetPostDiscussionById(ctx.Context(), discussionId)

	if err != nil {
		SendFailureResponse(
			ctx,
			fiber.StatusNotFound,
			err.Error(),
		)
	}

	err = h.Repo.RemovePostDiscussion(ctx.Context(), discussionId)
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

func (h *Handlers) AddCommentOnDiscussion(ctx *fiber.Ctx) error {
	commentData, err := parseCommentData(ctx)
	if err != nil {
		return err
	}

	args := db.AddCommentParams{
		DiscussionID:    commentData.DiscussionID,
		ParentCommentID: commentData.ParentCommentID,
		UserID:          commentData.UserID,
		Content:         commentData.Content,
	}

	addedComment, err := h.Repo.AddComment(ctx.Context(), args)
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
	return SendSuccessResponse(ctx, fiber.StatusOK, PostDiscussionCommentResponse{
		Content: addedComment.Content,
	},
	)
}

func (h *Handlers) RemoveCommentFromDiscussion(ctx *fiber.Ctx) error {
	commentId, err := strconv.ParseInt(ctx.Params("comment_id"), 10, 64)
	if err != nil {
		return SendFailureResponse(ctx, fiber.StatusInternalServerError, err.Error())
	}

	_, err = h.Repo.GetDiscussionCommentById(ctx.Context(), commentId)
	if err != nil {
		SendFailureResponse(
			ctx,
			fiber.StatusNotFound,
			err.Error(),
		)
	}

	err = h.Repo.RemoveDiscussionComments(ctx.Context(), commentId)
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

func (h *Handlers) GetPostDiscussions(ctx *fiber.Ctx) error {
	postId := ctx.Params("post_id")

	discussions, err := h.Repo.GetPostDiscussions(ctx.Context(), postId)
	if err != nil {
		return SendFailureResponse(ctx, fiber.StatusNotFound, err.Error())
	}

	return SendSuccessResponse(ctx, fiber.StatusOK, discussions)
}

func (h *Handlers) GetDiscussionComments(ctx *fiber.Ctx) error {
	discussionId, err := strconv.ParseInt(ctx.Params("discussion_id"), 10, 64)
	if err != nil {
		return SendFailureResponse(ctx, fiber.StatusInternalServerError, err.Error())
	}
	comments, err := h.Repo.GetAllDiscussionComments(ctx.Context(), discussionId)
	if err != nil {
		return SendFailureResponse(ctx, fiber.StatusNotFound, err.Error())
	}
	return SendSuccessResponse(ctx, fiber.StatusOK, comments)
}