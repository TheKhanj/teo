import { A } from "@solidjs/router";

export function IndexPage() {
  return (
    <>
      <ul class="nav nav-pills flex-column p-3">
        <li class="nav-item">
          <A class="nav-link" href="/login">
            login
          </A>
        </li>
        <li class="nav-item">
          <A class="nav-link" href="/dashboard">
            dashboard
          </A>
        </li>
      </ul>
    </>
  );
}
