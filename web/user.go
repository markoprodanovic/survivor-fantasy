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

func (app *AppHandler) handleGetUsersWithPicks(w http.ResponseWriter, r *http.Request) {
	users, err := db.GetUsers(app.dbSes)
	if err != nil {
		http.Error(w, "Error fetching users", http.StatusInternalServerError)
		return
	}

	users_with_picks := []model.UserWithPicks{}
	for _, user := range users {
		userID := user.ID
		user_picks, err := db.GetUserPicks(app.dbSes, int64(userID))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		picks := []int64{}
		if err == dbr.ErrNotFound {
			log.Printf("No points for episode")
		} else if err != nil {
			http.Error(w, "Error fetching episode points", http.StatusInternalServerError)
			return
		}

		for _, p := range user_picks {
			picks = append(picks, p.PlayerID)
		}

		user_with_picks := model.UserWithPicks{
			ID:        userID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			PlayerIDs: picks,
		}

		users_with_picks = append(users_with_picks, user_with_picks)
	}

	w.Header().Set("Content-Range", strconv.Itoa(len(users_with_picks)))

	respondWithJSON(w, r, http.StatusOK, users_with_picks)
}

func (app *AppHandler) handleCreateUserWithPicks(w http.ResponseWriter, r *http.Request) {
	user_with_picks := model.UserWithPicks{}
	if err := json.NewDecoder(r.Body).Decode(&user_with_picks); err != nil {
		http.Error(w, fmt.Sprintf("json decode failure: %v", err), http.StatusBadRequest)
		return
	}

	user := model.User{
		FirstName: user_with_picks.FirstName,
		LastName:  user_with_picks.LastName,
		Email:     user_with_picks.Email,
	}
	err := db.CreateUser(app.dbSes, &user)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
	}

	user_with_picks.ID = user.ID

	for _, player_id := range user_with_picks.PlayerIDs {
		user_pick := model.UserPick{
			UserID:   user.ID,
			PlayerID: player_id,
		}

		err := db.CreateUserPick(app.dbSes, &user_pick)
		if err != nil {
			http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		}
	}

	respondWithJSON(w, r, http.StatusOK, user_with_picks)
}

func (app *AppHandler) handleGetOneUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["userID"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	user, err := db.GetUser(app.dbSes, int64(userID))

	if err == dbr.ErrNotFound {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Error fetching user", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, r, http.StatusOK, user)
}

func (app *AppHandler) handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userID, err := strconv.Atoi(vars["userID"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	_, err = db.GetUser(app.dbSes, int64(userID))
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	err = db.DeleteUser(app.dbSes, int64(userID))
	if err != nil {
		http.Error(w, "Error deleting user", http.StatusInternalServerError)
	}

	log.Printf("Deleted User %d", userID)

	respondWithJSON(w, r, http.StatusOK, map[string]string{"id": vars["userID"], "result": "success"})
}

func (app *AppHandler) handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	givenUser := model.User{}
	if err := json.NewDecoder(r.Body).Decode(&givenUser); err != nil {
		http.Error(w, fmt.Sprintf("json decode failure: %v", err), http.StatusBadRequest)
		return
	}

	user := model.User{
		ID:        givenUser.ID,
		FirstName: givenUser.FirstName,
		LastName:  givenUser.LastName,
		Email:     givenUser.Email,
	}

	existingUser, err := db.GetUser(app.dbSes, user.ID)
	if err != nil {
		if err == dbr.ErrNotFound {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		user.ID = existingUser.ID
		if err := db.UpdateUser(app.dbSes, &user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	respondWithJSON(w, r, http.StatusOK, user)
}
