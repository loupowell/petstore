package main

import (
	"net/http"
	"petstore/breeds"
	"petstore/owners"
	"petstore/pets"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Path("/breeds").Queries("group", "{group}").HandlerFunc(breeds.GetBreeds).Methods("GET")
	router.Path("/breeds").Queries("origin-country", "{origin-country}").HandlerFunc(breeds.GetBreeds).Methods("GET")
	router.HandleFunc("/breeds", breeds.GetBreeds).Methods("GET")
	router.HandleFunc("/breeds", breeds.PostBreeds).Methods("POST")
	router.HandleFunc("/breeds/{friendly-id}", breeds.GetBreed).Methods("GET")
	router.HandleFunc("/breeds/{id}", breeds.DeleteBreeds).Methods("DELETE")
	router.HandleFunc("/breeds/{friendly-id}", breeds.UpdateBreed).Methods("PUT")
	router.HandleFunc("/owners", owners.GetOwners).Methods("GET")
	router.HandleFunc("/pets", pets.GetPets).Methods("GET")
	http.ListenAndServe(":8000", router)
}
