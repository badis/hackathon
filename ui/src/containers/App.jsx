import React from "react";
import { Layout } from "antd";

import Header from "../components/Header";
import "./App.css";

const App = (props) => {
  const { children } = props;
  const { Content } = Layout;

  return (
    <>
      <Layout>
        <Header />
        <Content>{children}</Content>
      </Layout>
    </>
  );
};

export default App;
