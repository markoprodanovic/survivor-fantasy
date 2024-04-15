package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"survivor_fantasy/db"
	"survivor_fantasy/model"

	"github.com/gocraft/dbr"
	"github.com/gorilla/mux"
)

// helper function
func isUniqueList(list []int64) bool {
	seen := make(map[int64]bool)

	for _, item := range list {
		if seen[item] {
			// If the item is already seen, it's not unique
			return false
		}
		seen[item] = true
	}

	// All items are unique if no duplicates were found
	return true
}

func (app *AppHandler) handleCreateUserPicks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userID"]
	_, err := db.GetUser(app.dbSes, userID)

	// make sure user exists in the database
	if err == dbr.ErrNotFound {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Error fetching user", http.StatusInternalServerError)
		return
	}

	// parse picks
	var picks []int64
	if err := json.NewDecoder(r.Body).Decode(&picks); err != nil {
		http.Error(w, fmt.Sprintf("json decode failure: %v", err), http.StatusBadRequest)
		return
	}

	// if user already has picks delete them
	user_picks, err := db.GetUserPicks(app.dbSes, userID)
	if err != nil {
		http.Error(w, "Error fetching user picks", http.StatusInternalServerError)
		return
	}
	if len(user_picks) != 0 {
		// delete existing picks
		err := db.DeleteUserPicks(app.dbSes, userID)
		if err != nil {
			http.Error(w, "Unable to delete existing user picks", http.StatusInternalServerError)
			return
		}
	}

	// picks list must contain 6 ids
	if len(picks) != 6 {
		http.Error(w, "Invalid operation. User must select 6 picks.", http.StatusBadRequest)
		return
	}

	// return error if pick id's aren't all unique
	if !isUniqueList(picks) {
		http.Error(w, "Invalid player ids. User must pick 6 unique players", http.StatusBadRequest)
		return
	}

	// all 6 pick ids must correspond to players
	players, err := db.GetPlayers(app.dbSes)

	if err != nil {
		http.Error(w, "Error fetching players", http.StatusInternalServerError)
		return
	}

	// Create a map to store player IDs for quick lookup
	playerIDs := make(map[int64]bool)
	for _, player := range players {
		playerIDs[player.ID] = true
	}

	// Create a map to store the number of picks per tribe
	tribePicksCount := make(map[int64]int)

	// Initialize the counts for each tribe to 0
	for _, player := range players {
		tribePicksCount[player.TribeID] = 0
	}

	// Iterate through picks and update tribe pick counts
	for _, pickID := range picks {
		for _, player := range players {
			if player.ID == pickID {
				tribePicksCount[player.TribeID]++
				break
			}
		}
	}

	// Check if any tribe has more than 2 picks
	for tribeID, count := range tribePicksCount {
		if count > 2 {
			http.Error(w, fmt.Sprintf("Too many picks from tribe %d", tribeID), http.StatusBadRequest)
			return
		}
	}

	for _, pickID := range picks {
		if _, exists := playerIDs[pickID]; !exists {
			http.Error(w, fmt.Sprintf("Invalid player ID: %v", pickID), http.StatusBadRequest)
			return
		}
	}

	for _, player_id := range picks {
		user_pick := model.UserPick{
			UserID:   userID,
			PlayerID: player_id,
		}

		err := db.CreateUserPick(app.dbSes, &user_pick)
		if err != nil {
			http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		}
	}

	respondWithJSON(w, r, http.StatusOK, picks)
}
