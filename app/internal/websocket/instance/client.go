package ws

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

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
