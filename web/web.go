package web

import (
	"log"
	"net/http"
	"time"

	"github.com/gocraft/dbr"
	"github.com/gorilla/mux"
)

func MakeMuxRouter(dbSes *dbr.Session) http.Handler {

	app := &AppHandler{dbSes: dbSes}

	muxRouter := mux.NewRouter()

	muxRouter.HandleFunc("/api/v1/tribes", app.handleGetTribes).Methods("GET")
	muxRouter.HandleFunc("/api/v1/tribes", app.handleCreateTribe).Methods("POST")
	muxRouter.HandleFunc("/api/v1/tribes/{tribeID:[0-9]}", app.handleGetOneTribe).Methods("GET")
	muxRouter.HandleFunc("/api/v1/tribes/{tribeID:[0-9]}", app.handleDeleteTribe).Methods("DELETE")
	muxRouter.HandleFunc("/api/v1/tribes/{tribeID:[0-9]}", app.handleUpdateTribe).Methods("PUT")

	return muxRouter
}

func Run(con *dbr.Connection) error {
	httpAddr := ":9090"

	dbSes := con.NewSession(nil)
	mux := MakeMuxRouter(dbSes)

	log.Printf("Listening on %s\n", httpAddr)
	s := &http.Server{
		Addr:           httpAddr,
		Handler:        mux,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Printf("Checkpoint")
	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("Server error: %v", err)
		return err
	}
	log.Printf("Stopping...")

	return nil
}
