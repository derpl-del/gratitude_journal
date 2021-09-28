package main

import (
	"gratitude_journal/derpl-del/author"
	"gratitude_journal/derpl-del/infrastructure"
	"gratitude_journal/derpl-del/journal"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var env = infrastructure.Environment{}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	r := mux.NewRouter()
	r = journal.Router(r)
	r = author.Router(r)
	log.Println("Server ready at 9000")
	log.Fatal(http.ListenAndServe(":9001", r))
}

func init() {
	env.LoadConfig()
	env.InitDatabase()
	infrastructure.PGdatabase.Table.AutoMigrate(&author.Author{}, &journal.Journal{}, &journal.Point{})
}
