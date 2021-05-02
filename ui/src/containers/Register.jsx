import React, { useState } from "react";
import { Form, Input, Checkbox, Button, AutoComplete, message } from "antd";
import { useHistory } from "react-router-dom";

import { connect } from "react-redux";
import { registerPatient } from "../redux/actions";

const formItemLayout = {
  labelCol: {
    xs: {
      span: 24,
    },
    sm: {
      span: 8,
    },
    lg: {
      span: 4,
    },
  },
  wrapperCol: {
    xs: {
      span: 24,
    },
    sm: {
      span: 16,
    },
    lg: {
      span: 8,
    },
  },
};

const tailFormItemLayout = {
  wrapperCol: {
    xs: {
      offset: 0,
      span: 24,
    },
    sm: {
      offset: 8,
      span: 16,
    },
    lg: {
      offset: 4,
      span: 8,
    },
  },
};

const Register = (props) => {
  const [form] = Form.useForm();
  const history = useHistory();

  const onFinish = async (values) => {
    try {
      await props.registerPatient(values);
      message.success("Registration successful !");
      setTimeout(() => {
        history.push("/login");
      }, 2000);
    } catch (e) {
      console.log(e);
      message.error("Registration Failed!");
    }
  };

  const [autoCompleteResult, setAutoCompleteResult] = useState([]);

  const onDiseaseChange = (value) => {
    if (!value) {
      setAutoCompleteResult([]);
    }
  };

  const diseaseOptions = autoCompleteResult.map((disease) => ({
    label: disease,
    value: disease,
  }));

  return (
    <Form
      {...formItemLayout}
      form={form}
      name="register"
      onFinish={onFinish}
      initialValues={{
        residence: ["zhejiang", "hangzhou", "xihu"],
        prefix: "86",
      }}
      scrollToFirstError
    >
      <Form.Item
        name="firstname"
        label="First Name"
        rules={[
          {
            required: true,
            message: "Please input your first name!",
          },
        ]}
      >
        <Input />
      </Form.Item>

      <Form.Item
        name="lastname"
        label="Last Name"
        rules={[
          {
            required: true,
            message: "Please input your last name!",
          },
        ]}
      >
        <Input />
      </Form.Item>

      <Form.Item
        name="email"
        label="Email"
        rules={[
          {
            type: "email",
            message: "The input is not valid Email!",
          },
          {
            required: true,
            message: "Please input your Email!",
          },
        ]}
      >
        <Input />
      </Form.Item>

      <Form.Item
        name="password"
        label="Password"
        rules={[
          {
            required: true,
            message: "Please input your password!",
          },
        ]}
        hasFeedback
      >
        <Input.Password />
      </Form.Item>

      <Form.Item
        name="confirm"
        label="Confirm Password"
        dependencies={["password"]}
        hasFeedback
        rules={[
          {
            required: true,
            message: "Please confirm your password!",
          },
          ({ getFieldValue }) => ({
            validator(_, value) {
              if (!value || getFieldValue("password") === value) {
                return Promise.resolve();
              }

              return Promise.reject(
                new Error("The two passwords that you entered do not match!"),
              );
            },
          }),
        ]}
      >
        <Input.Password />
      </Form.Item>

      <Form.Item
        name="disease_name"
        label="Disease"
        rules={[
          {
            required: true,
            message: "Please enter the disease you have!",
          },
        ]}
      >
        <AutoComplete
          options={diseaseOptions}
          onChange={onDiseaseChange}
          placeholder="disease"
        >
          <Input />
        </AutoComplete>
      </Form.Item>

      <Form.Item
        name="agreement"
        valuePropName="checked"
        rules={[
          {
            validator: (_, value) =>
              value
                ? Promise.resolve()
                : Promise.reject(new Error("Should accept agreement")),
          },
        ]}
        {...tailFormItemLayout}
      >
        <Checkbox>I have read the agreement</Checkbox>
      </Form.Item>
      <Form.Item {...tailFormItemLayout}>
        <Button type="primary" htmlType="submit">
          Register
        </Button>
      </Form.Item>
    </Form>
  );
};

export default connect(null, { registerPatient })(Register);
