import { ConstructionOutlined } from "@mui/icons-material";
import NextAuth, { NextAuthConfig } from "next-auth";
import GitHub from "next-auth/providers/github";
import { getSession } from "next-auth/react";

export const authConfig = {
  providers: [GitHub],
  callbacks: {
    async redirect({ url, baseUrl }) {
      if (!url.startsWith("http://") && !url.startsWith("https://")) {
        // If in development, prepend the development base URL
        if (process.env.NODE_ENV === "development") {
          const newUrl = `http://localhost:9090${url}`;
          console.log("Modified URL:", newUrl);
          return newUrl;
        }
      }

      return url;
    },
    authorized({ auth, request: { nextUrl } }) {
      console.log("auth: ", auth);
      console.log("nextURL: ", nextUrl);
      const isLoggedIn = !!auth?.user;
      console.log("isLoggedIn: ", isLoggedIn);
      const paths = ["/admin", "/dashboard"];
      const isProtected = paths.some((path) =>
        nextUrl.pathname.startsWith(path)
      );

      if (isProtected && !isLoggedIn) {
        console.log("Rerouting!!!");
        // Adjust the base URL based on the environment
        const baseUrl =
          process.env.NODE_ENV === "development"
            ? "http://localhost:9090"
            : nextUrl.origin;

        const redirectUrl = new URL("api/auth/signin", baseUrl);
        redirectUrl.searchParams.append(
          "callbackUrl",
          baseUrl + nextUrl.pathname
        );

        return Response.redirect(redirectUrl);
      }

      console.log("Should not reroute");

      return "localhost:9090" + nextUrl.pathname;
      // return true;
    },
  },
} satisfies NextAuthConfig;

export const { handlers, auth } = NextAuth(authConfig);
