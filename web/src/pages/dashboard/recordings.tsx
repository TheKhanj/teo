import { Title } from "@solidjs/meta";

import { DashboardComponent } from "../../components/dashboard";
import { RecordingsViewport } from "../../components/recordings-viewport";

export function RecordingsPage() {
  return (
    <>
      <Title>Dashboard - Recordings</Title>
      <DashboardComponent sidebar={{ active: "recordings" }}>
        <RecordingsViewport />
      </DashboardComponent>
    </>
  );
}
