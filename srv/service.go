package srv

import (
	"context"
	"net/http"

	"github.com/sqsinformatique/backend/cfg"
	"github.com/sqsinformatique/backend/utils"

	"github.com/gorilla/mux"

	"github.com/sqsinformatique/backend/db"
)

var (
	srv *http.Server
)

type Server struct {
	R *mux.Router
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Stop here if its Preflighted OPTIONS request
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization,X-Request-Id")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// Lets Gorilla work
	s.R.ServeHTTP(w, r)
}

func Start() (err error) {
	// Init web-server
	router := mux.NewRouter()

	router.HandleFunc("/health", healthGetHandler).Methods("GET")

	router.HandleFunc("/api/models", modelsGetAllHandler).Methods("GET")
	router.HandleFunc("/api/models", modelsPostHandler).Methods("POST")
	router.HandleFunc("/api/models", modelsPutHandler).Methods("PUT")
	router.HandleFunc("/api/models/{id}", modelsGetHandler).Methods("GET")

	router.HandleFunc("/api/nodes", nodesGetAllHandler).Methods("GET")
	router.HandleFunc("/api/nodes", nodesPostHandler).Methods("POST")
	router.HandleFunc("/api/nodes/{id}", nodesGetHandler).Methods("GET")

	router.HandleFunc("/api/elevators", elevatorsGetAllHandler).Methods("GET")
	router.HandleFunc("/api/elevators", elevatorsPostHandler).Methods("POST")
	router.HandleFunc("/api/elevators/{id}", elevatorsGetHandler).Methods("GET")
	router.HandleFunc("/api/elevators/{id}/contractor", elevatorContractorGetHandler).Methods("GET")
	router.HandleFunc("/api/elevators/{id}/works", elevatorWorksGetHandler).Methods("GET")

	router.HandleFunc("/api/contractors", contractorsGetAllHandler).Methods("GET")
	router.HandleFunc("/api/contractors", contractorsPostHandler).Methods("POST")
	router.HandleFunc("/api/contractors/{id}", contractorsGetHandler).Methods("GET")

	router.HandleFunc("/api/worktype", worktypeGetAllHandler).Methods("GET")
	router.HandleFunc("/api/worktype", worktypePostHandler).Methods("POST")
	router.HandleFunc("/api/worktype/{id}", worktypeGetHandler).Methods("GET")

	router.HandleFunc("/api/works", worksPostHandler).Methods("POST")

	router.HandleFunc("/api/predictions/{id}/{fileid}", predictionGetHandler).Methods("GET")

	// Initialize our database connection
	err = db.InitDB()
	if err != nil {
		return
	}

	utils.Info("Service is listening...")
	port := cfg.DefaultCfg.Port
	utils.Infof("http://%s:%s", utils.IP, port)
	srv = &http.Server{Addr: ":" + port, Handler: &Server{router}}
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			utils.Fatal("ListenAndServe:", err)
		}
	}()

	return
}

func Stop() {
	defer db.CloseDB()
	if srv == nil {
		return
	}
	utils.Info("Stoping Service...")
	if err := srv.Shutdown(context.TODO()); err != nil {
		panic(err) // failure/timeout shutting down the server gracefully
	}
	utils.Info("Service stopped")
}
