import React from "react";
import ReactDOM from "react-dom";
import App from "./App";
import { Provider } from "react-redux";
import { CssBaseline } from "@material-ui/core";
import { createStore, applyMiddleware, compose } from "redux";
import thunk from "redux-thunk";
import reducers from "./reducers";
import TimeAgo from "javascript-time-ago";
import en from "javascript-time-ago/locale/en";
import it from "javascript-time-ago/locale/it";

TimeAgo.addDefaultLocale(en);
TimeAgo.addLocale(it);

const store = createStore(reducers, {}, compose(applyMiddleware(thunk)));

ReactDOM.render(
  <Provider store={store}>
    <React.StrictMode>
      <CssBaseline />
      <App />
    </React.StrictMode>
  </Provider>,
  document.getElementById("root")
);
