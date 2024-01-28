package db

import (
	"fmt"
	"log"
	"survivor_fantasy/model"

	"github.com/gocraft/dbr"
)

func GetUsers(dbSes *dbr.Session) ([]model.User, error) {
	var users []model.User
	stmt := dbSes.Select("*").From("user")
	rows, err := stmt.Rows()
	if err != nil {
		log.Println("Error getting users")
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		log.Println("Starting scan for user")
		user := model.User{}
		err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.EmailVerified,
			&user.Image,
			&user.IsAdmin,
		)

		if err != nil {
			log.Println("Error with scan")
			log.Println(err)
			return nil, err
		}

		users = append(users, user)
	}

	return users, err
}

func GetUserPicks(dbSes *dbr.Session, userID string) ([]model.UserPick, error) {
	var user_picks []model.UserPick
	stmt := dbSes.Select("*").From("user_picks").Where("user_id = ?", userID)
	rows, err := stmt.Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		pick := model.UserPick{}
		err = rows.Scan(
			&pick.ID,
			&pick.UserID,
			&pick.PlayerID,
		)

		if err != nil {
			return nil, err
		}

		user_picks = append(user_picks, pick)
	}

	return user_picks, err
}

func CreateUser(dbSes *dbr.Session, user *model.User) error {
	insertColumns := []string{"first_name", "last_name", "email"}

	err := dbSes.InsertInto("users").Columns(insertColumns...).Record(user).Returning("id").Load(&user.ID)

	return err
}

func CreateUserPick(dbSes *dbr.Session, user_pick *model.UserPick) error {
	insertColumns := []string{"user_id", "player_id"}

	err := dbSes.InsertInto("user_picks").Columns(insertColumns...).Record(user_pick).Returning("id").Load(&user_pick.ID)

	return err
}

func GetUser(dbSes *dbr.Session, userID string) (model.User, error) {
	var user model.User
	stmt := dbSes.Select("*").From("user").Where("id = ?", userID)
	err := stmt.LoadOne(&user)
	if err != nil {
		return user, err
	}

	return user, err
}

func DeleteUser(dbSes *dbr.Session, userID string) error {
	_, err := dbSes.DeleteFrom("users").Where("id = ?", userID).Exec()
	if err != nil {
		return fmt.Errorf("couldn't delete %v from user table: %v", userID, err)
	}

	return err
}

func UpdateUser(dbSes *dbr.Session, user *model.User) error {
	fields := map[string]interface{}{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
	}

	_, err := dbSes.Update("users").SetMap(fields).Where("id = ?", user.ID).Exec()
	return err
}
