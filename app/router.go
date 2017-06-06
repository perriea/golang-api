package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/perriea/Golang-API/app/logs"
	"github.com/perriea/Golang-API/routes/index"
	"github.com/perriea/Golang-API/routes/search"
)

// API is the server mux for handling API calls
var API = mux.NewRouter().StrictSlash(true)

func Router() {

	API.Host("localhost")

	API.HandleFunc("/", index.Run)
	API.HandleFunc("/search/{domain}/{object}", search.Object)

	//http.Handle("/", middlewareOne(API))
	http.Handle("/", API)

	loggingHandler := logs.NewApacheLoggingHandler(API, os.Stderr)
	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", 8080),
		Handler: loggingHandler,
	}

	log.Fatal(server.ListenAndServe())
}
