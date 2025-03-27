import {
  Accessor,
  createContext,
  createSignal,
  FlowProps,
  useContext,
} from "solid-js";

import { Oath } from "../util";
import { useApi } from "./api";

type Camera = {};
type Cameras = Record<string, Camera>;

type AppState = {
  cameras: Accessor<Cameras>;
  updateCameras: () => Oath<Error | null>;
};

const AppStateContext = createContext<AppState>();

export function AppStateProvider(props: FlowProps) {
  const [cams, updateCams] = createSignal<Cameras>({});
  const api = useApi();

  const v: AppState = {
    cameras: cams,
    updateCameras: async () => {
      const [cameras, err] = await api.cameras();
      if (err != null) return err;

      updateCams(
        cameras.reduce((prev, curr) => {
          prev[curr] = {};
          return prev;
        }, {} as Cameras),
      );
      return null;
    },
  };

  // TODO: find a better place to put this at
  (async () => {
    const err = await v.updateCameras();
    if (err != null) console.log(err);
  })();

  return (
    <AppStateContext.Provider value={v}>
      {props.children}
    </AppStateContext.Provider>
  );
}

export function useAppState() {
  return useContext(AppStateContext)!;
}
