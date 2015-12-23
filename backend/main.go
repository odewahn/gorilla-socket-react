package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func pulsar(w http.ResponseWriter, r *http.Request) {
	// Upgrade the handler for websockets
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	// Set up a ticker that will push something in periodically
	ft := time.NewTicker(1000 * time.Millisecond)
	defer func() {
		ft.Stop()
	}()
	for {
		select {
		case <-ft.C:
			t := fmt.Sprintf(time.Now().Format(time.RFC3339))
			msg := "Tickles sent at " + t
			log.Println(msg)
			err := conn.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				log.Println(err)
			}
		}
	}

}

func main() {
	http.HandleFunc("/pulsar", pulsar)
	http.Handle("/", http.FileServer(http.Dir("./public")))
	err := http.ListenAndServe(":3001", nil)
	if err != nil {
		panic("Error: " + err.Error())
	}

}
