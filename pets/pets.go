package pets

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"petstore/shared"
)

type Pet struct {
	ID         string `json:"id"`
	FriendlyID string `json:"friendly-id" validate:"required"`
	PetName    string `json:"petName" validate:"required"`
	Breed      string `json:"breed" validate:"required"`
	Owner      string `json:"owner,omitempty"`
}

var petsList []Pet

//Retrieve the entire pets collection
func GetPets(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Correlation-Key", shared.GetID())
	w.Header().Set("API-Version", shared.Version)
	if len(petsList) == 0 {
		loadPets, _ := ioutil.ReadFile("pets/petslist.json")
		err := json.Unmarshal([]byte(loadPets), &petsList)
		if err != nil {
			fmt.Println(err)
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(petsList)
}
