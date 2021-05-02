import History from "../../utils/history";
import { REGISTER_PATIENT, AUTH_PATIENT, UNAUTH_PATIENT } from "./types";

import patientApi from "../../api/patientApi";

export const loginPatient = ({ email, password }) => async (dispatch) => {
  try {
    const res = await patientApi.login({ email, password });

    dispatch({
      type: AUTH_PATIENT,
      payload: res.data,
    });

    localStorage.setItem("token", res.data.token);
    History.push("/account");

    return Promise.resolve(res.data);
  } catch (err) {
    return Promise.reject(err);
  }
};

export const registerPatient = (patient) => async (dispatch) => {
  try {
    const res = await patientApi.register(patient);

    dispatch({
      type: REGISTER_PATIENT,
      payload: res.data,
    });

    localStorage.setItem("token", res.data.token);
    History.push("/account");

    return Promise.resolve(res.data);
  } catch (err) {
    return Promise.reject(err);
  }
};

export const logoutPatient = () => {
  localStorage.removeItem("token");
  return { type: UNAUTH_PATIENT };
};

/*
export const fetchAccount = () => (dispatch) => {

};
*/
