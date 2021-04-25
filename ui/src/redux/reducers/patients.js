import { ADD_PATIENT } from '../actionTypes';

const initialState = {
  list: [],
};

export default (state = initialState, action) => {
  switch (action.type) {
    case ADD_PATIENT: {
      const patient = action.payload;
      return {
        ...state,
        list: [...state.list, { ...patient }],
      };
    }

    default:
      return state;
  }
};
