package main

import (
	"log"
	"net/http"
	"time"

	"github.com/boltdb/bolt"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func boltDB(db *bolt.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Upgrade the handler for websockets
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		for {
			_, p, err := conn.ReadMessage()
			if err != nil {
				return
			}

			//Add the DB operation here
			log.Println("Got some data:", string(p))
			time.Sleep(3 * time.Second)

			msg := string(p) + " task is done"
			log.Println(msg)
			_ = conn.WriteMessage(websocket.TextMessage, []byte(msg))
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

	http.HandleFunc("/db", boltDB(db))
	http.Handle("/", http.FileServer(http.Dir("./public")))
	err = http.ListenAndServe(":3001", nil)
	if err != nil {
		panic("Error: " + err.Error())
	}

}
