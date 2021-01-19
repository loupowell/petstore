package owners

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

type Owner struct {
	ID          string `json:"id"`
	FirstName   string `json:"firstName" validate:"required"`
	LastName    string `json:"lastName" validate:"required"`
	FullName    string `json:"fullName, omitempty"`
	Email       string `json:"email" validate:"omitempty,email"`
	MobilePhone string `json:"mobilePhone" validate:"required,numeric,max=10,min=10"`
}

var ownersList []Owner

//Retrieve the entire owners collection
func GetOwners(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Correlation-Key", shared.GetID())
	w.Header().Set("API-Version", shared.Version)
	if len(ownersList) == 0 {
		loadOwners, _ := ioutil.ReadFile("owners/ownerslist.json")
		err := json.Unmarshal([]byte(loadOwners), &ownersList)
		if err != nil {
			fmt.Println(err)
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ownersList)
}

//GetOwners Retrieve a specific Owner based on id
func GetOwner(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Correlation-Key", shared.GetID())
	w.Header().Set("API-Version", shared.Version)
	params := mux.Vars(r)
	for index, item := range ownersList {
		if item.ID == params["id"] {
			requestedRecord := ownersList[index]
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(requestedRecord)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

//Add a new Owner to the Owners collection
func PostOwners(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Correlation-Key", shared.GetID())
	w.Header().Set("API-Version", shared.Version)
	var newOwner Owner
	_ = json.NewDecoder(r.Body).Decode(&newOwner)
	newOwner.ID = strconv.Itoa(rand.Intn(100000000))
	newOwner.FullName = newOwner.FirstName + " " + newOwner.LastName

	//Validate the posted record
	if ok, errors := shared.ValidateInputs(newOwner); ok == false {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Correlation-Key", shared.GetID())
		w.Header().Set("API-Version", shared.Version)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors)
		return
	}

	ownersList = append(ownersList, newOwner)
	w.Header().Set("Location", "/owners/"+newOwner.ID)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&newOwner)
}
