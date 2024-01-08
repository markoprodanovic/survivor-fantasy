package db

import (
	"fmt"
	"log"
	"survivor_fantasy/model"

	"github.com/gocraft/dbr"
)

func GetEpisodes(dbSes *dbr.Session) ([]model.Episode, error) {
	var episodes []model.Episode
	stmt := dbSes.Select("*").From("episodes")
	rows, err := stmt.Rows()

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		episode := model.Episode{}
		err = rows.Scan(
			&episode.ID,
			&episode.EpisodeNumber,
			&episode.EpisodeDate,
		)

		if err != nil {
			return nil, err
		}

		episodes = append(episodes, episode)
	}

	return episodes, err
}

func GetEpisode(dbSes *dbr.Session, episodeID int64) (model.Episode, error) {
	var episode model.Episode
	stmt := dbSes.Select("*").From("episodes").Where("id = ?", episodeID)
	err := stmt.LoadOne(&episode)
	if err != nil {
		log.Println(err)
		return episode, err
	}

	return episode, err
}

func GetEpisodePoints(dbSes *dbr.Session, episodeID int64) ([]model.EpisodePoints, error) {
	var episode_points []model.EpisodePoints
	stmt := dbSes.Select("*").From("episode_points").Where("episode_id = ?", episodeID)
	rows, err := stmt.Rows()

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		eps := model.EpisodePoints{}
		err = rows.Scan(
			&eps.ID,
			&eps.EpisodeID,
			&eps.CastID,
			&eps.Points,
		)

		if err != nil {
			return nil, err
		}

		episode_points = append(episode_points, eps)
	}

	return episode_points, err
}

func CreateEpisode(dbSes *dbr.Session, episode *model.Episode) error {

	insertColumns := []string{"episode_number", "episode_date"}

	err := dbSes.InsertInto("episodes").Columns(insertColumns...).Record(episode).Returning("id").Load(&episode.ID)

	return err
}

func CreateEpisodePoints(dbSes *dbr.Session, episode_points *model.EpisodePoints) error {

	insertColumns := []string{"episode_id", "cast_id", "points"}

	err := dbSes.InsertInto("episode_points").Columns(insertColumns...).Record(episode_points).Returning("id").Load(&episode_points.ID)

	return err
}

func DeleteEpisode(dbSes *dbr.Session, episodeID int64) error {

	_, err := dbSes.DeleteFrom("episode_points").Where("episode_id = ?", episodeID).Exec()
	if err != nil {
		return fmt.Errorf("couldn't delete %v from episode_points table: %v", episodeID, err)
	}

	_, err = dbSes.DeleteFrom("episodes").Where("id = ?", episodeID).Exec()
	if err != nil {
		return fmt.Errorf("couldn't delete %v from episodes table: %v", episodeID, err)
	}

	return err
}
