import { integer, text, sqliteTable } from "drizzle-orm/sqlite-core";
import { tribes } from "./tribes";

export const players = sqliteTable("players", {
  id: integer("id").primaryKey({ autoIncrement: true }),
  firstName: text("first_name").notNull(),
  lastName: text("last_name").notNull(),
  age: integer("age").notNull(),
  tribeId: integer("tribe_id").references(() => tribes.id),
  eliminated: integer("eliminated", { mode: "boolean" }).default(false),
});
