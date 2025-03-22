import { LivePage } from "./pages/dashboard/live";
import { LoginPage } from "./pages/login";
import { NotFoundPage } from "./pages/error";

/**
 * The comments are necessary for build process, take a look at Makefile
 */
export const ROUTES = {
  "/login": LoginPage, // !route
  "/dashboard/live": LivePage, // !route
  "/error/not-found": NotFoundPage, // !route
};

export type Route = keyof typeof ROUTES;
