import React from "react";
import { auth } from "../../../auth";
import { redirect } from "next/navigation";

const UserDashboard = async () => {
  const session = await auth();

  if (!session?.user) {
    redirect("/api/auth/signin");
  }
  return <div>UserDashboard</div>;
};

export default UserDashboard;
