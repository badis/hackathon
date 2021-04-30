import { createStore, applyMiddleware } from "redux";
import reduxThunk from "redux-thunk";
import { composeWithDevTools } from "redux-devtools-extension";

import rootReducer from "./reducers";
import { AUTH_PATIENT } from "./actions/types";

// UNCOMMENT IT FOR PRODUCTION
// const createStoreWithMiddleware = applyMiddleware(reduxThunk)(createStore);
// const store = createStoreWithMiddleware(rootReducer);

/* COMMENT IT OUT FOR PRODUCTION */
const store = createStore(
  rootReducer,
  composeWithDevTools(applyMiddleware(reduxThunk)),
);

const token = localStorage.getItem("token");
// if we have a token, consider the user to be signed in
if (token) {
  // we need to update application state
  store.dispatch({ type: AUTH_PATIENT });
}

export default store;
