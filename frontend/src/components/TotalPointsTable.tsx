import React, { useState, useEffect } from "react";
import { Player } from "../../types/Player";
import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableCell from "@mui/material/TableCell";
import TableContainer from "@mui/material/TableContainer";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import Paper from "@mui/material/Paper";
import { Episode } from "../../types/Episode";

type PlayerTotalPoints = {
  playerName: string;
  totalPoints: number;
};

interface TotalPointsTableProps {
  players: Player[];
  episodes: Episode[];
}

const TotalPointsTable = ({ players, episodes }: TotalPointsTableProps) => {
  const [playersTotalPoints, setPlayersTotalPoints] = useState<
    PlayerTotalPoints[]
  >([]);

  useEffect(() => {
    if (players.length === 0 || episodes.length === 0) return;
    const pointsList = episodes.map((x) => x.points).flat();

    const pTotalPoints: PlayerTotalPoints[] = [];
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
  }, [players, episodes]);
  return (
    <div className="w1/4 mx-auto">
      <h3>Total Points</h3>
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
  );
};

export default TotalPointsTable;
