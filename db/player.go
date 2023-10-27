package db

import (
	"fmt"
	"survivor_fantasy/model"

	"github.com/gocraft/dbr"
)

func GetPlayer(dbSes *dbr.Session, playerID int64) (model.Player, error) {
	var player model.Player
	stmt := dbSes.Select("*").From("players").Where("id = ?", playerID)
	err := stmt.LoadOne(&player)
	if err != nil {
		return player, err
	}

	return player, err
}

func GetPlayers(dbSes *dbr.Session) ([]model.Player, error) {
	var players []model.Player
	stmt := dbSes.Select("*").From("players")
	rows, err := stmt.Rows()

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		player := model.Player{}
		err = rows.Scan(
			&player.ID,
			&player.FirstName,
			&player.LastName,
			&player.Age,
			&player.TribeID,
			&player.Eliminated,
		)

		if err != nil {
			return nil, err
		}

		players = append(players, player)
	}

	return players, err
}

func DeletePlayer(dbSes *dbr.Session, playerID int64) error {

	_, err := dbSes.DeleteFrom("players").Where("id = ?", playerID).Exec()

	if err != nil {
		return fmt.Errorf("couldn't delete %v from players table: %v", playerID, err)
	}

	return err
}

func CreatePlayer(dbSes *dbr.Session, player *model.Player) error {

	insertColumns := []string{"first_name", "last_name", "age", "tribe_id", "eliminated"}

	err := dbSes.InsertInto("players").Columns(insertColumns...).Record(player).Returning("id").Load(&player.ID)

	return err
}

func UpdatePlayer(dbSes *dbr.Session, player *model.Player) error {
	fields := map[string]interface{}{
		"id":         player.ID,
		"first_name": player.FirstName,
		"last_name":  player.LastName,
		"age":        player.Age,
		"tribe_id":   player.TribeID,
		"eliminated": player.Eliminated,
	}

	_, err := dbSes.Update("players").SetMap(fields).Where("id = ?", player.ID).Exec()
	return err
}
