import React from "react";
import ReactDOM from "react-dom";
import App from "./App";
import { Provider } from "react-redux";
import { createStore, applyMiddleware } from "redux";
import thunk from "redux-thunk";
import reducer from "./reducers";
import logger from "redux-logger";

const store = createStore(reducer, applyMiddleware(thunk, logger));

const rootElement = document.getElementById("root");
ReactDOM.render(
  <Provider store={store}>
    {" "}
    <App />{" "}
  </Provider>,
  rootElement
);
