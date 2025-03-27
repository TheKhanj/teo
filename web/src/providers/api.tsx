import { createContext, FlowProps, useContext } from "solid-js";

import { Oath, To, to } from "../util";
import { sessionGet, sessionSet } from "../session";

export class Api {
  private readonly headers = new Headers();

  public constructor() {
    const auth = sessionGet("authorization-header");
    if (auth) this.headers.set("Authorization", auth);
  }

  public async authLogin(
    username: string,
    password: string,
  ): Oath<Error | null> {
    const auth = `Basic ${btoa(`${username}:${password}`)}`;
    const [res, err] = await to(
      fetch("/api/auth/login", {
        headers: new Headers({ Authorization: auth }),
      }),
    );

    if (err !== null) return err;
    if (res.ok === false) {
      const [text, err] = await to(res.text());
      if (err != null) return err;
      return new Error(text);
    }

    this.headers.set("Authorization", auth);
    sessionSet("authorization-header", auth);

    return null;
  }

  public async cameras(): Oath<To<string[], Error>> {
    const [res, err] = await to(
      fetch("/api/cameras/", { headers: this.headers }),
    );
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
