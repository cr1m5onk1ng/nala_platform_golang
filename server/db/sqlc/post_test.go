package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/cr1m5onk1ng/nala_platform_app/util"
	"github.com/stretchr/testify/require"
)

func createRandomResource(t *testing.T) Resource {
	resourceParams := AddResourceParams{
		Url:          util.RandomString(20),
		Language:     util.RandomString(2),
		Difficulty:   sql.NullString{String: util.RandomString(5), Valid: true},
		Title:        util.RandomString(30),
		Description:  sql.NullString{String: util.RandomString(50), Valid: true},
		MediaType:    sql.NullString{String: util.RandomString(5), Valid: true},
		Category:     util.RandomString(5),
		ThumbnailUrl: sql.NullString{String: util.RandomString(5), Valid: true},
	}

	addedResource, err := testQueries.AddResource(context.Background(), resourceParams)
	require.NoError(t, err)
	require.NotEmpty(t, addedResource)
	return addedResource
}

func createRandomPost(t *testing.T) UserPost {
	resource := createRandomResource(t)
	user := createRandomUser(t)
	params := AddPostParams{
		ID:              util.RandomString(20),
		UserID:          user.ID,
		ResourceID:      resource.ID,
		PostTitle:       util.RandomString(30),
		PostDescription: sql.NullString{String: util.RandomString(50), Valid: true},
	}
	addedPost, err := testQueries.AddPost(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, addedPost)
	return addedPost
}

func TestCreatePost(t *testing.T) {
	post := createRandomPost(t)
	require.NotEmpty(t, post)
}

func TestGetPost(t *testing.T) {
	post := createRandomPost(t)
	retrievedPost, err := testQueries.GetPostById(context.Background(), post.ID)
	require.NoError(t, err)
	require.NotEmpty(t, retrievedPost)
	require.Equal(t, post.ID, retrievedPost.ID)
	require.Equal(t, post.UserID, retrievedPost.UserID)
	require.Equal(t, post.PostTitle, retrievedPost.PostTitle)
	require.WithinDuration(t, post.PostTime, retrievedPost.PostTime, time.Second)
}
