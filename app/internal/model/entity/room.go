package entity

import (
	"time"

	"github.com/google/uuid"
)

type RoomStatus string

const (
	RoomStandby  RoomStatus = "LOBBY"
	RoomActive   RoomStatus = "IN_GAME"
	RoomFinished RoomStatus = "FINISHED"
)

type Room struct {
	ID               uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	RoomCode         string     `gorm:"type:varchar(4);not null;index:idx_active_room_code,unique,where:status != 'FINISHED'" json:"room_code"`
	UserID           uuid.UUID  `gorm:"type:uuid;uniqueIndex;not null" json:"user_id"`
	QuestionCategory string     `gorm:"type:varchar(50);not null" json:"question_category"`
	QuestionCount    int        `gorm:"default:10" json:"question_count"`
	MaxPlayers       int        `gorm:"default:8" json:"max_players"`
	Status           RoomStatus `gorm:"type:varchar(20);default:'LOBBY'" json:"status"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`

	User *User `gorm:"foreignKey:UserID;" json:"-"`
}
