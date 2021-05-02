import React from "react";
import { Form, Input, Button, Checkbox, message } from "antd";
import { useHistory } from "react-router-dom";
import { connect } from "react-redux";
import { loginPatient } from "../redux/actions";

const layout = {
  labelCol: {
    span: 4,
  },
  wrapperCol: {
    span: 8,
  },
};

const tailLayout = {
  wrapperCol: {
    offset: 4,
    span: 8,
  },
};

const Login = (props) => {
  const history = useHistory();

  const onFinish = async (values) => {
    let response;
    try {
      response = await props.loginPatient(values);
      console.log(response);
      localStorage.setItem("firstname", response.authPatient.firstname);
      localStorage.setItem("lastname", response.authPatient.lastname);

      message.success("Authentication successful !");
      setTimeout(() => {
        history.push("/account");
      }, 2000);
    } catch (e) {
      console.log(e);
      message.error("Authentication Failed!");
    }
  };

  const onFinishFailed = (errorInfo) => {
    console.log("Failed:", errorInfo);
  };

  return (
    <Form
      {...layout}
      name="basic"
      initialValues={{
        remember: true,
      }}
      onFinish={onFinish}
      onFinishFailed={onFinishFailed}
    >
      <Form.Item
        label="Email"
        name="email"
        rules={[
          {
            required: true,
            message: "Please input your email!",
          },
        ]}
      >
        <Input />
      </Form.Item>

      <Form.Item
        label="Password"
        name="password"
        rules={[
          {
            required: true,
            message: "Please input your password!",
          },
        ]}
      >
        <Input.Password />
      </Form.Item>

      <Form.Item {...tailLayout} name="remember" valuePropName="checked">
        <Checkbox>Remember me</Checkbox>
      </Form.Item>

      <Form.Item {...tailLayout}>
        <Button type="primary" htmlType="submit">
          Submit
        </Button>
      </Form.Item>
    </Form>
  );
};

export default connect(null, { loginPatient })(Login);
