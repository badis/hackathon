import React from "react";
import { useHistory } from "react-router-dom";

export default () => {
  const firstname = localStorage.getItem("firstname");
  const lastname = localStorage.getItem("lastname");
  const token = localStorage.getItem("token");

  if (!token) {
    const history = useHistory();
    history.push("/login");
  }

  return <div>Account: Welcome {firstname + " " + lastname}</div>;
};
