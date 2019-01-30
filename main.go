package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type customer struct {
	ID        string `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,motiempty"`
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(customers)
}

var customers []customer

func main() {
	router := mux.NewRouter()
	customers = append(customers, customer{ID: "1", Firstname: "John", Lastname: "Smith"})
	router.HandleFunc("/customers", getCustomers).Methods("GET")
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
