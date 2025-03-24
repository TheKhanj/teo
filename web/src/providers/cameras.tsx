import {
  Accessor,
  createContext,
  createSignal,
  FlowProps,
  useContext,
} from "solid-js";

import { Oath } from "../util";
import { useApi } from "./api";

type Cameras = {
  cameras: Accessor<string[]>;
  update: () => Oath<Error | null>;
};

const CamerasContext = createContext<Cameras>();

export function CamerasProvider(props: FlowProps) {
  const [cams, updateCams] = createSignal<string[]>([]);
  const api = useApi();

  const v: Cameras = {
    cameras: cams,
    update: async () => {
      const [cameras, err] = await api.cameras();
      if (err != null) return err;

      updateCams(cameras);
      return null;
    },
  };

  // TODO: find a better place to put this at
  (async () => {
    const err = await v.update();
    if (err != null) console.log(err);
  })();

  return (
    <CamerasContext.Provider value={v}>
      {props.children}
    </CamerasContext.Provider>
  );
}

export function useCameras() {
  return useContext(CamerasContext)!;
}
