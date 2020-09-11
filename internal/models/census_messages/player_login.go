package models

type PlayerLogin struct {
	CensusEventBase
	CharacterID string `json:"character_id"`
}
