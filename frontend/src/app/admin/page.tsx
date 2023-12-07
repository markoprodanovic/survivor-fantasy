// "use client";

import { NextPage } from "next";
import dynamic from "next/dynamic";
import { auth } from "../../auth";
import { redirect } from "next/navigation";

const DynamicAdminApp = dynamic(() => import("../../components/AdminApp"), {
  ssr: false,
  loading: () => <p>Loading ...</p>,
});

const Admin: NextPage = async () => {
  const session = await auth();

  if (!session?.user) {
    redirect("/api/auth/signin");
  }

  if (!session.user.isAdmin) {
    redirect(`/dashboard/${session.user.id}`);
  }
  return <DynamicAdminApp user={session.user} />;
};

export default Admin;
