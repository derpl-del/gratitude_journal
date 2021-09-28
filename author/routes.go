package author

import "github.com/gorilla/mux"

func Router(r *mux.Router) *mux.Router {
	r.HandleFunc("/api/v1/author", CreateAuthor).Methods("POST")
	r.HandleFunc("/api/v1/author", QueryAuthor).Methods("GET")
	r.HandleFunc("/api/v1/author", DeleteAuthor).Methods("DELETE")
	r.HandleFunc("/api/v1/author", UpdateJournal).Methods("PATCH")
	return r
}
