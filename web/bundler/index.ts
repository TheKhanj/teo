import { build } from "esbuild";
import { solidPlugin } from "esbuild-plugin-solid";

const env = process.env.NODE_ENV ?? "dev";

build({
  entryPoints: ["src/index.tsx"],
  bundle: true,
  sourcemap: true,
  outfile: "cdn/bundle.js",
  plugins: [solidPlugin()],
  define: {
    "process.env.NODE_ENV": JSON.stringify(env),
  },
  conditions: [env],
}).catch(() => process.exit(1));
