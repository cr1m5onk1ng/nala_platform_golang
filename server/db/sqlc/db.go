// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.addCommentStmt, err = db.PrepareContext(ctx, addComment); err != nil {
		return nil, fmt.Errorf("error preparing query AddComment: %w", err)
	}
	if q.addPostStmt, err = db.PrepareContext(ctx, addPost); err != nil {
		return nil, fmt.Errorf("error preparing query AddPost: %w", err)
	}
	if q.addPostDiscussionStmt, err = db.PrepareContext(ctx, addPostDiscussion); err != nil {
		return nil, fmt.Errorf("error preparing query AddPostDiscussion: %w", err)
	}
	if q.addResourceStmt, err = db.PrepareContext(ctx, addResource); err != nil {
		return nil, fmt.Errorf("error preparing query AddResource: %w", err)
	}
	if q.addResourceToStudyListStmt, err = db.PrepareContext(ctx, addResourceToStudyList); err != nil {
		return nil, fmt.Errorf("error preparing query AddResourceToStudyList: %w", err)
	}
	if q.addStudyListStmt, err = db.PrepareContext(ctx, addStudyList); err != nil {
		return nil, fmt.Errorf("error preparing query AddStudyList: %w", err)
	}
	if q.addUserTargetLanguageStmt, err = db.PrepareContext(ctx, addUserTargetLanguage); err != nil {
		return nil, fmt.Errorf("error preparing query AddUserTargetLanguage: %w", err)
	}
	if q.createUserStmt, err = db.PrepareContext(ctx, createUser); err != nil {
		return nil, fmt.Errorf("error preparing query CreateUser: %w", err)
	}
	if q.getAllDiscussionCommentsStmt, err = db.PrepareContext(ctx, getAllDiscussionComments); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllDiscussionComments: %w", err)
	}
	if q.getAllUserCommentsStmt, err = db.PrepareContext(ctx, getAllUserComments); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllUserComments: %w", err)
	}
	if q.getAllUserPostCommentsStmt, err = db.PrepareContext(ctx, getAllUserPostComments); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllUserPostComments: %w", err)
	}
	if q.getAllUsersStmt, err = db.PrepareContext(ctx, getAllUsers); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllUsers: %w", err)
	}
	if q.getCommentDirectResponsesStmt, err = db.PrepareContext(ctx, getCommentDirectResponses); err != nil {
		return nil, fmt.Errorf("error preparing query GetCommentDirectResponses: %w", err)
	}
	if q.getCommentLikesStmt, err = db.PrepareContext(ctx, getCommentLikes); err != nil {
		return nil, fmt.Errorf("error preparing query GetCommentLikes: %w", err)
	}
	if q.getCommentLikesCountStmt, err = db.PrepareContext(ctx, getCommentLikesCount); err != nil {
		return nil, fmt.Errorf("error preparing query GetCommentLikesCount: %w", err)
	}
	if q.getDiscussionCommentByIdStmt, err = db.PrepareContext(ctx, getDiscussionCommentById); err != nil {
		return nil, fmt.Errorf("error preparing query GetDiscussionCommentById: %w", err)
	}
	if q.getPostByIdStmt, err = db.PrepareContext(ctx, getPostById); err != nil {
		return nil, fmt.Errorf("error preparing query GetPostById: %w", err)
	}
	if q.getPostDifficultyVotesStmt, err = db.PrepareContext(ctx, getPostDifficultyVotes); err != nil {
		return nil, fmt.Errorf("error preparing query GetPostDifficultyVotes: %w", err)
	}
	if q.getPostDiscussionByIdStmt, err = db.PrepareContext(ctx, getPostDiscussionById); err != nil {
		return nil, fmt.Errorf("error preparing query GetPostDiscussionById: %w", err)
	}
	if q.getPostDiscussionsStmt, err = db.PrepareContext(ctx, getPostDiscussions); err != nil {
		return nil, fmt.Errorf("error preparing query GetPostDiscussions: %w", err)
	}
	if q.getPostDiscussionsByUserStmt, err = db.PrepareContext(ctx, getPostDiscussionsByUser); err != nil {
		return nil, fmt.Errorf("error preparing query GetPostDiscussionsByUser: %w", err)
	}
	if q.getPostLikesStmt, err = db.PrepareContext(ctx, getPostLikes); err != nil {
		return nil, fmt.Errorf("error preparing query GetPostLikes: %w", err)
	}
	if q.getPostTagsStmt, err = db.PrepareContext(ctx, getPostTags); err != nil {
		return nil, fmt.Errorf("error preparing query GetPostTags: %w", err)
	}
	if q.getPostsByCategoryStmt, err = db.PrepareContext(ctx, getPostsByCategory); err != nil {
		return nil, fmt.Errorf("error preparing query GetPostsByCategory: %w", err)
	}
	if q.getPostsByDifficultyStmt, err = db.PrepareContext(ctx, getPostsByDifficulty); err != nil {
		return nil, fmt.Errorf("error preparing query GetPostsByDifficulty: %w", err)
	}
	if q.getPostsByLanguageStmt, err = db.PrepareContext(ctx, getPostsByLanguage); err != nil {
		return nil, fmt.Errorf("error preparing query GetPostsByLanguage: %w", err)
	}
	if q.getPostsByMediaTypeStmt, err = db.PrepareContext(ctx, getPostsByMediaType); err != nil {
		return nil, fmt.Errorf("error preparing query GetPostsByMediaType: %w", err)
	}
	if q.getPostsByUserStmt, err = db.PrepareContext(ctx, getPostsByUser); err != nil {
		return nil, fmt.Errorf("error preparing query GetPostsByUser: %w", err)
	}
	if q.getResourceByIdStmt, err = db.PrepareContext(ctx, getResourceById); err != nil {
		return nil, fmt.Errorf("error preparing query GetResourceById: %w", err)
	}
	if q.getResourceByUrlStmt, err = db.PrepareContext(ctx, getResourceByUrl); err != nil {
		return nil, fmt.Errorf("error preparing query GetResourceByUrl: %w", err)
	}
	if q.getResourcePostStmt, err = db.PrepareContext(ctx, getResourcePost); err != nil {
		return nil, fmt.Errorf("error preparing query GetResourcePost: %w", err)
	}
	if q.getResourcesByLanguageStmt, err = db.PrepareContext(ctx, getResourcesByLanguage); err != nil {
		return nil, fmt.Errorf("error preparing query GetResourcesByLanguage: %w", err)
	}
	if q.getResourcesPostsByUserStmt, err = db.PrepareContext(ctx, getResourcesPostsByUser); err != nil {
		return nil, fmt.Errorf("error preparing query GetResourcesPostsByUser: %w", err)
	}
	if q.getStudyListStmt, err = db.PrepareContext(ctx, getStudyList); err != nil {
		return nil, fmt.Errorf("error preparing query GetStudyList: %w", err)
	}
	if q.getStudyListResourcesStmt, err = db.PrepareContext(ctx, getStudyListResources); err != nil {
		return nil, fmt.Errorf("error preparing query GetStudyListResources: %w", err)
	}
	if q.getUserStmt, err = db.PrepareContext(ctx, getUser); err != nil {
		return nil, fmt.Errorf("error preparing query GetUser: %w", err)
	}
	if q.getUserByEmailStmt, err = db.PrepareContext(ctx, getUserByEmail); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserByEmail: %w", err)
	}
	if q.getUserFollowersStmt, err = db.PrepareContext(ctx, getUserFollowers); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserFollowers: %w", err)
	}
	if q.getUserSavedResourcesStmt, err = db.PrepareContext(ctx, getUserSavedResources); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserSavedResources: %w", err)
	}
	if q.getUserStudyListsStmt, err = db.PrepareContext(ctx, getUserStudyLists); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserStudyLists: %w", err)
	}
	if q.getUserTargetLanguagesStmt, err = db.PrepareContext(ctx, getUserTargetLanguages); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserTargetLanguages: %w", err)
	}
	if q.getVoteStmt, err = db.PrepareContext(ctx, getVote); err != nil {
		return nil, fmt.Errorf("error preparing query GetVote: %w", err)
	}
	if q.getVotesStmt, err = db.PrepareContext(ctx, getVotes); err != nil {
		return nil, fmt.Errorf("error preparing query GetVotes: %w", err)
	}
	if q.likeCommentStmt, err = db.PrepareContext(ctx, likeComment); err != nil {
		return nil, fmt.Errorf("error preparing query LikeComment: %w", err)
	}
	if q.removeAllUserStudyListsStmt, err = db.PrepareContext(ctx, removeAllUserStudyLists); err != nil {
		return nil, fmt.Errorf("error preparing query RemoveAllUserStudyLists: %w", err)
	}
	if q.removeCommentStmt, err = db.PrepareContext(ctx, removeComment); err != nil {
		return nil, fmt.Errorf("error preparing query RemoveComment: %w", err)
	}
	if q.removeDiscussionCommentsStmt, err = db.PrepareContext(ctx, removeDiscussionComments); err != nil {
		return nil, fmt.Errorf("error preparing query RemoveDiscussionComments: %w", err)
	}
	if q.removePostStmt, err = db.PrepareContext(ctx, removePost); err != nil {
		return nil, fmt.Errorf("error preparing query RemovePost: %w", err)
	}
	if q.removePostDiscussionStmt, err = db.PrepareContext(ctx, removePostDiscussion); err != nil {
		return nil, fmt.Errorf("error preparing query RemovePostDiscussion: %w", err)
	}
	if q.removePostDiscussionsByCreatorStmt, err = db.PrepareContext(ctx, removePostDiscussionsByCreator); err != nil {
		return nil, fmt.Errorf("error preparing query RemovePostDiscussionsByCreator: %w", err)
	}
	if q.removeResourceFromStudyListStmt, err = db.PrepareContext(ctx, removeResourceFromStudyList); err != nil {
		return nil, fmt.Errorf("error preparing query RemoveResourceFromStudyList: %w", err)
	}
	if q.removeStudyListStmt, err = db.PrepareContext(ctx, removeStudyList); err != nil {
		return nil, fmt.Errorf("error preparing query RemoveStudyList: %w", err)
	}
	if q.removeUserStmt, err = db.PrepareContext(ctx, removeUser); err != nil {
		return nil, fmt.Errorf("error preparing query RemoveUser: %w", err)
	}
	if q.removeUserCommentsStmt, err = db.PrepareContext(ctx, removeUserComments); err != nil {
		return nil, fmt.Errorf("error preparing query RemoveUserComments: %w", err)
	}
	if q.removeUserPostsStmt, err = db.PrepareContext(ctx, removeUserPosts); err != nil {
		return nil, fmt.Errorf("error preparing query RemoveUserPosts: %w", err)
	}
	if q.removeVoteStmt, err = db.PrepareContext(ctx, removeVote); err != nil {
		return nil, fmt.Errorf("error preparing query RemoveVote: %w", err)
	}
	if q.unlikeCommentStmt, err = db.PrepareContext(ctx, unlikeComment); err != nil {
		return nil, fmt.Errorf("error preparing query UnlikeComment: %w", err)
	}
	if q.updatePostStmt, err = db.PrepareContext(ctx, updatePost); err != nil {
		return nil, fmt.Errorf("error preparing query UpdatePost: %w", err)
	}
	if q.updatePostDiscussionStmt, err = db.PrepareContext(ctx, updatePostDiscussion); err != nil {
		return nil, fmt.Errorf("error preparing query UpdatePostDiscussion: %w", err)
	}
	if q.updateResourceStmt, err = db.PrepareContext(ctx, updateResource); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateResource: %w", err)
	}
	if q.updateStudyListStmt, err = db.PrepareContext(ctx, updateStudyList); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateStudyList: %w", err)
	}
	if q.updateUserLanguageStmt, err = db.PrepareContext(ctx, updateUserLanguage); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateUserLanguage: %w", err)
	}
	if q.updateUserRoleStmt, err = db.PrepareContext(ctx, updateUserRole); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateUserRole: %w", err)
	}
	if q.updateVoteStmt, err = db.PrepareContext(ctx, updateVote); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateVote: %w", err)
	}
	if q.votePostStmt, err = db.PrepareContext(ctx, votePost); err != nil {
		return nil, fmt.Errorf("error preparing query VotePost: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.addCommentStmt != nil {
		if cerr := q.addCommentStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addCommentStmt: %w", cerr)
		}
	}
	if q.addPostStmt != nil {
		if cerr := q.addPostStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addPostStmt: %w", cerr)
		}
	}
	if q.addPostDiscussionStmt != nil {
		if cerr := q.addPostDiscussionStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addPostDiscussionStmt: %w", cerr)
		}
	}
	if q.addResourceStmt != nil {
		if cerr := q.addResourceStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addResourceStmt: %w", cerr)
		}
	}
	if q.addResourceToStudyListStmt != nil {
		if cerr := q.addResourceToStudyListStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addResourceToStudyListStmt: %w", cerr)
		}
	}
	if q.addStudyListStmt != nil {
		if cerr := q.addStudyListStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addStudyListStmt: %w", cerr)
		}
	}
	if q.addUserTargetLanguageStmt != nil {
		if cerr := q.addUserTargetLanguageStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addUserTargetLanguageStmt: %w", cerr)
		}
	}
	if q.createUserStmt != nil {
		if cerr := q.createUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createUserStmt: %w", cerr)
		}
	}
	if q.getAllDiscussionCommentsStmt != nil {
		if cerr := q.getAllDiscussionCommentsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllDiscussionCommentsStmt: %w", cerr)
		}
	}
	if q.getAllUserCommentsStmt != nil {
		if cerr := q.getAllUserCommentsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllUserCommentsStmt: %w", cerr)
		}
	}
	if q.getAllUserPostCommentsStmt != nil {
		if cerr := q.getAllUserPostCommentsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllUserPostCommentsStmt: %w", cerr)
		}
	}
	if q.getAllUsersStmt != nil {
		if cerr := q.getAllUsersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllUsersStmt: %w", cerr)
		}
	}
	if q.getCommentDirectResponsesStmt != nil {
		if cerr := q.getCommentDirectResponsesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getCommentDirectResponsesStmt: %w", cerr)
		}
	}
	if q.getCommentLikesStmt != nil {
		if cerr := q.getCommentLikesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getCommentLikesStmt: %w", cerr)
		}
	}
	if q.getCommentLikesCountStmt != nil {
		if cerr := q.getCommentLikesCountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getCommentLikesCountStmt: %w", cerr)
		}
	}
	if q.getDiscussionCommentByIdStmt != nil {
		if cerr := q.getDiscussionCommentByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getDiscussionCommentByIdStmt: %w", cerr)
		}
	}
	if q.getPostByIdStmt != nil {
		if cerr := q.getPostByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPostByIdStmt: %w", cerr)
		}
	}
	if q.getPostDifficultyVotesStmt != nil {
		if cerr := q.getPostDifficultyVotesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPostDifficultyVotesStmt: %w", cerr)
		}
	}
	if q.getPostDiscussionByIdStmt != nil {
		if cerr := q.getPostDiscussionByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPostDiscussionByIdStmt: %w", cerr)
		}
	}
	if q.getPostDiscussionsStmt != nil {
		if cerr := q.getPostDiscussionsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPostDiscussionsStmt: %w", cerr)
		}
	}
	if q.getPostDiscussionsByUserStmt != nil {
		if cerr := q.getPostDiscussionsByUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPostDiscussionsByUserStmt: %w", cerr)
		}
	}
	if q.getPostLikesStmt != nil {
		if cerr := q.getPostLikesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPostLikesStmt: %w", cerr)
		}
	}
	if q.getPostTagsStmt != nil {
		if cerr := q.getPostTagsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPostTagsStmt: %w", cerr)
		}
	}
	if q.getPostsByCategoryStmt != nil {
		if cerr := q.getPostsByCategoryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPostsByCategoryStmt: %w", cerr)
		}
	}
	if q.getPostsByDifficultyStmt != nil {
		if cerr := q.getPostsByDifficultyStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPostsByDifficultyStmt: %w", cerr)
		}
	}
	if q.getPostsByLanguageStmt != nil {
		if cerr := q.getPostsByLanguageStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPostsByLanguageStmt: %w", cerr)
		}
	}
	if q.getPostsByMediaTypeStmt != nil {
		if cerr := q.getPostsByMediaTypeStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPostsByMediaTypeStmt: %w", cerr)
		}
	}
	if q.getPostsByUserStmt != nil {
		if cerr := q.getPostsByUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPostsByUserStmt: %w", cerr)
		}
	}
	if q.getResourceByIdStmt != nil {
		if cerr := q.getResourceByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getResourceByIdStmt: %w", cerr)
		}
	}
	if q.getResourceByUrlStmt != nil {
		if cerr := q.getResourceByUrlStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getResourceByUrlStmt: %w", cerr)
		}
	}
	if q.getResourcePostStmt != nil {
		if cerr := q.getResourcePostStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getResourcePostStmt: %w", cerr)
		}
	}
	if q.getResourcesByLanguageStmt != nil {
		if cerr := q.getResourcesByLanguageStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getResourcesByLanguageStmt: %w", cerr)
		}
	}
	if q.getResourcesPostsByUserStmt != nil {
		if cerr := q.getResourcesPostsByUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getResourcesPostsByUserStmt: %w", cerr)
		}
	}
	if q.getStudyListStmt != nil {
		if cerr := q.getStudyListStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getStudyListStmt: %w", cerr)
		}
	}
	if q.getStudyListResourcesStmt != nil {
		if cerr := q.getStudyListResourcesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getStudyListResourcesStmt: %w", cerr)
		}
	}
	if q.getUserStmt != nil {
		if cerr := q.getUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserStmt: %w", cerr)
		}
	}
	if q.getUserByEmailStmt != nil {
		if cerr := q.getUserByEmailStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserByEmailStmt: %w", cerr)
		}
	}
	if q.getUserFollowersStmt != nil {
		if cerr := q.getUserFollowersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserFollowersStmt: %w", cerr)
		}
	}
	if q.getUserSavedResourcesStmt != nil {
		if cerr := q.getUserSavedResourcesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserSavedResourcesStmt: %w", cerr)
		}
	}
	if q.getUserStudyListsStmt != nil {
		if cerr := q.getUserStudyListsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserStudyListsStmt: %w", cerr)
		}
	}
	if q.getUserTargetLanguagesStmt != nil {
		if cerr := q.getUserTargetLanguagesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserTargetLanguagesStmt: %w", cerr)
		}
	}
	if q.getVoteStmt != nil {
		if cerr := q.getVoteStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getVoteStmt: %w", cerr)
		}
	}
	if q.getVotesStmt != nil {
		if cerr := q.getVotesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getVotesStmt: %w", cerr)
		}
	}
	if q.likeCommentStmt != nil {
		if cerr := q.likeCommentStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing likeCommentStmt: %w", cerr)
		}
	}
	if q.removeAllUserStudyListsStmt != nil {
		if cerr := q.removeAllUserStudyListsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing removeAllUserStudyListsStmt: %w", cerr)
		}
	}
	if q.removeCommentStmt != nil {
		if cerr := q.removeCommentStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing removeCommentStmt: %w", cerr)
		}
	}
	if q.removeDiscussionCommentsStmt != nil {
		if cerr := q.removeDiscussionCommentsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing removeDiscussionCommentsStmt: %w", cerr)
		}
	}
	if q.removePostStmt != nil {
		if cerr := q.removePostStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing removePostStmt: %w", cerr)
		}
	}
	if q.removePostDiscussionStmt != nil {
		if cerr := q.removePostDiscussionStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing removePostDiscussionStmt: %w", cerr)
		}
	}
	if q.removePostDiscussionsByCreatorStmt != nil {
		if cerr := q.removePostDiscussionsByCreatorStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing removePostDiscussionsByCreatorStmt: %w", cerr)
		}
	}
	if q.removeResourceFromStudyListStmt != nil {
		if cerr := q.removeResourceFromStudyListStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing removeResourceFromStudyListStmt: %w", cerr)
		}
	}
	if q.removeStudyListStmt != nil {
		if cerr := q.removeStudyListStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing removeStudyListStmt: %w", cerr)
		}
	}
	if q.removeUserStmt != nil {
		if cerr := q.removeUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing removeUserStmt: %w", cerr)
		}
	}
	if q.removeUserCommentsStmt != nil {
		if cerr := q.removeUserCommentsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing removeUserCommentsStmt: %w", cerr)
		}
	}
	if q.removeUserPostsStmt != nil {
		if cerr := q.removeUserPostsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing removeUserPostsStmt: %w", cerr)
		}
	}
	if q.removeVoteStmt != nil {
		if cerr := q.removeVoteStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing removeVoteStmt: %w", cerr)
		}
	}
	if q.unlikeCommentStmt != nil {
		if cerr := q.unlikeCommentStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing unlikeCommentStmt: %w", cerr)
		}
	}
	if q.updatePostStmt != nil {
		if cerr := q.updatePostStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updatePostStmt: %w", cerr)
		}
	}
	if q.updatePostDiscussionStmt != nil {
		if cerr := q.updatePostDiscussionStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updatePostDiscussionStmt: %w", cerr)
		}
	}
	if q.updateResourceStmt != nil {
		if cerr := q.updateResourceStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateResourceStmt: %w", cerr)
		}
	}
	if q.updateStudyListStmt != nil {
		if cerr := q.updateStudyListStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateStudyListStmt: %w", cerr)
		}
	}
	if q.updateUserLanguageStmt != nil {
		if cerr := q.updateUserLanguageStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateUserLanguageStmt: %w", cerr)
		}
	}
	if q.updateUserRoleStmt != nil {
		if cerr := q.updateUserRoleStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateUserRoleStmt: %w", cerr)
		}
	}
	if q.updateVoteStmt != nil {
		if cerr := q.updateVoteStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateVoteStmt: %w", cerr)
		}
	}
	if q.votePostStmt != nil {
		if cerr := q.votePostStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing votePostStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                                 DBTX
	tx                                 *sql.Tx
	addCommentStmt                     *sql.Stmt
	addPostStmt                        *sql.Stmt
	addPostDiscussionStmt              *sql.Stmt
	addResourceStmt                    *sql.Stmt
	addResourceToStudyListStmt         *sql.Stmt
	addStudyListStmt                   *sql.Stmt
	addUserTargetLanguageStmt          *sql.Stmt
	createUserStmt                     *sql.Stmt
	getAllDiscussionCommentsStmt       *sql.Stmt
	getAllUserCommentsStmt             *sql.Stmt
	getAllUserPostCommentsStmt         *sql.Stmt
	getAllUsersStmt                    *sql.Stmt
	getCommentDirectResponsesStmt      *sql.Stmt
	getCommentLikesStmt                *sql.Stmt
	getCommentLikesCountStmt           *sql.Stmt
	getDiscussionCommentByIdStmt       *sql.Stmt
	getPostByIdStmt                    *sql.Stmt
	getPostDifficultyVotesStmt         *sql.Stmt
	getPostDiscussionByIdStmt          *sql.Stmt
	getPostDiscussionsStmt             *sql.Stmt
	getPostDiscussionsByUserStmt       *sql.Stmt
	getPostLikesStmt                   *sql.Stmt
	getPostTagsStmt                    *sql.Stmt
	getPostsByCategoryStmt             *sql.Stmt
	getPostsByDifficultyStmt           *sql.Stmt
	getPostsByLanguageStmt             *sql.Stmt
	getPostsByMediaTypeStmt            *sql.Stmt
	getPostsByUserStmt                 *sql.Stmt
	getResourceByIdStmt                *sql.Stmt
	getResourceByUrlStmt               *sql.Stmt
	getResourcePostStmt                *sql.Stmt
	getResourcesByLanguageStmt         *sql.Stmt
	getResourcesPostsByUserStmt        *sql.Stmt
	getStudyListStmt                   *sql.Stmt
	getStudyListResourcesStmt          *sql.Stmt
	getUserStmt                        *sql.Stmt
	getUserByEmailStmt                 *sql.Stmt
	getUserFollowersStmt               *sql.Stmt
	getUserSavedResourcesStmt          *sql.Stmt
	getUserStudyListsStmt              *sql.Stmt
	getUserTargetLanguagesStmt         *sql.Stmt
	getVoteStmt                        *sql.Stmt
	getVotesStmt                       *sql.Stmt
	likeCommentStmt                    *sql.Stmt
	removeAllUserStudyListsStmt        *sql.Stmt
	removeCommentStmt                  *sql.Stmt
	removeDiscussionCommentsStmt       *sql.Stmt
	removePostStmt                     *sql.Stmt
	removePostDiscussionStmt           *sql.Stmt
	removePostDiscussionsByCreatorStmt *sql.Stmt
	removeResourceFromStudyListStmt    *sql.Stmt
	removeStudyListStmt                *sql.Stmt
	removeUserStmt                     *sql.Stmt
	removeUserCommentsStmt             *sql.Stmt
	removeUserPostsStmt                *sql.Stmt
	removeVoteStmt                     *sql.Stmt
	unlikeCommentStmt                  *sql.Stmt
	updatePostStmt                     *sql.Stmt
	updatePostDiscussionStmt           *sql.Stmt
	updateResourceStmt                 *sql.Stmt
	updateStudyListStmt                *sql.Stmt
	updateUserLanguageStmt             *sql.Stmt
	updateUserRoleStmt                 *sql.Stmt
	updateVoteStmt                     *sql.Stmt
	votePostStmt                       *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                                 tx,
		tx:                                 tx,
		addCommentStmt:                     q.addCommentStmt,
		addPostStmt:                        q.addPostStmt,
		addPostDiscussionStmt:              q.addPostDiscussionStmt,
		addResourceStmt:                    q.addResourceStmt,
		addResourceToStudyListStmt:         q.addResourceToStudyListStmt,
		addStudyListStmt:                   q.addStudyListStmt,
		addUserTargetLanguageStmt:          q.addUserTargetLanguageStmt,
		createUserStmt:                     q.createUserStmt,
		getAllDiscussionCommentsStmt:       q.getAllDiscussionCommentsStmt,
		getAllUserCommentsStmt:             q.getAllUserCommentsStmt,
		getAllUserPostCommentsStmt:         q.getAllUserPostCommentsStmt,
		getAllUsersStmt:                    q.getAllUsersStmt,
		getCommentDirectResponsesStmt:      q.getCommentDirectResponsesStmt,
		getCommentLikesStmt:                q.getCommentLikesStmt,
		getCommentLikesCountStmt:           q.getCommentLikesCountStmt,
		getDiscussionCommentByIdStmt:       q.getDiscussionCommentByIdStmt,
		getPostByIdStmt:                    q.getPostByIdStmt,
		getPostDifficultyVotesStmt:         q.getPostDifficultyVotesStmt,
		getPostDiscussionByIdStmt:          q.getPostDiscussionByIdStmt,
		getPostDiscussionsStmt:             q.getPostDiscussionsStmt,
		getPostDiscussionsByUserStmt:       q.getPostDiscussionsByUserStmt,
		getPostLikesStmt:                   q.getPostLikesStmt,
		getPostTagsStmt:                    q.getPostTagsStmt,
		getPostsByCategoryStmt:             q.getPostsByCategoryStmt,
		getPostsByDifficultyStmt:           q.getPostsByDifficultyStmt,
		getPostsByLanguageStmt:             q.getPostsByLanguageStmt,
		getPostsByMediaTypeStmt:            q.getPostsByMediaTypeStmt,
		getPostsByUserStmt:                 q.getPostsByUserStmt,
		getResourceByIdStmt:                q.getResourceByIdStmt,
		getResourceByUrlStmt:               q.getResourceByUrlStmt,
		getResourcePostStmt:                q.getResourcePostStmt,
		getResourcesByLanguageStmt:         q.getResourcesByLanguageStmt,
		getResourcesPostsByUserStmt:        q.getResourcesPostsByUserStmt,
		getStudyListStmt:                   q.getStudyListStmt,
		getStudyListResourcesStmt:          q.getStudyListResourcesStmt,
		getUserStmt:                        q.getUserStmt,
		getUserByEmailStmt:                 q.getUserByEmailStmt,
		getUserFollowersStmt:               q.getUserFollowersStmt,
		getUserSavedResourcesStmt:          q.getUserSavedResourcesStmt,
		getUserStudyListsStmt:              q.getUserStudyListsStmt,
		getUserTargetLanguagesStmt:         q.getUserTargetLanguagesStmt,
		getVoteStmt:                        q.getVoteStmt,
		getVotesStmt:                       q.getVotesStmt,
		likeCommentStmt:                    q.likeCommentStmt,
		removeAllUserStudyListsStmt:        q.removeAllUserStudyListsStmt,
		removeCommentStmt:                  q.removeCommentStmt,
		removeDiscussionCommentsStmt:       q.removeDiscussionCommentsStmt,
		removePostStmt:                     q.removePostStmt,
		removePostDiscussionStmt:           q.removePostDiscussionStmt,
		removePostDiscussionsByCreatorStmt: q.removePostDiscussionsByCreatorStmt,
		removeResourceFromStudyListStmt:    q.removeResourceFromStudyListStmt,
		removeStudyListStmt:                q.removeStudyListStmt,
		removeUserStmt:                     q.removeUserStmt,
		removeUserCommentsStmt:             q.removeUserCommentsStmt,
		removeUserPostsStmt:                q.removeUserPostsStmt,
		removeVoteStmt:                     q.removeVoteStmt,
		unlikeCommentStmt:                  q.unlikeCommentStmt,
		updatePostStmt:                     q.updatePostStmt,
		updatePostDiscussionStmt:           q.updatePostDiscussionStmt,
		updateResourceStmt:                 q.updateResourceStmt,
		updateStudyListStmt:                q.updateStudyListStmt,
		updateUserLanguageStmt:             q.updateUserLanguageStmt,
		updateUserRoleStmt:                 q.updateUserRoleStmt,
		updateVoteStmt:                     q.updateVoteStmt,
		votePostStmt:                       q.votePostStmt,
	}
}
