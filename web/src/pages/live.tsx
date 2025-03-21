import { LiveViewport } from "../components/live-viewport";
import { DashboardComponent } from "../components/dashboard";

export function LivePage() {
  return (
    <DashboardComponent sidebar={{ active: "live" }}>
      <LiveViewport />
    </DashboardComponent>
  );
}
