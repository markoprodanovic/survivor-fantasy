import { redirect } from "next/navigation";

import { auth } from "@/src/auth";

const HomePage = async () => {
  const session = await auth();
  const baseUrl = process.env.NEXT_PUBLIC_BASE_URL;

  console.log(`redirecting to: ${baseUrl}/admin`);
  redirect(`${baseUrl}/admin`);
  // if (!session?.user) {
  //   redirect(`${baseUrl}/api/auth/signin?callbackUrl=/${baseUrl}admin`);
  // } else {
  //   console.log(`redirecting to: ${baseUrl}/admin`);
  //   redirect(`${baseUrl}/admin`);
  // }
};

export default HomePage;
