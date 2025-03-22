import { Title } from "@solidjs/meta";

import { LiveViewport } from "../../components/live-viewport";
import { DashboardComponent } from "../../components/dashboard";

export function LivePage() {
  return (
    <>
      <Title>Dashboard</Title>
      <DashboardComponent sidebar={{ active: "live" }}>
        <LiveViewport />
      </DashboardComponent>
    </>
  );
}
