package breeds

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"

	"petstore/shared"

	"github.com/gorilla/mux"
)

type Breed struct {
	ID         string `json:"id"`
	FriendlyID string `json:"friendly-id" validate:"required"`
	Breed      string `json:"breed" validate:"required"`
	Group      string `json:"group" validate:"required,oneof=herding hound toy non-sporting terrier working miscellaneous foundation-stock-service"`
	Origin     string `json:"origin-country,omitempty"`
}

var breedsList []Breed

//Retrieve the entire breeds collection
func GetBreeds(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Correlation-Key", shared.GetID())
	w.Header().Set("API-Version", shared.Version)
	if len(breedsList) == 0 {
		loadBreeds, _ := ioutil.ReadFile("breeds/breedslist.json")
		err := json.Unmarshal([]byte(loadBreeds), &breedsList)
		if err != nil {
			fmt.Println(err)
		}
	}

	//IF there is a Group Filter, pull only those records
	params := mux.Vars(r)
	if params["group"] != "" {
		var filterdBreedsList []Breed
		for index, item := range breedsList {
			if item.Group == params["group"] {
				filterdBreedsList = append(filterdBreedsList, breedsList[index])
			}
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(filterdBreedsList)
		return
	}

	//IF there is an Origin Country Filter, pull only those records
	if params["origin-country"] != "" {
		var filterdBreedsList []Breed
		for index, item := range breedsList {
			if item.Origin == params["origin-country"] {
				filterdBreedsList = append(filterdBreedsList, breedsList[index])
			}
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(filterdBreedsList)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(breedsList)
}

//GetBreed Retrieve a specific breed based on friendly-id
func GetBreed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Correlation-Key", shared.GetID())
	w.Header().Set("API-Version", shared.Version)
	params := mux.Vars(r)
	for index, item := range breedsList {
		if item.FriendlyID == params["friendly-id"] {
			requestedRecord := breedsList[index]
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(requestedRecord)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

//Add a new breed to the collection
func PostBreeds(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Correlation-Key", shared.GetID())
	w.Header().Set("API-Version", shared.Version)
	var newBreed Breed
	_ = json.NewDecoder(r.Body).Decode(&newBreed)
	newBreed.ID = strconv.Itoa(rand.Intn(100000000))

	//Validate the posted record
	if ok, errors := shared.ValidateInputs(newBreed); ok == false {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Correlation-Key", shared.GetID())
		w.Header().Set("API-Version", shared.Version)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors)
		return
	}

	//Check to see if the friendly id already exists in the collection
	for i := range breedsList {
		if breedsList[i].FriendlyID == newBreed.FriendlyID {
			w.WriteHeader(http.StatusBadRequest)

			theseErrors := new(shared.Errors)
			thisError := shared.ModelError{Code: "W-GLBL-STD-1005", Title: "The friendly-id " + newBreed.FriendlyID + " already exists in the breeds collection."}
			theseErrors.Errors = append(theseErrors.Errors, thisError)
			json.NewEncoder(w).Encode(&theseErrors)
			return
		}
	}

	breedsList = append(breedsList, newBreed)
	w.Header().Set("Location", "localhost:8000/breeds/"+newBreed.FriendlyID)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&newBreed)
}

func UpdateBreed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Correlation-Key", shared.GetID())
	w.Header().Set("API-Version", shared.Version)
	params := mux.Vars(r)
	var UpdatedBreed Breed
	_ = json.NewDecoder(r.Body).Decode(&UpdatedBreed)

	for index, item := range breedsList {
		if item.FriendlyID == params["friendly-id"] {
			if UpdatedBreed.Breed == "" {
				UpdatedBreed.Breed = breedsList[index].Breed
			}
			if UpdatedBreed.Group == "" {
				UpdatedBreed.Group = breedsList[index].Group
			}
			if UpdatedBreed.FriendlyID == "" {
				UpdatedBreed.FriendlyID = breedsList[index].FriendlyID
			}
			if UpdatedBreed.Origin == "" {
				UpdatedBreed.Origin = breedsList[index].Origin
			}
			UpdatedBreed.ID = breedsList[index].ID

			//Validate the updated record
			if ok, errors := shared.ValidateInputs(UpdatedBreed); ok == false {
				w.Header().Set("Content-Type", "application/json")
				w.Header().Set("Correlation-Key", shared.GetID())
				w.Header().Set("API-Version", shared.Version)
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(errors)
				return
			}

			breedsList[index] = UpdatedBreed

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(&breedsList[index])
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func DeleteBreeds(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Correlation-Key", shared.GetID())
	w.Header().Set("API-Version", shared.Version)
	params := mux.Vars(r)
	for index, item := range breedsList {
		if item.ID == params["id"] {
			breedsList = append(breedsList[:index], breedsList[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}
