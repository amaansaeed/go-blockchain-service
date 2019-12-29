package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type app struct {
	Router *mux.Router
	BC     *Blockchain
}

func (a *app) Initialize() {
	// connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)
	// var err error
	// a.DB, err = sql.Open("postgres", connectionString)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Successfully connected to DB!")

	a.BC = NewBlockchain()
	fmt.Println("blockchain created")

	a.Router = mux.NewRouter()

	api := a.Router.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("server healthy"))
	}).Methods(http.MethodGet)

	api.HandleFunc("/blockchain", a.getBlockChain).Methods(http.MethodGet)
	api.HandleFunc("/blockchain", a.addBlock).Methods(http.MethodPost)

	a.Router.Use(middlewareLogger)
}

func (a *app) Run(addr string) {
	fmt.Printf("Server listening on port: %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, a.Router))
}
