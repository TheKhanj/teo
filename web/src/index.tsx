import { render } from "solid-js/web";
import { MetaProvider } from "@solidjs/meta";
import { Router, Route } from "@solidjs/router";

import { ROUTES } from "./routes";
import { NotFoundPage } from "./pages/error";

function App() {
  return (
    <MetaProvider>
      <Router>
        {Object.entries(ROUTES).map(([path, Component]) => (
          <Route path={path} component={Component} />
        ))}

        <Route path="*" component={NotFoundPage}></Route>
      </Router>
    </MetaProvider>
  );
}

render(() => <App />, document.getElementById("app")!);
