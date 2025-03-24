import { Title } from "@solidjs/meta";
import { useNavigate } from "@solidjs/router";
import { createSignal } from "solid-js";

import { useApi } from "../api";
import { Route } from "../routes";

type Status =
  | {
      ok: null;
    }
  | {
      ok: boolean;
      message: string;
    };

function Alert(props: { status: Status }) {
  return (
    <>
      {props.status.ok === null ? (
        <div class="mb-3"></div>
      ) : (
        <div class="px-3">
          <div class="border-bottom m-3"></div>
          <div
            class="alert"
            classList={{
              "alert-success": props.status.ok,
              "alert-danger": !props.status.ok,
            }}
          >
            {props.status.message}
          </div>
        </div>
      )}
    </>
  );
}

export function LoginPage() {
  const [status, setStatus] = createSignal<Status>({ ok: null });
  const api = useApi();
  const nav = useNavigate();

  const handleSubmit = async (e: SubmitEvent) => {
    e.preventDefault();

    const form = document.getElementById("login-form") as HTMLFormElement;
    const formData = new FormData(form);
    const username = formData.get("username")?.toString() ?? "";
    const password = formData.get("password")?.toString() ?? "";

    const err = await api.authLogin(username, password);

    if (err === null) {
      setStatus({ ok: true, message: "Succeeded" });

      const route: Route = "/dashboard/live";
      nav(route);
    } else {
      setStatus({
        ok: false,
        message: err.message,
      });
    }
  };

  return (
    <>
      <Title>Login</Title>
      <form
        id="login-form"
        class="vh-100 vw-100 d-flex align-items-center justify-content-center p-2"
        onsubmit={handleSubmit}
      >
        <div class="card overflow-hidden col-11 col-sm-6 col-md-4 col-lg-3 d-flex flex-column">
          <h3 class="d-flex align-items-center justify-content-center text-primary py-3">
            Teo 🐶
          </h3>
          <div class="border-bottom mx-3 mb-3"></div>
          <div class="input-group mb-3 px-3">
            <input
              class="form-control"
              type="text"
              name="username"
              placeholder="Username"
            />
          </div>
          <div class="input-group mb-3 px-3">
            <input
              class="form-control"
              type="password"
              name="password"
              placeholder="Password"
            />
          </div>
          <div class="input-group px-3">
            <button type="submit" class="btn btn-primary">
              Login
            </button>
          </div>
          <Alert status={status()} />
        </div>
      </form>
    </>
  );
}
