package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/cr1m5onk1ng/nala_platform_app/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		ID:             util.RandomString(20),
		Username:       util.RandomString(10),
		Email:          util.RandomString(10),
		Role:           sql.NullString{String: util.RandomString(5), Valid: true},
		NativeLanguage: util.RandomString(2),
	}
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.ID, user.ID)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Role, user.Role)
	require.Equal(t, arg.NativeLanguage, user.NativeLanguage)
	require.NotZero(t, user.ID)
	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.NativeLanguage, user2.NativeLanguage)
	require.Equal(t, user1.Role, user2.Role)
	require.WithinDuration(t, user1.RegistrationDate, user2.RegistrationDate, time.Second)
}

func TestDeleteUser(t *testing.T) {
	user1 := createRandomUser(t)
	err := testQueries.RemoveUser(context.Background(), user1.ID)
	require.NoError(t, err)

	account2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

func TestGetAllUsers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUser(t)
	}

	users, err := testQueries.GetAllUsers(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, users)

	for _, user := range users {
		require.NotEmpty(t, user)
	}
}

func TestUpdateUserLanguage(t *testing.T) {
	user := createRandomUser(t)
	params := UpdateUserLanguageParams{
		ID:             user.ID,
		NativeLanguage: "ja",
	}
	user2, err := testQueries.UpdateUserLanguage(context.Background(), params)
	require.NoError(t, err)
	require.NotEqual(t, user.NativeLanguage, user2.NativeLanguage)
	require.Equal(t, user2.NativeLanguage, "ja")
}
