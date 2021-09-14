import {
  UPDATE_USER,
  UPDATE_TOKEN,
  START_LOADING_USER,
  END_LOADING_USER,
  INVALIDATE_USER,
  LOGIN_ERROR,
} from "../constants/actionTypes";

const USER_INITIAL_STATE = {
  user: JSON.parse(localStorage.getItem("user")) || null,
  accessToken: null,
  refreshToken: null,
  isLoading: false,
  error: false,
};

export default (state = USER_INITIAL_STATE, action) => {
  switch (action.type) {
    case UPDATE_USER:
      return action.payload;
    case UPDATE_TOKEN:
      return {
        ...state,
        accessToken: action.payload.accessToken,
        refreshToken: action.payload.refreshToken,
      };
    case INVALIDATE_USER:
      localStorage.setItem("user", null);
      return { user: null, isLoading: false, error: false };
    case START_LOADING_USER:
      return { ...state, isLoading: true };
    case END_LOADING_USER:
      return { ...state, isLoading: false };
    case LOGIN_ERROR:
      return { user: null, isLoading: false, error: true };
    default:
      return state;
  }
};
