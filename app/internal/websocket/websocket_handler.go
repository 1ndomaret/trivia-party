package ws

import (
	"app/app/internal/model/entity"
	"log"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/labstack/echo/v5"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type ClientRole string

const (
	RoleHost   ClientRole = "HOST"
	RolePlayer ClientRole = "PLAYER"
)

type Client struct {
	ID       uuid.UUID
	Conn     *websocket.Conn
	Send     chan []byte
	Role     ClientRole
	Nickname string
	Score    int
}

type WsRoom struct {
	ID               uuid.UUID
	RoomCode         string
	HostID           uuid.UUID
	QuestionCategory string
	MaxPlayers       int
	QuestionCount    int

	Clients    *sync.Map
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}

type WsRoomManager struct {
	Rooms *sync.Map
}

func NewWsRoom(room *entity.Room) *WsRoom {
	return &WsRoom{
		ID:               room.ID,
		RoomCode:         room.RoomCode,
		HostID:           room.UserID,
		QuestionCategory: room.QuestionCategory,
		MaxPlayers:       room.MaxPlayers,
		QuestionCount:    room.QuestionCount,

		Clients:    &sync.Map{},
		Broadcast:  make(chan []byte, 256),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

func StartRoom(c *echo.Context, room *entity.Room) {
	wsRoom := NewWsRoom(room)
	go wsRoom.Run()
}

func (wr *WsRoom) Run() {
	for {
		select {
		case client := <-wr.Register:
			wr.Clients.Store(client.Nickname, client)
			wr.Broadcast <- []byte(client.Nickname + " joined! 🎉")
		case client := <-wr.Unregister:
			wr.Clients.Delete(client.Nickname)
			wr.Broadcast <- []byte(client.Nickname + " left. 👋")
		case message := <-wr.Broadcast:
			wr.Clients.Range(func(key, value interface{}) bool {
				client := value.(*Client)
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					wr.Clients.Delete(key)
				}
				return true
			})
		}
	}
}

func (wr *WsRoom) HandleConnection(c *echo.Context) {
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}

	nickname := c.QueryParam("nickname")
	if nickname == "" {
		nickname = "Anonymous"
	}

	client := &Client{Conn: conn, Send: make(chan []byte, 256), Nickname: nickname}
	wr.Register <- client

	go client.write()
	go client.read(wr)
}

func (c *Client) write() {
	defer c.Conn.Close()
	for message := range c.Send {
		if err := c.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Println("Write error:", err)
			return
		}
	}
}

// read processes incoming messages
func (c *Client) read(wr *WsRoom) {
	defer func() {
		wr.Unregister <- c
		c.Conn.Close()
	}()
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			return
		}
		wr.Broadcast <- []byte(c.Nickname + ": " + string(message))
	}
}
