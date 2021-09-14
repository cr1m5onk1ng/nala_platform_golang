import {
  END_LOADING_POSTS_POSTS,
  START_LOADING_POSTS_POSTS,
  FETCH_POSTS_BY_LANGUAGE,
  FETCH_POST,
  FETCH_POSTS_BY_SEARCH,
  CREATE_POST,
  UPDATE_POST,
  DELETE_POST,
  LIKE_POST,
  COMMENT_POST,
  FETCH_POST_BY_CREATOR,
  FETCH_ERROR,
  POST_ERROR,
} from "../constants/actionTypes";
import * as api from "../api/index.js";

export const getPost = (id) => async (dispatch) => {
  try {
    dispatch({ type: START_LOADING_POSTS_POSTS });

    const { error, message, data } = await api.fetchPost(id);
    if (error) {
      dispatch({ type: FETCH_ERROR, payload: message });
      return;
    }

    dispatch({ type: FETCH_POST, payload: data });
    dispatch({ type: END_LOADING_POSTS_POSTS });
  } catch (error) {
    console.log(error);
  }
};

export const getPosts = (lang, page) => async (dispatch) => {
  try {
    dispatch({ type: START_LOADING_POSTS });
    const { error, message, data } = await api.fetchPostsByLanguage(lang, page);
    if (error) {
      dispatch({ type: FETCH_ERROR, payload: message });
      return;
    }
    dispatch({
      type: FETCH_POSTS_BY_LANGUAGE,
      payload: data,
    });
    dispatch({ type: END_LOADING_POSTS });
  } catch (error) {
    console.log(error);
  }
};

export const getPostsByCreator = (name) => async (dispatch) => {
  try {
    dispatch({ type: START_LOADING_POSTS });
    const { error, message, data } = await api.fetchPostsByCreator(name);
    if (error) {
      dispatch({ type: FETCH_ERROR, payload: message });
      return;
    }
    dispatch({ type: FETCH_POST_BY_CREATOR, payload: { data } });
    dispatch({ type: END_LOADING_POSTS });
  } catch (error) {
    console.log(error);
  }
};

export const getPostsBySearch = (searchQuery) => async (dispatch) => {
  try {
    dispatch({ type: START_LOADING_POSTS });
    const { error, message, data } = await api.fetchPostsBySearch(searchQuery);
    if (error) {
      dispatch({ type: FETCH_ERROR, payload: message });
      return;
    }
    dispatch({ type: FETCH_POSTS_BY_SEARCH, payload: data });
    dispatch({ type: END_LOADING_POSTS });
  } catch (error) {
    console.log(error);
  }
};

export const createPost = (post, history) => async (dispatch) => {
  try {
    dispatch({ type: START_LOADING_POSTS });
    const { error, message, data } = await api.createPost(post);
    if (error) {
      dispatch({ type: POST_ERROR, payload: message });
      return;
    }
    dispatch({ type: CREATE_POST, payload: data });
    history.push(`/posts/${data._id}`);
  } catch (error) {
    console.log(error);
  }
};

export const updatePost = (id, post) => async (dispatch) => {
  try {
    const { error, message, data } = await api.updatePost(id, post);
    if (error) {
      dispatch({ type: POST_ERROR, payload: message });
      return;
    }
    dispatch({ type: UPDATE_POST, payload: data });
  } catch (error) {
    console.log(error);
  }
};

export const likePost = (id) => async (dispatch) => {
  const user = JSON.parse(localStorage.getItem("profile"));

  try {
    const { error, message, data } = await api.likePost(id, user?.token);
    if (error) {
      dispatch({ type: POST_ERROR, payload: message });
      return;
    }
    dispatch({ type: LIKE_POST, payload: data });
  } catch (error) {
    console.log(error);
  }
};

export const commentPost = (value, id) => async (dispatch) => {
  try {
    const { error, message, data } = await api.commentPost(value, id);
    if (error) {
      dispatch({ type: POST_ERROR, payload: message });
      return;
    }
    dispatch({ type: COMMENT_POST, payload: data });
    return data.COMMENT_POSTs;
  } catch (error) {
    console.log(error);
  }
};

export const deletePost = (id) => async (dispatch) => {
  try {
    await api.deletePost(id);
    dispatch({ type: DELETE_POST, payload: id });
  } catch (error) {
    console.log(error);
  }
};
