package journal

import (
	"github.com/gorilla/mux"
)

func Router(r *mux.Router) *mux.Router {
	r.HandleFunc("/api/v1/journal", CreateJournal).Methods("POST")
	r.HandleFunc("/api/v1/journal", QueryJournal).Methods("GET")
	r.HandleFunc("/api/v1/journal", UpdateJournal).Methods("PATCH")
	r.HandleFunc("/api/v1/journal", DeleteJournal).Methods("Delete")
	return r
}
