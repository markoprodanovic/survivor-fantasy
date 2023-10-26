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

	muxRouter.HandleFunc("/api/v1/tribes", app.handleGetTribes).Methods("GET")
	muxRouter.HandleFunc("/api/v1/tribes", app.handleCreateTribe).Methods("POST")
	muxRouter.HandleFunc("/api/v1/tribes/{tribeID:[0-9]}", app.handleGetOneTribe).Methods("GET")
	muxRouter.HandleFunc("/api/v1/tribes/{tribeID:[0-9]}", app.handleDeleteTribe).Methods("DELETE")
	muxRouter.HandleFunc("/api/v1/tribes/{tribeID:[0-9]}", app.handleUpdateTribe).Methods("PUT")

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
