package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"petstore/breeds"
	"petstore/owners"
	"petstore/pets"

	"github.com/gorilla/mux"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var message = "Hello Pet Lovers"
	json.NewEncoder(w).Encode(message)
}

func main() {

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}
	router := mux.NewRouter()
	router.HandleFunc("/", helloWorld).Methods("GET")
	router.Path("/breeds").Queries("group", "{group}").HandlerFunc(breeds.GetBreeds).Methods("GET")
	router.Path("/breeds").Queries("origin-country", "{origin-country}").HandlerFunc(breeds.GetBreeds).Methods("GET")
	router.HandleFunc("/breeds", breeds.GetBreeds).Methods("GET")
	router.HandleFunc("/breeds", breeds.PostBreeds).Methods("POST")
	router.HandleFunc("/breeds/{friendly-id}", breeds.GetBreed).Methods("GET")
	router.HandleFunc("/breeds/{id}", breeds.DeleteBreeds).Methods("DELETE")
	router.HandleFunc("/breeds/{friendly-id}", breeds.UpdateBreed).Methods("PUT")
	router.HandleFunc("/owners", owners.GetOwners).Methods("GET")
	router.HandleFunc("/owners/{id}", owners.GetOwner).Methods("GET")
	router.HandleFunc("/owners", owners.PostOwners).Methods("POST")
	router.HandleFunc("/pets", pets.GetPets).Methods("GET")
	http.ListenAndServe(":"+port, router)
}
