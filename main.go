package main

import (
	"net/http"
	"petstore/breeds"

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
	http.ListenAndServe(":8000", router)
}
