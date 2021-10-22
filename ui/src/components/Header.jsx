import React from "react";
import { Layout } from "antd";
import { Link } from "react-router-dom";

import logo from "./../logo.png";

export default () => {
  const { Header } = Layout;

  const token = localStorage.getItem("token");

  return (
    <Header>
      <img
        style={{ maxWidth: "552px", width: "100%" }}
        src={logo}
        alt="Logo"
        max-width="552"
        width="100%"
      />
      <br />

      {!token ? (
        <>
          <Link to="/">Home.</Link> | <Link to="/login">Login</Link> |{" "}
          <Link to="/register">Register</Link>
        </>
      ) : (
        <>
          <Link to="/">Home</Link> | <Link to="/account">Account</Link> |{" "}
          <Link to="/logout">Logout</Link>
        </>
      )}
      <hr />
    </Header>
  );
};
