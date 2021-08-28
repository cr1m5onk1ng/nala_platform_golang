package domain

import _ "github.com/go-playground/validator/v10"

type MappedPostDiscussion struct {
	CreatorID   string `json:"creator_id" validate:"required"`
	PostID      string `json:"post_id" validate:"required"`
	Title       string `json:"title" validate:"required,alphanum"`
	Description string `json:"description" validate:"required,alphanum"`
}

type MappedPostDiscussionComment struct {
	DiscussionID    int64  `json:"discussion_id" validate:"required"`
	ParentCommentID int64  `json:"parent_comment_id"`
	UserID          string `json:"user_id" validate:"required"`
	Content         string `json:"content" validate:"required,alphanum"`
}
