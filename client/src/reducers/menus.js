import { OPEN_DRAWER, CLOSE_DRAWER } from "../constants/actionTypes";

export default (state = { isOpen: false }, action) => {
  switch (action.type) {
    case OPEN_DRAWER:
      return { isOpen: true };
    case CLOSE_DRAWER:
      return { isOpen: false };
    default:
      return state;
  }
};
