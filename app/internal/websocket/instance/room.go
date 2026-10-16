package ws

import (
	"sync"

	"github.com/google/uuid"
)

type Room struct {
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
