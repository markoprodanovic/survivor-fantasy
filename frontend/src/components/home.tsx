import { useEffect, useState } from "react";
import * as React from "react";
import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableCell from "@mui/material/TableCell";
import TableContainer from "@mui/material/TableContainer";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import Paper from "@mui/material/Paper";
import { Rowdies } from "next/font/google";

export const Home = (props: any) => {
  const [players, setPlayers] = useState([]);
  const [episodesWithPoints, setEpisodesWithPoints] = useState([]);
  const [playersTotalPoints, setPlayersTotalPoints] = useState([]);
  // {playerID: 3, totalPoints: 36}

  useEffect(() => {
    fetch("/api/v1/players")
      .then((res) => res.json())
      .then((data) => {
        setPlayers(data);
      });
  }, []);

  useEffect(() => {
    fetch("/api/v1/episodes")
      .then((res) => res.json())
      .then((data) => {
        setEpisodesWithPoints(data);
      });
  }, []);

  useEffect(() => {
    if (players.length === 0 || episodesWithPoints.length === 0) return;
    const pointsList = episodesWithPoints.map((x) => x.points).flat();

    const pTotalPoints = [];
    players.forEach((player) => {
      console.log(player);
      const playerPoints = pointsList.filter((x) => x.castId === player.id);
      let totalPoints = 0;
      if (playerPoints.length !== 0) {
        totalPoints = playerPoints
          .map((x) => x.points)
          .reduce((total, curr) => total + curr);
      }
      pTotalPoints.push({
        playerName: `${player.first_name} ${player.last_name}`,
        totalPoints: totalPoints,
      });
    });
    pTotalPoints.sort((a, b) => {
      return b.totalPoints - a.totalPoints;
    });
    setPlayersTotalPoints(pTotalPoints);
  }, [players, episodesWithPoints]);

  return (
    <div className="flex justify-center">
      <div className="w1/4 mx-auto">
        <TableContainer component={Paper}>
          <Table sx={{ minWidth: 650 }} aria-label="simple table">
            <TableHead>
              <TableRow>
                <TableCell>Player</TableCell>
                <TableCell align="right">Total Points</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {playersTotalPoints.map((row) => (
                <TableRow
                  key={row.playerName}
                  sx={{ "&:last-child td, &:last-child th": { border: 0 } }}
                >
                  <TableCell component="th" scope="row">
                    {row.playerName}
                  </TableCell>
                  <TableCell align="right">{row.totalPoints}</TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      </div>
    </div>
  );
};
