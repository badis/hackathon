import React from "react";
import { useHistory } from "react-router-dom";

export default () => {
  localStorage.removeItem("token");
  const history = useHistory();
  history.push("/login");
  return <div>Logout ...</div>;
};
