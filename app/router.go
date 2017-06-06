package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/perriea/Golang-API/routes/index"
)

// API is the server mux for handling API calls
var API = mux.NewRouter().StrictSlash(true)

func Router() {

	API.Host("localhost")

	API.HandleFunc("/", index.Run)

	http.Handle("/", API)
	http.ListenAndServe(":8080", nil)
}
