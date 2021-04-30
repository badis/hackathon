import { REGISTER_PATIENT } from "../actions/types";

const initialState = {};

export default (patient = initialState, action) => {
  switch (action.type) {
    case REGISTER_PATIENT: {
      return { ...action.payload };
    }

    default:
      return patient;
  }
};
