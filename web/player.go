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

func (app *AppHandler) handleGetOnePlayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	playerID, err := strconv.Atoi(vars["playerID"])
	if err != nil {
		http.Error(w, "Invalid player ID", http.StatusBadRequest)
		return
	}
	player, err := db.GetPlayer(app.dbSes, int64(playerID))

	if err == dbr.ErrNotFound {
		http.Error(w, "Player not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Error fetching player", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, r, http.StatusOK, player)
}

func (app *AppHandler) handleGetPlayers(w http.ResponseWriter, r *http.Request) {
	players, err := db.GetPlayers(app.dbSes)

	if err != nil {
		http.Error(w, "Error fetching player", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Range", strconv.Itoa(len(players)))

	respondWithJSON(w, r, http.StatusOK, players)
}

func (app *AppHandler) handleDeletePlayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	playerID, err := strconv.Atoi(vars["playerID"])
	if err != nil {
		http.Error(w, "Invalid player ID", http.StatusBadRequest)
		return
	}

	_, err = db.GetPlayer(app.dbSes, int64(playerID))
	if err != nil {
		http.Error(w, "Player not found", http.StatusNotFound)
		return
	}

	err = db.DeletePlayer(app.dbSes, int64(playerID))
	if err != nil {
		http.Error(w, "Error deleting player", http.StatusInternalServerError)
	}

	log.Printf("Deleted Player %d", playerID)

	respondWithJSON(w, r, http.StatusOK, map[string]string{"id": vars["playerID"], "result": "success"})
}

func (app *AppHandler) handleCreatePlayer(w http.ResponseWriter, r *http.Request) {
	player := model.Player{}
	if err := json.NewDecoder(r.Body).Decode(&player); err != nil {
		http.Error(w, fmt.Sprintf("json decode failure: %v", err), http.StatusBadRequest)
		return
	}

	err := db.CreatePlayer(app.dbSes, &player)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
	}

	respondWithJSON(w, r, http.StatusOK, player)
}

func (app *AppHandler) handleUpdatePlayer(w http.ResponseWriter, r *http.Request) {
	givenPlayer := model.Player{}
	if err := json.NewDecoder(r.Body).Decode(&givenPlayer); err != nil {
		http.Error(w, fmt.Sprintf("json decode failure: %v", err), http.StatusBadRequest)
		return
	}

	player := model.Player{
		ID:         givenPlayer.ID,
		FirstName:  givenPlayer.FirstName,
		LastName:   givenPlayer.LastName,
		Age:        givenPlayer.Age,
		TribeID:    givenPlayer.TribeID,
		Eliminated: givenPlayer.Eliminated,
	}

	existingPlayer, err := db.GetPlayer(app.dbSes, player.ID)
	if err != nil {
		if err == dbr.ErrNotFound {
			http.Error(w, "Player not found", http.StatusNotFound)
			return
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		player.ID = existingPlayer.ID
		if err := db.UpdatePlayer(app.dbSes, &player); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	respondWithJSON(w, r, http.StatusOK, player)
}
