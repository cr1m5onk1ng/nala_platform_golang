// Code generated by sqlc. DO NOT EDIT.
// source: post.sql

package db

import (
	"context"
	"database/sql"
)

const addPost = `-- name: AddPost :one
INSERT INTO user_post (
  id,
  user_id, 
  resource_id, 
  post_title, 
  post_description 
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING id, user_id, resource_id, post_time, post_title, post_description
`

type AddPostParams struct {
	ID              string         `json:"id"`
	UserID          string         `json:"user_id"`
	ResourceID      int64          `json:"resource_id"`
	PostTitle       string         `json:"post_title"`
	PostDescription sql.NullString `json:"post_description"`
}

func (q *Queries) AddPost(ctx context.Context, arg AddPostParams) (UserPost, error) {
	row := q.queryRow(ctx, q.addPostStmt, addPost,
		arg.ID,
		arg.UserID,
		arg.ResourceID,
		arg.PostTitle,
		arg.PostDescription,
	)
	var i UserPost
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ResourceID,
		&i.PostTime,
		&i.PostTitle,
		&i.PostDescription,
	)
	return i, err
}

const getPostById = `-- name: GetPostById :one
SELECT id, user_id, resource_id, post_time, post_title, post_description FROM user_post
WHERE id = $1
`

func (q *Queries) GetPostById(ctx context.Context, id string) (UserPost, error) {
	row := q.queryRow(ctx, q.getPostByIdStmt, getPostById, id)
	var i UserPost
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ResourceID,
		&i.PostTime,
		&i.PostTitle,
		&i.PostDescription,
	)
	return i, err
}

const getPostDifficultyVotes = `-- name: GetPostDifficultyVotes :many
SELECT user_id, post_id, difficulty, comment from votes
WHERE post_id = $1
`

