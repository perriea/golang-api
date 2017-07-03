package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// API is the server mux for handling API calls
var API = mux.NewRouter().StrictSlash(true)

func redirect(w http.ResponseWriter, req *http.Request) {
	// remove/add not default ports from req.Host
	target := "https://" + req.Host + req.URL.Path
	if len(req.URL.RawQuery) > 0 {
		target += "?" + req.URL.RawQuery
	}
	log.Printf("redirect to: %s", target)
	http.Redirect(w, req, target, http.StatusTemporaryRedirect)
}

func index(w http.ResponseWriter, req *http.Request) {
	// all calls to unknown url paths should return 404
	if req.URL.Path != "/" {
		log.Printf("404: %s", req.URL.String())
		http.NotFound(w, req)
	} else {
		w.Write([]byte("Hello !!!\n"))
	}
}

func Router() {

	go http.ListenAndServe(":80", http.HandlerFunc(redirect))

	//http.HandleFunc("/", index.Run)

	API.Host("localhost")
	http.Handle("/", API)

	err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
