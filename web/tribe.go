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

type AppHandler struct {
	dbSes *dbr.Session
}

func (app *AppHandler) handleGetOneTribe(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tribeID, err := strconv.Atoi(vars["tribeID"])
	if err != nil {
		http.Error(w, "Invalid tribe ID", http.StatusBadRequest)
		return
	}
	tribe, err := db.GetTribe(app.dbSes, int64(tribeID))

	if err == dbr.ErrNotFound {
		http.Error(w, "Tribe not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Error fetching tribe", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, r, http.StatusOK, tribe)
}

func (app *AppHandler) handleGetTribes(w http.ResponseWriter, r *http.Request) {
	tribes, err := db.GetTribes(app.dbSes)

	if err != nil {
		http.Error(w, "Error fetching tribe", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, r, http.StatusOK, tribes)
}

func (app *AppHandler) handleDeleteTribe(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	tribeID, err := strconv.Atoi(vars["tribeID"])
	if err != nil {
		http.Error(w, "Invalid tribe ID", http.StatusBadRequest)
		return
	}

	_, err = db.GetTribe(app.dbSes, int64(tribeID))
	if err != nil {
		http.Error(w, "Tribe not found", http.StatusNotFound)
		return
	}

	err = db.DeleteTribe(app.dbSes, int64(tribeID))
	if err != nil {
		http.Error(w, "Error deleting tribe", http.StatusInternalServerError)
	}

	log.Printf("Deleted Tribe %d", tribeID)

	respondWithJSON(w, r, http.StatusOK, map[string]string{"id": vars["tribeID"], "result": "success"})
}

func (app *AppHandler) handleCreateTribe(w http.ResponseWriter, r *http.Request) {
	tribe := model.Tribe{}
	if err := json.NewDecoder(r.Body).Decode(&tribe); err != nil {
		http.Error(w, fmt.Sprintf("json decode failure: %v", err), http.StatusBadRequest)
		return
	}

	err := db.CreateTribe(app.dbSes, &tribe)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
	}

	respondWithJSON(w, r, http.StatusOK, tribe)
}

func (app *AppHandler) handleUpdateTribe(w http.ResponseWriter, r *http.Request) {
	givenTribe := model.Tribe{}
	if err := json.NewDecoder(r.Body).Decode(&givenTribe); err != nil {
		http.Error(w, fmt.Sprintf("json decode failure: %v", err), http.StatusBadRequest)
		return
	}

	tribe := model.Tribe{
		ID:     givenTribe.ID,
		Name:   givenTribe.Name,
		Colour: givenTribe.Colour,
	}

	existingTribe, err := db.GetTribe(app.dbSes, tribe.ID)
	if err != nil {
		if err == dbr.ErrNotFound {
			http.Error(w, "Tribe not found", http.StatusNotFound)
			return
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		tribe.ID = existingTribe.ID
		if err := db.UpdateTribe(app.dbSes, &tribe); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	respondWithJSON(w, r, http.StatusOK, tribe)
}
