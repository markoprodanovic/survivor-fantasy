import { integer, text, sqliteTable } from "drizzle-orm/sqlite-core";
import { users } from "./users";
import { players } from "./players";

export const userPicks = sqliteTable("user_picks", {
  id: integer("id").primaryKey({ autoIncrement: true }),
  userId: text("user_id")
    .notNull()
    .references(() => users.id),
  playerId: integer("player_id")
    .notNull()
    .references(() => players.id),
});
