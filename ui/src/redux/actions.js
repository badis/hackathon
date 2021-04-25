import { ADD_PATIENT } from './actionTypes';

export const addPatient = (patient) => ({
  type: ADD_PATIENT,
  payload: {
    ...patient,
  },
});