func (q *Queries) GetPostDifficultyVotes(ctx context.Context, postID string) ([]Vote, error) {
	rows, err := q.query(ctx, q.getPostDifficultyVotesStmt, getPostDifficultyVotes, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Vote{}
	for rows.Next() {
		var i Vote
		if err := rows.Scan(
			&i.UserID,
			&i.PostID,
			&i.Difficulty,
			&i.Comment,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPostLikes = `-- name: GetPostLikes :many
SELECT user_id, post_id from likes
WHERE post_id = $1
`

func (q *Queries) GetPostLikes(ctx context.Context, postID string) ([]Like, error) {
	rows, err := q.query(ctx, q.getPostLikesStmt, getPostLikes, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Like{}
	for rows.Next() {
		var i Like
		if err := rows.Scan(&i.UserID, &i.PostID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPostTags = `-- name: GetPostTags :many
SELECT post_id, tag FROM tags
WHERE post_id = $1
`

func (q *Queries) GetPostTags(ctx context.Context, postID string) ([]Tag, error) {
	rows, err := q.query(ctx, q.getPostTagsStmt, getPostTags, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Tag{}
	for rows.Next() {
		var i Tag
		if err := rows.Scan(&i.PostID, &i.Tag); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPostsByCategory = `-- name: GetPostsByCategory :many
SELECT p.id, p.user_id, p.resource_id, p.post_time, p.post_title, p.post_description FROM user_post AS p
JOIN resources AS r
ON p.resource_id = r.id
WHERE r.category = $1
ORDER BY p.resource_id
LIMIT $2 OFFSET $3
`

type GetPostsByCategoryParams struct {
	Category string `json:"category"`
	Limit    int32  `json:"limit"`
	Offset   int32  `json:"offset"`
}

func (q *Queries) GetPostsByCategory(ctx context.Context, arg GetPostsByCategoryParams) ([]UserPost, error) {
	rows, err := q.query(ctx, q.getPostsByCategoryStmt, getPostsByCategory, arg.Category, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []UserPost{}
	for rows.Next() {
		var i UserPost
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.ResourceID,
			&i.PostTime,
			&i.PostTitle,
			&i.PostDescription,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPostsByDifficulty = `-- name: GetPostsByDifficulty :many
SELECT p.id, p.user_id, p.resource_id, p.post_time, p.post_title, p.post_description FROM user_post AS p
JOIN resources AS r
ON p.resource_id = r.id
WHERE r.difficulty = $1
ORDER BY p.resource_id
LIMIT $2 OFFSET $3
`

type GetPostsByDifficultyParams struct {
	Difficulty sql.NullString `json:"difficulty"`
	Limit      int32          `json:"limit"`
	Offset     int32          `json:"offset"`
}

func (q *Queries) GetPostsByDifficulty(ctx context.Context, arg GetPostsByDifficultyParams) ([]UserPost, error) {
	rows, err := q.query(ctx, q.getPostsByDifficultyStmt, getPostsByDifficulty, arg.Difficulty, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []UserPost{}
	for rows.Next() {
		var i UserPost
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.ResourceID,
			&i.PostTime,
			&i.PostTitle,
			&i.PostDescription,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPostsByLanguage = `-- name: GetPostsByLanguage :many
SELECT p.id, p.user_id, p.resource_id, p.post_time, p.post_title, p.post_description FROM user_post AS p
JOIN resources AS r
ON p.resource_id = r.id
WHERE r.language = $1
ORDER BY p.resource_id
LIMIT $2 OFFSET $3
`

type GetPostsByLanguageParams struct {
	Language string `json:"language"`
	Limit    int32  `json:"limit"`
	Offset   int32  `json:"offset"`
}

func (q *Queries) GetPostsByLanguage(ctx context.Context, arg GetPostsByLanguageParams) ([]UserPost, error) {
	rows, err := q.query(ctx, q.getPostsByLanguageStmt, getPostsByLanguage, arg.Language, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []UserPost{}
	for rows.Next() {
		var i UserPost
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.ResourceID,
			&i.PostTime,
			&i.PostTitle,
			&i.PostDescription,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPostsByMediaType = `-- name: GetPostsByMediaType :many
SELECT p.id, p.user_id, p.resource_id, p.post_time, p.post_title, p.post_description FROM user_post AS p
JOIN resources AS r
ON p.resource_id = r.id
WHERE r.media_type = $1
ORDER BY p.resource_id
LIMIT $2 OFFSET $3
`

type GetPostsByMediaTypeParams struct {
	MediaType sql.NullString `json:"media_type"`
	Limit     int32          `json:"limit"`
	Offset    int32          `json:"offset"`
}

func (q *Queries) GetPostsByMediaType(ctx context.Context, arg GetPostsByMediaTypeParams) ([]UserPost, error) {
	rows, err := q.query(ctx, q.getPostsByMediaTypeStmt, getPostsByMediaType, arg.MediaType, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []UserPost{}
	for rows.Next() {
		var i UserPost
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.ResourceID,
			&i.PostTime,
			&i.PostTitle,
			&i.PostDescription,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPostsByUser = `-- name: GetPostsByUser :many
SELECT id, user_id, resource_id, post_time, post_title, post_description FROM user_post
WHERE user_id = $1
ORDER BY user_id
LIMIT $2 OFFSET $3
`

type GetPostsByUserParams struct {
	UserID string `json:"user_id"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

func (q *Queries) GetPostsByUser(ctx context.Context, arg GetPostsByUserParams) ([]UserPost, error) {
	rows, err := q.query(ctx, q.getPostsByUserStmt, getPostsByUser, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []UserPost{}
	for rows.Next() {
		var i UserPost
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.ResourceID,
			&i.PostTime,
			&i.PostTitle,
			&i.PostDescription,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getVote = `-- name: GetVote :one
SELECT user_id, post_id, difficulty, comment FROM votes
WHERE user_id = $1 AND post_id = $2
`

type GetVoteParams struct {
	UserID string `json:"user_id"`
	PostID string `json:"post_id"`
}

func (q *Queries) GetVote(ctx context.Context, arg GetVoteParams) (Vote, error) {
	row := q.queryRow(ctx, q.getVoteStmt, getVote, arg.UserID, arg.PostID)
	var i Vote
	err := row.Scan(
		&i.UserID,
		&i.PostID,
		&i.Difficulty,
		&i.Comment,
	)
	return i, err
}

const getVotes = `-- name: GetVotes :many
SELECT user_id, post_id, difficulty, comment FROM votes
WHERE post_id = $1
`

func (q *Queries) GetVotes(ctx context.Context, postID string) ([]Vote, error) {
	rows, err := q.query(ctx, q.getVotesStmt, getVotes, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Vote{}
	for rows.Next() {
		var i Vote
		if err := rows.Scan(
			&i.UserID,
			&i.PostID,
			&i.Difficulty,
			&i.Comment,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const removePost = `-- name: RemovePost :exec
DELETE FROM user_post 
WHERE id = $1
`

func (q *Queries) RemovePost(ctx context.Context, id string) error {
	_, err := q.exec(ctx, q.removePostStmt, removePost, id)
	return err
}

const removeUserPosts = `-- name: RemoveUserPosts :exec
DELETE FROM user_post
WHERE user_id = $1
`

func (q *Queries) RemoveUserPosts(ctx context.Context, userID string) error {
	_, err := q.exec(ctx, q.removeUserPostsStmt, removeUserPosts, userID)
	return err
}

const removeVote = `-- name: RemoveVote :exec
DELETE FROM votes
WHERE user_id = $1 AND post_id = $2
`

type RemoveVoteParams struct {
	UserID string `json:"user_id"`
	PostID string `json:"post_id"`
}

func (q *Queries) RemoveVote(ctx context.Context, arg RemoveVoteParams) error {
	_, err := q.exec(ctx, q.removeVoteStmt, removeVote, arg.UserID, arg.PostID)
	return err
}

const updatePost = `-- name: UpdatePost :one
UPDATE user_post 
SET post_title = $3, post_description = $4
WHERE user_id = $1 AND resource_id = $2
RETURNING id, user_id, resource_id, post_time, post_title, post_description
`

type UpdatePostParams struct {
	UserID          string         `json:"user_id"`
	ResourceID      int64          `json:"resource_id"`
	PostTitle       string         `json:"post_title"`
	PostDescription sql.NullString `json:"post_description"`
}

func (q *Queries) UpdatePost(ctx context.Context, arg UpdatePostParams) (UserPost, error) {
	row := q.queryRow(ctx, q.updatePostStmt, updatePost,
		arg.UserID,
		arg.ResourceID,
		arg.PostTitle,
		arg.PostDescription,
	)
	var i UserPost
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ResourceID,
		&i.PostTime,
		&i.PostTitle,
		&i.PostDescription,
	)
	return i, err
}

const updateVote = `-- name: UpdateVote :one
UPDATE votes
SET difficulty = $3, comment = $4
WHERE user_id = $1 AND post_id = $2
RETURNING user_id, post_id, difficulty, comment
`

type UpdateVoteParams struct {
	UserID     string         `json:"user_id"`
	PostID     string         `json:"post_id"`
	Difficulty string         `json:"difficulty"`
	Comment    sql.NullString `json:"comment"`
}

func (q *Queries) UpdateVote(ctx context.Context, arg UpdateVoteParams) (Vote, error) {
	row := q.queryRow(ctx, q.updateVoteStmt, updateVote,
		arg.UserID,
		arg.PostID,
		arg.Difficulty,
		arg.Comment,
	)
	var i Vote
	err := row.Scan(
		&i.UserID,
		&i.PostID,
		&i.Difficulty,
		&i.Comment,
	)
	return i, err
}

const votePost = `-- name: VotePost :one
INSERT INTO votes (
  user_id,
  post_id,
  difficulty,
  comment
) VALUES (
  $1, $2, $3, $4
) RETURNING user_id, post_id, difficulty, comment
`

type VotePostParams struct {
	UserID     string         `json:"user_id"`
	PostID     string         `json:"post_id"`
	Difficulty string         `json:"difficulty"`
	Comment    sql.NullString `json:"comment"`
}

func (q *Queries) VotePost(ctx context.Context, arg VotePostParams) (Vote, error) {
	row := q.queryRow(ctx, q.votePostStmt, votePost,
		arg.UserID,
		arg.PostID,
		arg.Difficulty,
		arg.Comment,
	)
	var i Vote
	err := row.Scan(
		&i.UserID,
		&i.PostID,
		&i.Difficulty,
		&i.Comment,
	)
	return i, err
}
