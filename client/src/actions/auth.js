import {
  UPDATE_USER,
  INVALIDATE_USER,
  START_LOADING_USER,
  END_LOADING_USER,
  LOGIN_ERROR,
  LOGOUT_ERROR,
} from "../constants/actionTypes";
import * as api from "../api/index.js";

export const updateUser = (newUser) => async (dispatch) => {
  dispatch({ type: UPDATE_USER, payload: newUser });
};

export const login = (userData) => async (dispatch) => {
  dispatch({ type: START_LOADING_USER });
  try {
    const { message, error, data } = await api.signIn(userData);
    if (error) {
      dipatch({ type: LOGIN_ERROR, payload: message });
      return;
    }
    dispatch({ type: UPDATE_USER, payload: data });
    dispatch({ type: END_LOADING_USER });
  } catch (err) {
    console.log(err);
  }
};

export const logout = (data) => async (dispatch) => {
  try {
    await api.logout(data);
    dispatch({ type: INVALIDATE_USER });
  } catch (err) {
    dispatch({ type: LOGOUT_ERROR, payload: err });
  }
};
