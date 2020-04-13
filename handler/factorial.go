package handler

import (
	"encoding/json"
	"factorial/compute"
	"factorial/validation"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//FactRequestHandler :
func FactRequestHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	request, err := ioutil.ReadAll(r.Body) //Fetching request body
	if err != nil {
		log.Print(err)
	}
	a, b, validInput := validation.InputValidator(request) //Parsing and validating request

	if validInput == true {
		fact := make(chan int64, 2)   //Buffered channel to get output from go routines
		go compute.Factorial(fact, a) //Calling goroutine
		factA := <-fact
		go compute.Factorial(fact, b)
		factB := <-fact
		out := factA * factB

		if factA < 0 || factB < 0 || out <= 0 { //Validating if factorial and output is not too long
			http.Error(w, `{"Error":"Please enter small numbers"}`, 400)
		} else { //Setting output
			w.Header().Set("Content-Type", "application/json")
			Output := output{
				Output: out,
			}
			json.NewEncoder(w).Encode(Output)

		}
	} else {
		http.Error(w, `{"Error":"Incorrect Input"}`, 400)
	}

}

// Success Response structure
type output struct {
	Output int64 `json:"output"`
}
