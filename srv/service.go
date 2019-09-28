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

	// ref_type := router.GetRoute("/api/ref_type")

	// ref_type.HandleFunc("/resource", ref_resourceGetAllHandler).Methods("GET")
	// ref_type.HandleFunc("/resource", ref_resourcePostHandler).Methods("POST")
	// ref_type.HandleFunc("/resource/{id}", ref_resourceDeleteHandler).Methods("DELETE")
	// ref_type.HandleFunc("/resource/{id}", ref_resourceGetHandler).Methods("GET")

	// ref_type.HandleFunc("/objects", ref_objectsGetAllHandler).Methods("GET")
	// ref_type.HandleFunc("/objects", ref_objectsPostHandler).Methods("POST")
	// ref_type.HandleFunc("/objects/{id}", ref_objectsDeleteHandler).Methods("DELETE")
	// ref_type.HandleFunc("/objects/{id}", ref_objectsGetHandler).Methods("GET")

	// ref_type.HandleFunc("/status", ref_statusGetAllHandler).Methods("GET")
	// ref_type.HandleFunc("/status", ref_statusPostHandler).Methods("POST")
	// ref_type.HandleFunc("/status/{id}", ref_statusDeleteHandler).Methods("DELETE")
	// ref_type.HandleFunc("/status/{id}", statusGetHandler).Methods("GET")

	// ref_type.HandleFunc("/maintenance", ref_maintenanceGetAllHandler).Methods("GET")
	// ref_type.HandleFunc("/maintenance", ref_maintenancePostHandler).Methods("POST")
	// ref_type.HandleFunc("/maintenance/{id}", ref_maintenanceDeleteHandler).Methods("DELETE")
	// ref_type.HandleFunc("/maintenance/{id}", ref_maintenanceGetHandler).Methods("GET")

	// router.HandleFunc("/api/supply_organizations", supply_organizationsGetAllHandler).Methods("GET")
	router.HandleFunc("/api/supply_organizations", supply_organizationsPostHandler).Methods("POST")
	// router.HandleFunc("/api/supply_organizations/{id}", supply_organizationsDeleteHandler).Methods("DELETE")
	router.HandleFunc("/api/supply_organizations/{id}", supply_organizationsGetHandler).Methods("GET")

	// router.HandleFunc("/api/incidents", incidentsGetAllHandler).Methods("GET")
	// router.HandleFunc("/api/incidents/{supply_organizations_id}", incidentsPostHandler).Methods("POST")
	// router.HandleFunc("/api/incidents/{supply_organizations_id}/{object_id}", incidentsDeleteHandler).Methods("DELETE")
	// router.HandleFunc("/api/incidents/{supply_organizations_id}/{object_id}", incidentsGetHandler).Methods("GET")

	router.HandleFunc("/api/objects", objectsGetAllHandler).Methods("GET")
	router.HandleFunc("/api/objects/{type}", objectsGetAllByTypeHandler).Methods("GET")
	router.HandleFunc("/api/objects/{supply_organizations_id}", objectsPostHandler).Methods("POST")
	// router.HandleFunc("/api/objects/{supply_organizations_id}/{object_id}", objectsDeleteHandler).Methods("DELETE")
	// router.HandleFunc("/api/objects/{supply_organizations_id}/{object_id}", objectsGetHandler).Methods("GET")

	// router.HandleFunc("/api/maintenance", maintenanceGetAllHandler).Methods("GET")
	// router.HandleFunc("/api/maintenance/{supply_organizations_id}", maintenancePostHandler).Methods("POST")
	// router.HandleFunc("/api/maintenance/{supply_organizations_id}/{object_id}", maintenanceDeleteHandler).Methods("DELETE")
	// router.HandleFunc("/api/maintenance/{supply_organizations_id}/{object_id}", maintenanceGetHandler).Methods("GET")

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
