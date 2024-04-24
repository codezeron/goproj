package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/codezeron/apigo/services/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	//subrouter para v1
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	//registrando endpoints de user
	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subrouter)

	log.Println("Server running: ", s.addr)
	return http.ListenAndServe(s.addr, router)
}
