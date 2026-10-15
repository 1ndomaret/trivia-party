package socket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v5"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func CreateSession(c *echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return err
	}
	defer ws.Close()

	fmt.Println("Client connected!")

	for {
		messageType, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("Client disconnected or read error:", err)
			break
		}

		fmt.Printf("Received from client: %s\n", message)

		response := []byte("Hello back from Gorilla WebSocket!")
		err = ws.WriteMessage(messageType, response)
		if err != nil {
			log.Println("Write error:", err)
			break
		}
	}

	return nil
}
