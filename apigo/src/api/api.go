package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/codezeron/apigo/services/cart"
	"github.com/codezeron/apigo/services/order"
	"github.com/codezeron/apigo/services/product"
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
	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	//registrando endpoints de product
	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(productStore, userStore)
	productHandler.RegisterRoutes(subrouter)

	//registrando endpoints de order
	orderStore := order.NewStore(s.db)
	//registrando endpoints de cart

	cartHandler := cart.NewHandler(productStore, orderStore, userStore)
	cartHandler.RegisterRoutes(subrouter)

	// Serve static files
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

	log.Println("Server running: ", s.addr)
	return http.ListenAndServe(s.addr, router)
}
