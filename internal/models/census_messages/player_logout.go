package models

type PlayerLogout struct {
	CensusEventBase
	CharacterID string `json:"character_id"`
}
