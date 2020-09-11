package models

type FacilityControl struct {
	CensusEventBase
	FacilityID   int    `json:"facility_id,string"`
	NewFactionID int    `json:"new_faction_id,string"`
	OldFactionID int    `json:"old_faction_id,string"`
	DurationHeld int    `json:"duration_held,string"`
	OutfitID     string `json:"outfit_id"`
}
