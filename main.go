package main

import (
	"factorial/handler"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.POST("/calculate", handler.FactRequestHandler)
	http.ListenAndServe(":8989", router)
}
