import React from "react";
import { auth } from "../../../auth";
import { redirect } from "next/navigation";
import TotalPointsTable from "@/src/components/TotalPointsTable";

const fetchPlayers = async () => {
  const res = await fetch("http://localhost:9090/api/v1/players");
  return res.json();
};

const fetchEpisodes = async () => {
  const res = await fetch("http://localhost:9090/api/v1/episodes");
  return res.json();
};

const fetchCurrentUser = async (userID: string) => {
  const res = await fetch(`http://localhost:9090/api/v1/users/${userID}`);
  return res.json();
};

const UserDashboard = async ({
  params,
}: {
  params: { "user-id": "string" };
}) => {
  const session = await auth();

  // fetch data
  const players = await fetchPlayers();
  const episodes = await fetchEpisodes();
  const currentUser = await fetchCurrentUser(params["user-id"]);

  if (!session?.user) {
    redirect("/api/auth/signin");
  }

  console.log(session.user);
  return (
    <div className="flex flex-col justify-center items-center p-5">
      <h1 className="text-2xl font-bold">User Dashboard</h1>
      <p>{JSON.stringify(currentUser)}</p>
      <div className="w-3/4">
        <TotalPointsTable players={players} episodes={episodes} />
      </div>
      <div>{JSON.stringify(players)}</div>
    </div>
  );
};

export default UserDashboard;
