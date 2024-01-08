import NextAuth, { NextAuthConfig } from "next-auth";
import GitHub from "next-auth/providers/github";
import { DrizzleAdapter } from "@auth/drizzle-adapter";
import { users } from "@/db/schema/users";
import { db } from "@/db";

export const authConfig = {
  adapter: DrizzleAdapter(db),
  providers: [GitHub],
  callbacks: {
    async signIn({ user, account, profile }) {
      console.log("User from DB on signIn:", user);
      return true;
    },
    async session({ session, user }) {
      const fullUser = db.select().from(users).get({ id: user.id });
      console.log("database user:");

      session.user.id = user.id;
      session.user.admin = fullUser.admin;
      return session;
    },
    authorized({ auth, request: { nextUrl } }) {
      const isLoggedIn = !!auth?.user;
      const paths = ["/admin", "/dashboard"];
      const isProtected = paths.some((path) =>
        nextUrl.pathname.startsWith(path)
      );

      if (isProtected && !isLoggedIn) {
        const redirectUrl = new URL("api/auth/signin", nextUrl.origin);
        redirectUrl.searchParams.append("callbackUrl", nextUrl.href);
        return Response.redirect(redirectUrl);
      }

      return true;
    },
  },
} satisfies NextAuthConfig;

export const { handlers, auth } = NextAuth(authConfig);
