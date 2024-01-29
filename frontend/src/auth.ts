import NextAuth, { NextAuthConfig } from "next-auth";
import GitHub from "next-auth/providers/github";
import Google from "next-auth/providers/google";
import { DrizzleAdapter } from "@auth/drizzle-adapter";
import { users } from "@/db/schema/users";
import { eq } from "drizzle-orm";
import { db } from "@/db";

export const authConfig = {
  adapter: DrizzleAdapter(db),
  providers: [GitHub, Google],
  callbacks: {
    async session({ session, token }) {
      if (session.user && token.sub) {
        session.user.id = token.sub;
      }
      if (session.user && token.admin) {
        // @ts-ignore
        session.user.admin = token.admin;
      }
      return session;
    },
    async jwt({ token }) {
      if (!token.sub) return token;
      const userQueryResult = await db
        .select()
        .from(users)
        .where(eq(users.id, token.sub));
      const existingUser = userQueryResult[0];
      if (!existingUser) return token;
      token.admin = existingUser.admin;
      return token;
    },
  },
  session: { strategy: "jwt" },
} satisfies NextAuthConfig;

export const { handlers, auth } = NextAuth(authConfig);
