import { JSX } from "solid-js";

type SideBarProps = {
  active: "live" | "recordings";
};

function Sidebar(props: SideBarProps) {
  // TODO: fix active element
  return (
    <div
      class="col-2 vh-100 nav nav-pills p-2 d-flex flex-column bg-body-secondary"
      style="min-width: 200px"
    >
      <h3 class="text-center text-primary py-3">Teo 🐶</h3>
      <a class="nav-item nav-link m-1 active" href="#">
        Live
      </a>
      <a class="nav-item nav-link m-1" href="#">
        Recordings
      </a>
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
