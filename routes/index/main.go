package index

import (
	"encoding/json"
	"net/http"
)

func Run(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	methods := Methods{
		Method{Name: "/", Desc: "Index"},
		Method{Name: "/search/{domain}/{object}", Desc: "Search an object"},
	}

	index := IndexPage{
		Count:   len(methods),
		Methods: methods,
	}

	json.NewEncoder(w).Encode(index)
}
