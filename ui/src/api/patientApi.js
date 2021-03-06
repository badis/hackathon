import http from "./http-common";

const register = (data) => http.post("/register", data);
const login = (data) => http.post("/login", data);

const patientAPI = {
  register,
  login,
};

export default patientAPI;
