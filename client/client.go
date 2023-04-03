package main

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

func main() {
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for {
		var message string
		fmt.Print("Enter message: ")
		fmt.Scanln(&message)

		if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
			log.Println(err)
			break
		}

		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}

		fmt.Printf("Received message: %s\n", msg)
	}
}
