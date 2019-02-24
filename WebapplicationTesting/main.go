package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()
	
	
	r.HandleFunc("/bird", getBirdHandler).Methods("GET")
	r.HandleFunc("/bird", createBirdHandler).Methods("POST")
	return r
}

func main() {
	fmt.Println("Starting server...")
	db, err := sql.Open("postgres","postgres://postgres:root@localhost/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	InitStore(&dbStore{db: db})

	r := newRouter()
	fmt.Println("Serving on port 8080")
	http.ListenAndServe(":8080", r)
}

