import React from "react";
import { Route } from "react-router-dom";

import App from "../containers/App";
import Login from "../containers/Login";
import Logout from "../containers/Logout";
import Register from "../containers/Register";
import Account from "../containers/Account";
import Landing from "../containers/Landing";

const Routes = () => (
  <App>
    <Route exact path="/" component={Landing} />
    <Route exact path="/login" component={Login} />
    <Route exact path="/logout" component={Logout} />
    <Route exact path="/register" component={Register} />
    <Route exact path="/account" component={Account} />
  </App>
);

export default Routes;
