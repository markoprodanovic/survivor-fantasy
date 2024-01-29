import { useEffect, useState } from "react";
import TotalPointsTable from "./TotalPointsTable";
import { Player } from "../../types/Player";
import { Episode } from "../../types/Episode";
import { useSession } from "next-auth/react";
import Link from "next/link";

export const Home = (props: any) => {
  const [players, setPlayers] = useState<Player[]>([]);
  const [episodes, setEpisodes] = useState<Episode[]>([]);
  const { data: session } = useSession();

  useEffect(() => {
    fetch("http://localhost:9090/api/v1/players")
      .then((res) => res.json())
      .then((data) => {
        setPlayers(data);
      });
  }, []);

  useEffect(() => {
    fetch("http://localhost:9090/api/v1/episodes")
      .then((res) => res.json())
      .then((data) => {
        setEpisodes(data);
      });
  }, []);

  return (
    <>
      <div className="flex flex-col items-center justify-center">
        <TotalPointsTable players={players} episodes={episodes} />
        <Link
          className="mt-8 bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
          href={`/dashboard/${session.user.id}`}
        >
          User Dashboard
        </Link>
      </div>
    </>
  );
};
