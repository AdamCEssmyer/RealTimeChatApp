package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

var upgrader = websocket.Upgrader{}

type Message struct {
	Email string `json:"email"`
	Username string `json:"username`
	Message string `json:"message`
}
var upgrader = websocket.Upgrader{}

func main() {
	//Create a simple file server
	fs :=http.FileServer(http.Dir("../public"))
	http.Handle("/", fs)
	//Configure websocket route
	http.HandleFunc("/ws", handleConnections)
}
