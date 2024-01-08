import { integer, sqliteTable } from "drizzle-orm/sqlite-core";
import { episodes } from "./episodes";
import { players } from "./players";

export const episodePoints = sqliteTable("episode_points", {
  id: integer("id").primaryKey({ autoIncrement: true }),
  episodeId: integer("episode_id")
    .notNull()
    .references(() => episodes.id),
  castId: integer("cast_id")
    .notNull()
    .references(() => players.id),
  points: integer("points").notNull(),
});
