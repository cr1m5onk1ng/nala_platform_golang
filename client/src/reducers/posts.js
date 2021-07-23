import {
  START_LOADING,
  END_LOADING,
  FETCH_ALL_POSTS,
  FETCH_POSTS_BY_LANGUAGE,
  FETCH_POST,
  FETCH_POSTS_BY_SEARCH,
  CREATE_POST,
  UPDATE_POST,
  DELETE_POST,
  LIKE_POST,
  COMMENT_POST,
  FETCH_POST_BY_CREATOR,
} from "../constants/actionTypes";

export default (state = { isLoading: true, posts: [] }, action) => {
  switch (action.type) {
    case START_LOADING:
      return { ...state, isLoading: true };
    case END_LOADING:
      return { ...state, isLoading: false };
    case FETCH_ALL_POSTS:
      return {
        ...state,
        posts: action.payload.data,
      };
    case FETCH_POSTS_BY_LANGUAGE:
      return {
        ...state,
        posts: action.payload.data,
        currentPage: action.payload.currentPage,
        numberOfPages: action.payload.numberOfPages,
      };
    case FETCH_POSTS_BY_SEARCH:
    case FETCH_POST_BY_CREATOR:
      return { ...state, posts: action.payload.data };
    case FETCH_POST:
      return { ...state, post: action.payload.post };
    case LIKE_POST:
      return {
        ...state,
        posts: state.posts.map((post) =>
          post._id === action.payload._id ? action.payload : post
        ),
      };
    case COMMENT_POST:
      return {
        ...state,
        posts: state.posts.map((post) => {
          if (post._id === +action.payload._id) {
            return action.payload;
          }
          return post;
        }),
      };
    case CREATE_POST:
      return { ...state, posts: [...state.posts, action.payload] };
    case UPDATE_POST:
      return {
        ...state,
        posts: state.posts.map((post) =>
          post._id === action.payload._id ? action.payload : post
        ),
      };
    case DELETE_POST:
      return {
        ...state,
        posts: state.posts.filter((post) => post._id !== action.payload),
      };
    default:
      return state;
  }
};
