package search

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {

}

// Object : Launch Search Object
func Object(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	// Save var GET
	vars := mux.Vars(r)

	domain := vars["domain"]
	object := vars["object"]

	result := Result{
		Domain: domain,
		Object: object,
	}

	json.NewEncoder(w).Encode(result)
}
