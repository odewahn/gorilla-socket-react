package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/boltdb/bolt"
	"github.com/gorilla/mux"
)

// Message holds a decoded message passed from the frontend
type Message struct {
	Bucket string
	Key    string
	Value  string
}

func writeMsg(db *bolt.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		fmt.Println(string(body))
		// Create a message struct
		var msg Message
		json.Unmarshal(body, &msg)
		fmt.Println(string(msg.Key), string(msg.Value))
		// Save to the bolt DB
		err := db.Update(func(tx *bolt.Tx) error {
			log.Println("Writing record")
			bucket, err := tx.CreateBucketIfNotExists([]byte(msg.Bucket))
			if err != nil {
				return err
			}
			err = bucket.Put([]byte(msg.Key), []byte(msg.Value))
			if err != nil {
				return err
			}
			return nil
		})

		if err != nil {
			log.Fatal(err)
		}
	}
}

//Read a key value in the form /db/:bucket/:key
func readMsg(db *bolt.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		log.Println(vars["bucket"], vars["key"])
		bucketName := []byte(vars["bucket"])
		keyName := []byte(vars["key"])
		// retrieve the data
		err := db.View(func(tx *bolt.Tx) error {
			bucket := tx.Bucket(bucketName)
			if bucket == nil {
				return fmt.Errorf("Bucket %q not found!", bucketName)
			}

			val := bucket.Get(keyName)

			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, string(val))

			return nil
		})

		if err != nil {
			log.Fatal(err)
		}
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

	r.Handle("/db/{bucket}/{key}", readMsg(db)).Methods("GET")
	r.Handle("/db", writeMsg(db)).Methods("POST")
	http.Handle("/", http.FileServer(http.Dir("./public")))

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public")))

	err = http.ListenAndServe(":3001", r)
	if err != nil {
		panic("Error: " + err.Error())
	}

}
