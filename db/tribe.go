package db

import (
	"fmt"
	"survivor_fantasy/model"

	"github.com/gocraft/dbr"
)

func GetTribe(dbSes *dbr.Session, tribeID int64) (model.Tribe, error) {
	var tribe model.Tribe
	stmt := dbSes.Select("*").From("tribes").Where("id = ?", tribeID)
	err := stmt.LoadOne(&tribe)
	if err != nil {
		return tribe, err
	}

	return tribe, err
}

func GetTribes(dbSes *dbr.Session) ([]model.Tribe, error) {
	var tribes []model.Tribe
	stmt := dbSes.Select("*").From("tribes")
	rows, err := stmt.Rows()

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		tribe := model.Tribe{}
		err = rows.Scan(
			&tribe.ID,
			&tribe.Name,
			&tribe.Colour,
		)

		if err != nil {
			return nil, err
		}

		tribes = append(tribes, tribe)
	}

	return tribes, err
}

func DeleteTribe(dbSes *dbr.Session, tribeID int64) error {

	_, err := dbSes.DeleteFrom("tribes").Where("id = ?", tribeID).Exec()

	if err != nil {
		return fmt.Errorf("couldn't delete %v from tribe table: %v", tribeID, err)
	}

	return err
}

func CreateTribe(dbSes *dbr.Session, tribe *model.Tribe) error {

	insertColumns := []string{"name", "colour"}

	err := dbSes.InsertInto("tribes").Columns(insertColumns...).Record(tribe).Returning("id").Load(&tribe.ID)

	return err
}

func UpdateTribe(dbSes *dbr.Session, tribe *model.Tribe) error {
	fields := map[string]interface{}{
		"id":     tribe.ID,
		"name":   tribe.Name,
		"colour": tribe.Colour,
	}

	_, err := dbSes.Update("tribes").SetMap(fields).Where("id = ?", tribe.ID).Exec()
	return err
}
