import { render } from "solid-js/web";
import { MetaProvider } from "@solidjs/meta";
import { Router, Route } from "@solidjs/router";

import { ROUTES } from "./routes";
import { ApiProvider } from "./providers/api";
import { NotFoundPage } from "./pages/error";
import { CamerasProvider } from "./providers/cameras";

function App() {
  return (
    <ApiProvider>
      <CamerasProvider>
        <MetaProvider>
          <Router>
            {Object.entries(ROUTES).map(([path, component]) => (
              <Route path={path} component={component} />
            ))}

            <Route path="*" component={NotFoundPage}></Route>
          </Router>
        </MetaProvider>
      </CamerasProvider>
    </ApiProvider>
  );
}

render(() => <App />, document.getElementById("app")!);
