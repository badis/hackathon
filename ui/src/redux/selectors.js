export const getPatientsState = store => store.patients;

export const getPatientList = store =>
  getPatientsState(store) ? getPatientsState(store).patients : [];

