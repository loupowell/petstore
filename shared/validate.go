package shared

import (
	"fmt"
	"reflect"
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

type ModelError struct {
	Code   string `json:"code"`
	Title  string `json:"title"`
	Detail string `json:"detail,omitempty"`
	//Links  []Link `json:"links,omitempty"`
}

type Errors struct {
	Errors []ModelError `json:"errors"`
}

const Version = "01.00.12"

func ValidateInputs(dataSet interface{}) (state bool, errors interface{}) {

	validate := validator.New()
	err := validate.Struct(dataSet)

	if err != nil {

		//Validation syntax is invalid
		if err, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println("Validation: Error validating struct.")
			panic(err)
		}

		//Validation errors occurred
		theseErrors := new(Errors)

		//Use reflector to reverse engineer struct
		reflected := reflect.ValueOf(dataSet)
		for _, err := range err.(validator.ValidationErrors) {

			// Attempt to find field by name and get json tag name
			field, _ := reflected.Type().FieldByName(err.StructField())
			var name string
			//If json tag doesn't exist, use lower case of name
			if name = field.Tag.Get("json"); name == "" {
				name = strings.ToLower(err.StructField())
			}

			switch err.Tag() {
			case "required":
				thisError := ModelError{Code: "W-GLBL-STD-1001", Title: "The " + name + " field is required."}
				theseErrors.Errors = append(theseErrors.Errors, thisError)
				break
			case "email":
				thisError := ModelError{Code: "W-GLBL-STD-1002", Title: "The " + name + " field should be a valid email address."}
				theseErrors.Errors = append(theseErrors.Errors, thisError)
				break
			case "eqfield":
				thisError := ModelError{Code: "W-GLBL-STD-1003", Title: "The " + name + " field should be equal to the " + err.Param() + " field."}
				theseErrors.Errors = append(theseErrors.Errors, thisError)
				break
			case "max":
				thisError := ModelError{Code: "W-GLBL-STD-1004", Title: "The " + name + " field is too long."}
				theseErrors.Errors = append(theseErrors.Errors, thisError)
				break
			case "oneof":
				thisError := ModelError{Code: "W-GLBL-STD-1005", Title: "The " + name + " field must containt a valid enumeration."}
				theseErrors.Errors = append(theseErrors.Errors, thisError)
				break
			default:
				thisError := ModelError{Code: "W-GLBL-STD-1000", Title: "General Error. The " + name + " field is invalid."}
				theseErrors.Errors = append(theseErrors.Errors, thisError)
				break
			}
		}

		return false, theseErrors
	}
	return true, nil
}
