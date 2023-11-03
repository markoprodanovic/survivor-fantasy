package web

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/gocraft/dbr"
	"github.com/gorilla/mux"
)

var Version = "dev"

func MakeMuxRouter(dbSes *dbr.Session) http.Handler {

	app := &AppHandler{dbSes: dbSes}

	muxRouter := mux.NewRouter()

	// tribes
	muxRouter.HandleFunc("/api/v1/tribes", app.handleGetTribes).Methods("GET")
	muxRouter.HandleFunc("/api/v1/tribes", app.handleCreateTribe).Methods("POST")
	muxRouter.HandleFunc("/api/v1/tribes/{tribeID:[0-9]}", app.handleGetOneTribe).Methods("GET")
	muxRouter.HandleFunc("/api/v1/tribes/{tribeID:[0-9]}", app.handleDeleteTribe).Methods("DELETE")
	muxRouter.HandleFunc("/api/v1/tribes/{tribeID:[0-9]}", app.handleUpdateTribe).Methods("PUT")

	// players
	muxRouter.HandleFunc("/api/v1/players", app.handleGetPlayers).Methods("GET")
	muxRouter.HandleFunc("/api/v1/players", app.handleCreatePlayer).Methods("POST")
	muxRouter.HandleFunc("/api/v1/players/{playerID:[0-9]}", app.handleGetOnePlayer).Methods("GET")
	muxRouter.HandleFunc("/api/v1/players/{playerID:[0-9]}", app.handleDeletePlayer).Methods("DELETE")
	muxRouter.HandleFunc("/api/v1/players/{playerID:[0-9]}", app.handleUpdatePlayer).Methods("PUT")

	// episodes (with points)
	muxRouter.HandleFunc("/api/v1/episodes", app.handleGetEpisodesWithPoints).Methods("GET")
	muxRouter.HandleFunc("/api/v1/episodes", app.handleCreateEpisodeWithPoints).Methods("POST")
	muxRouter.HandleFunc("/api/v1/episodes/{episodeID:[0-9]}", app.handleGetOneEpisodesWithPoints).Methods("GET")
	muxRouter.HandleFunc("/api/v1/episodes/{episodeID:[0-9]}", app.handleDeleteEpisodeWithPoints).Methods("DELETE")
	// muxRouter.HandleFunc("/api/v1/episodes/{episodeID:[0-9]}", app.handleUpdateEpisodeWithPoints).Methods("UPDATE")

	// users
	muxRouter.HandleFunc("/api/v1/users", app.handleGetUsersWithPicks).Methods("GET")
	muxRouter.HandleFunc("/api/v1/users", app.handleCreateUserWithPicks).Methods("POST")
	muxRouter.HandleFunc("/api/v1/users/{userID:[0-9]}", app.handleGetOneUser).Methods("GET")
	muxRouter.HandleFunc("/api/v1/users/{userID:[0-9]}", app.handleDeleteUser).Methods("DELETE")
	muxRouter.HandleFunc("/api/v1/users/{userID:[0-9]}", app.handleUpdateUser).Methods("PUT")

	// Frontend HTML and related assets
	if Version == "dev" {
		// In dev mode, we proxy everything else to React's Webpack dev server
		frontendServer := "http://localhost:3000"
		remote, err := url.Parse(frontendServer)
		if err != nil {
			panic(err)
		}
		log.Printf("Development mode: Proxying / to " + frontendServer)

		proxy := httputil.NewSingleHostReverseProxy(remote)
		proxyHandler := func(w http.ResponseWriter, r *http.Request) {
			r.URL.Path = mux.Vars(r)["rest"]
			proxy.ServeHTTP(w, r)
		}
		muxRouter.HandleFunc("/{rest:.*}", proxyHandler)
	} else {
		assets := []string{}
		for _, p := range assets {
			muxRouter.PathPrefix(p).Handler(http.FileServer(http.Dir("./frontend/build/")))
		}

		// Everything else serves up index.html (and relies on client-side routing)
		muxRouter.HandleFunc("/{rest:.*}", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("cache-control", "no-cache")
			http.ServeFile(w, r, "./frontend/build/index.html")
		})
	}

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
	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("Server error: %v", err)
		return err
	}
	log.Printf("Stopping...")

	return nil
}
