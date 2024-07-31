package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade to WebSocket:", err)
		return
	}
	defer conn.Close()

	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			return
		}

		log.Printf("Received message: %s\n", msg)

		if err = conn.WriteMessage(msgType, msg); err != nil {
			log.Println("Error writing message:", err)
			return
		}
	}
}
