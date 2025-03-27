import { render } from "solid-js/web";
import { MetaProvider } from "@solidjs/meta";
import { Router, Route } from "@solidjs/router";

import { ROUTES } from "./routes";
import { ApiProvider } from "./providers/api";
import { NotFoundPage } from "./pages/error";
import { AppStateProvider } from "./providers/app-state";

function App() {
  return (
    <ApiProvider>
      <AppStateProvider>
        <MetaProvider>
          <Router>
            {Object.entries(ROUTES).map(([path, component]) => (
              <Route path={path} component={component} />
            ))}

            <Route path="*" component={NotFoundPage}></Route>
          </Router>
        </MetaProvider>
      </AppStateProvider>
    </ApiProvider>
  );
}

render(() => <App />, document.getElementById("app")!);
