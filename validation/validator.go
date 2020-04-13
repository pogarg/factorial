package validation

import (
	"encoding/json"
	"log"
)

//Method parsing input json and validating input
func InputValidator(request []byte) (int, int, bool) {
	var validInput bool
	var a float64
	var b float64
	var parserequest map[string]interface{}
	err := json.Unmarshal(request, &parserequest)
	if err != nil {
		log.Print(err)
	}
	if parserequest["a"] != nil && parserequest["b"] != nil { //Checking if input exist
		a = parserequest["a"].(float64)
		b = parserequest["b"].(float64)
		if a > 0 && b > 0 { //Checking if input numbers are non negative
			validInput = true
		}
	} else {
		validInput = false
	}
	return int(a), int(b), validInput
}
