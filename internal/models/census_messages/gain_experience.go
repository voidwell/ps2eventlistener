package models

type GainExperience struct {
	CensusEventBase
	CharacterID  string `json:"character_id"`
	ExperienceID int    `json:"experience_id,string"`
	Amount       int    `json:"Amount,string"`
	LoadoutID    *int   `json:"loadout_id,string,omitempty"`
	OtherID      string `json:"other_id"`
}
