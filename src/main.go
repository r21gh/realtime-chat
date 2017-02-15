package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)  // connected clients
var broadcast = make(chan Message) // broadcast channel

// configure the upgrader
var upgrader = websocket.Upgrader{}

// define our message object
type Message struct {
	Email string `json:"email"`
	Username string `json:"username"`
	Message string `json:"message"`
}

func main()  {
	// create a simple file server
	fs := http.FileServer(http.Dir("../public"))
	http.Handle("/", fs)

	// configure websocket route
	http.HandleFunc("/ws", handleConnections)
}

