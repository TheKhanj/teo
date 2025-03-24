import { A } from "@solidjs/router";
import { JSX } from "solid-js";

type SideBarProps = {
  active: "live" | "recordings";
};

function Sidebar(props: SideBarProps) {
  const { active } = props;

  const getClasses = (current: SideBarProps["active"]) => {
    let c = "nav-item nav-link m-1";
    if (active === current) c += " active";
    return c;
  };

  return (
    <div
      class="col-2 vh-100 nav nav-pills p-2 d-flex flex-column bg-body-secondary"
      style="min-width: 200px"
    >
      <h3 class="text-center text-primary py-3">Teo 🐶</h3>
      <A class={getClasses("live")} href="/dashboard/live">
        Live
      </A>
      <A class={getClasses("recordings")} href="/dashboard/recordings">
        Recordings
      </A>
    </div>
  );
}

export function DashboardComponent(props: {
  sidebar: SideBarProps;
  children: JSX.Element;
}) {
  console.log(props);

  return (
    <div class="d-flex d-print-table">
      {Sidebar(props.sidebar)}
      {props.children}
    </div>
  );
}
