import { useEffect, useState } from "react";
import TotalPointsTable from "./TotalPointsTable";
import { Player } from "../../types/Player";
import { Episode } from "../../types/Episode";

export const Home = (props: any) => {
  const [players, setPlayers] = useState<Player[]>([]);
  const [episodes, setEpisodes] = useState<Episode[]>([]);

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
    <div className="flex justify-center">
      <TotalPointsTable players={players} episodes={episodes} />
    </div>
  );
};
