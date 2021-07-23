import { combineReducers } from "redux";

import posts from "./posts";
import menus from "./menus";
//import auth from './auth';

const reducers = combineReducers({ posts, menus });

export default reducers;
