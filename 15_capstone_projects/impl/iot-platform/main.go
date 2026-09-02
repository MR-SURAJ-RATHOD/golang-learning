package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	fmt.Println("New Client Connected")

	for {
		// Echo for now
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}
		log.Printf("Received: %s", message)

		// Broadcast to all (implementation omitted for brevity)
		err = conn.WriteMessage(mt, message)
		if err != nil {
			log.Println("Write error:", err)
			break
		}
	}
}

func main() {
	fmt.Println("=== Capstone Project: IoT Platform ===")

	http.HandleFunc("/ws", handleWebSocket)

	fmt.Println("WebSocket Server listening on :8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
