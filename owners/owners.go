package owners

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"petstore/shared"
)

type Owner struct {
	ID          string `json:"id"`
	FirstName   string `json:"firstName" validate:"required"`
	LastName    string `json:"lastName" validate:"required"`
	FullName    string `json:"fullName", omitempty`
	Email       string `json:"email", validate:"email,required"`
	MobilePhone string `json:"mobilePhone", validate:"numeric,len=10,required"`
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
