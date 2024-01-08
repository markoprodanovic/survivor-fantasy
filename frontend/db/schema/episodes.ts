import { integer, text, sqliteTable } from "drizzle-orm/sqlite-core";

export const episodes = sqliteTable("episodes", {
  id: integer("id").primaryKey({ autoIncrement: true }),
  episodeNumber: integer("episode_number").notNull(),
  episodeDate: text("episode_date").notNull(),
});
