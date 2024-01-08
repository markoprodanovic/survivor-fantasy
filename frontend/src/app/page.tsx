import { redirect } from "next/navigation";

import { auth } from "@/src/auth";

const HomePage = async () => {
  const session = await auth();

  if (!session?.user) {
    redirect("/api/auth/signin?callbackUrl=/");
  } else {
    if (session.user.admin) {
      redirect("/admin");
    }
    redirect(`/dashboard/${session.user.id}`);
  }
};

export default HomePage;
