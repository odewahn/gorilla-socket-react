package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/boltdb/bolt"
	"github.com/gorilla/mux"
)

func boltDB(db *bolt.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		fmt.Println(string(body))
	}
}

func main() {

	// Set up the database
	db, err := bolt.Open("/Users/odewahn/.launchbot/config.db", 0644, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := mux.NewRouter()

	r.Handle("/db", boltDB(db)).Methods("POST")
	http.Handle("/", http.FileServer(http.Dir("./public")))

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public")))

	err = http.ListenAndServe(":3001", r)
	if err != nil {
		panic("Error: " + err.Error())
	}

}
