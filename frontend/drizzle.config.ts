import type { Config } from "drizzle-kit";

export default {
  schema: "./db/schema/*",
  driver: "better-sqlite",
  out: "./drizzle",
  dbCredentials: {
    url: "../survivor_fantasy.db",
  },
} satisfies Config;
