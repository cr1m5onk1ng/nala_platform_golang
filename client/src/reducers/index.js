import { combineReducers } from "redux";

import posts from "./posts";
import menus from "./menus";
import user from "./user.js";

const reducers = combineReducers({ posts, menus, user });

export default reducers;
