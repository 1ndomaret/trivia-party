package dto

type RoomUpdateRequest struct {
	QuestionCategory string `json:"question_category"`
	QuestionCount    int    `json:"question_count"`
	MaxPlayers       int    `json:"max_players"`
}
