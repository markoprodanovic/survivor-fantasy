package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"survivor_fantasy/db"
	"survivor_fantasy/model"

	"github.com/gocraft/dbr"
	"github.com/gorilla/mux"
)

func (app *AppHandler) handleGetOneEpisodesWithPoints(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	episodeID, err := strconv.Atoi(vars["episodeID"])
	if err != nil {
		http.Error(w, "Invalid episode ID", http.StatusBadRequest)
		return
	}
	episode, err := db.GetEpisode(app.dbSes, int64(episodeID))

	if err == dbr.ErrNotFound {
		http.Error(w, "Episode not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Error fetching episode", http.StatusInternalServerError)
		return
	}

	episode_points, err := db.GetEpisodePoints(app.dbSes, int64(episodeID))
	points := []model.SimpleEpisodePoint{}

	if err == dbr.ErrNotFound {
		log.Printf("No points for episode")
	} else if err != nil {
		http.Error(w, "Error fetching episode points", http.StatusInternalServerError)
		return
	}

	for _, ep := range episode_points {
		p := model.SimpleEpisodePoint{
			CastID: ep.CastID,
			Points: ep.Points,
		}
		points = append(points, p)
	}
	episode_with_points := model.EpisodeWithPoints{
		ID:            episode.ID,
		EpisodeNumber: episode.EpisodeNumber,
		EpisodeDate:   episode.EpisodeDate,
		Points:        points,
	}

	respondWithJSON(w, r, http.StatusOK, episode_with_points)
}

func (app *AppHandler) handleGetEpisodesWithPoints(w http.ResponseWriter, r *http.Request) {
	episodes, err := db.GetEpisodes(app.dbSes)
	if err != nil {
		http.Error(w, "Error fetching episode", http.StatusInternalServerError)
		return
	}

	eps_with_points := []model.EpisodeWithPoints{}
	for _, episode := range episodes {
		episodeID := episode.ID
		episode_points, err := db.GetEpisodePoints(app.dbSes, int64(episodeID))
		if err != nil {
			http.Error(w, "Error fetching episode points", http.StatusInternalServerError)
			return
		}

		points := []model.SimpleEpisodePoint{}
		for _, ep := range episode_points {
			p := model.SimpleEpisodePoint{
				CastID: ep.CastID,
				Points: ep.Points,
			}
			points = append(points, p)
		}

		episode_with_points := model.EpisodeWithPoints{
			ID:            episodeID,
			EpisodeNumber: episode.EpisodeNumber,
			EpisodeDate:   episode.EpisodeDate,
			Points:        points,
		}

		eps_with_points = append(eps_with_points, episode_with_points)
	}

	w.Header().Set("Content-Range", strconv.Itoa(len(eps_with_points)))

	respondWithJSON(w, r, http.StatusOK, eps_with_points)
}

func (app *AppHandler) handleCreateEpisodeWithPoints(w http.ResponseWriter, r *http.Request) {
	episode_with_points := model.EpisodeWithPoints{}
	if err := json.NewDecoder(r.Body).Decode(&episode_with_points); err != nil {
		http.Error(w, fmt.Sprintf("json decode failure: %v", err), http.StatusBadRequest)
		return
	}

	episode := model.Episode{
		EpisodeNumber: episode_with_points.EpisodeNumber,
		EpisodeDate:   episode_with_points.EpisodeDate,
	}
	err := db.CreateEpisode(app.dbSes, &episode)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
	}

	episode_with_points.ID = episode.ID

	for _, ep := range episode_with_points.Points {
		episode_points := model.EpisodePoints{
			EpisodeID: episode.ID,
			CastID:    ep.CastID,
			Points:    ep.Points,
		}

		err := db.CreateEpisodePoints(app.dbSes, &episode_points)
		if err != nil {
			http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		}
	}

	respondWithJSON(w, r, http.StatusOK, episode_with_points)
}

func (app *AppHandler) handleDeleteEpisodeWithPoints(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	episodeID, err := strconv.Atoi(vars["episodeID"])
	if err != nil {
		http.Error(w, "Invalid episode ID", http.StatusBadRequest)
		return
	}

	_, err = db.GetEpisode(app.dbSes, int64(episodeID))
	if err != nil {
		http.Error(w, "Episode not found", http.StatusNotFound)
		return
	}

	err = db.DeleteEpisode(app.dbSes, int64(episodeID))
	if err != nil {
		http.Error(w, "Error deleting episode", http.StatusInternalServerError)
		return
	}

	log.Printf("Deleted Episode %d", episodeID)

	respondWithJSON(w, r, http.StatusOK, map[string]string{"id": vars["episodeID"], "result": "success"})
}

// func (app *AppHandler) handleUpdateEpisodeWithPoints(w http.ResponseWriter, r *http.Request) {
// 	givenEpisodeWithPoints := model.EpisodeWithPoints{}
// 	if err := json.NewDecoder(r.Body).Decode(&givenEpisodeWithPoints); err != nil {
// 		http.Error(w, fmt.Sprintf("json decode failure: %v", err), http.StatusBadRequest)
// 		return
// 	}

// 	episode := model.Episode{
// 		ID:            givenEpisodeWithPoints.ID,
// 		EpisodeNumber: givenEpisodeWithPoints.EpisodeNumber,
// 		EpisodeDate:   givenEpisodeWithPoints.EpisodeDate,
// 	}

// 	existingTribe, err := db.GetEpisode(app.dbSes, episode.ID)
// 	if err != nil {
// 		if err == dbr.ErrNotFound {
// 			http.Error(w, "Episode not found", http.StatusNotFound)
// 			return
// 		} else {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 	}

// 	for _, ep := range givenEpisodeWithPoints.Points {
// 		existing_episode_points, err := db.GetEpisodePoints(app.dbSes, episode.ID)
// 		if err != nil {
// 			if err == dbr.ErrNotFound {
// 				http.Error(w, "Episode not found", http.StatusNotFound)
// 				return
// 			} else {
// 				http.Error(w, err.Error(), http.StatusInternalServerError)
// 				return
// 			}
// 		}

// 	}

// 	episode.ID = existingTribe.ID
// 	if err := db.UpdateEpisode(app.dbSes, &episode); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	respondWithJSON(w, r, http.StatusOK, tribe)
// }
