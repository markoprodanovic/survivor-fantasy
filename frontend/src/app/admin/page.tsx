import { NextPage } from "next";
import dynamic from "next/dynamic";
import { auth } from "../../auth";
import { redirect } from "next/navigation";
import { SessionProvider } from "next-auth/react";

const DynamicAdminApp = dynamic(() => import("../../components/AdminApp"), {
  ssr: false,
  loading: () => <p>Loading ...</p>,
});

const Admin: NextPage = async () => {
  const session = await auth();

  if (!session?.user) {
    redirect("/api/auth/signin");
  }

  if (!session.user.admin) {
    redirect(`/dashboard/${session.user.id}`);
  }
  return (
    <SessionProvider session={session}>
      <DynamicAdminApp />
    </SessionProvider>
  );
};

export default Admin;
