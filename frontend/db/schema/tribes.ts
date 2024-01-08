import { integer, text, sqliteTable } from "drizzle-orm/sqlite-core";

export const tribes = sqliteTable("tribes", {
  id: integer("id").primaryKey({ autoIncrement: true }),
  name: text("name").notNull(),
  colour: text("colour").notNull(),
});
