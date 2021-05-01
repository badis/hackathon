import React from "react";
import { Layout } from "antd";
import { Link } from "react-router-dom";

import logo from "./../logo.png";

export default () => {
  const { Header } = Layout;

  return (
    <Header>
      <Link to="/">
        <img
          style={{ maxWidth: "552px", width: "100%" }}
          src={logo}
          alt="Logo"
          max-width="552"
          width="100%"
        />
      </Link>
      <br />
      <Link to="/login">Login</Link> | <Link to="/register">Register</Link>
      <hr />
    </Header>
  );
};
