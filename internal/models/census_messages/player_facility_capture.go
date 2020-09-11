package models

type PlayerFacilityCapture struct {
	CensusEventBase
	CharacterID string `json:"character_id"`
	FacilityID  int    `json:"facility_id,string"`
	OutfitID    string `json:"outfit_id"`
}
