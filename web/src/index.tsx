import { render } from "solid-js/web";
import { Router, Route } from "@solidjs/router";

import { LivePage } from "./pages/live";
import { LoginPage } from "./pages/login";

function App() {
  return (
    <Router>
      <Route path="/login" component={LoginPage}></Route>
      <Route path="/dashboard" component={LivePage}></Route>
    </Router>
  );
}

render(() => <App />, document.getElementById("app")!);
