import { render } from "solid-js/web";
import { MetaProvider } from "@solidjs/meta";
import { Router, Route } from "@solidjs/router";

import { ROUTES } from "./routes";
import { NotFoundPage } from "./pages/error";
import { ApiProvider } from "./api";

function App() {
  return (
    <ApiProvider>
      <MetaProvider>
        <Router>
          {Object.entries(ROUTES).map(([path, component]) => (
            <Route path={path} component={component} />
          ))}

          <Route path="*" component={NotFoundPage}></Route>
        </Router>
      </MetaProvider>
    </ApiProvider>
  );
}

render(() => <App />, document.getElementById("app")!);
