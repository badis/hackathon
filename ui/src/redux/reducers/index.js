import { combineReducers } from "redux";
import { reducer as formReducer } from "redux-form";

import patientReducer from "./patient";

const rootReducer = combineReducers({
  form: formReducer,
  patient: patientReducer,
});

export default rootReducer;
