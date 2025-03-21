import { render } from "solid-js/web";
import { Router, Route } from "@solidjs/router";

import { LivePage } from "./pages/live";
import { IndexPage } from "./pages";
import { LoginPage } from "./pages/login";

function App() {
  return (
    <Router>
      <Route path="/static/index.html" component={IndexPage}></Route>
      <Route path="/login" component={LoginPage}></Route>
      <Route path="/dashboard" component={LivePage}></Route>
    </Router>
  );
}

render(() => <App />, document.getElementById("app")!);
