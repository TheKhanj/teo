import { LivePage } from "./pages/dashboard/live";
import { NotFoundPage } from "./pages/error";
import { RecordingsPage } from "./pages/dashboard/recordings";

/**
 * The comments are necessary for build process, take a look at Makefile
 */
export const ROUTES = {
  "/dashboard/live": LivePage, // !route
  "/error/not-found": NotFoundPage, // !route
  "/dashboard/recordings": RecordingsPage, // !route
};

export type Route = keyof typeof ROUTES;
