import { OPEN_DRAWER, CLOSE_DRAWER } from "../constants/actionTypes";

export const openDrawer = () => {
  return {
    type: OPEN_DRAWER,
    value: { isOpen: true },
  };
};

export const closeDrawer = () => {
  return {
    type: CLOSE_DRAWER,
    value: { isOpen: false },
  };
};
