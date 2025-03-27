import { createContext, FlowProps, useContext } from "solid-js";

import { Oath, To, to } from "../util";

export class Api {
  public async cameras(): Oath<To<string[], Error>> {
    const [res, err] = await to(fetch("/api/cameras/"));
    if (err !== null) return [null, err];

    if (!res.ok) {
      const [text, err] = await to(res.text());
      if (err !== null) return [null, err];
      return [null, new Error(text)];
    }

    const body: string[] = await res.json();
    return [body, null];
  }
}

const ApiContext = createContext<Api>();

export function ApiProvider(props: FlowProps) {
  const api = new Api();

  return (
    <ApiContext.Provider value={api}>{props.children}</ApiContext.Provider>
  );
}

export function useApi() {
  return useContext(ApiContext)!;
}
