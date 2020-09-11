package models

type VehicleDestroy struct {
	CensusEventBase
	AttackerCharacterID string `json:"attacker_character_id"`
	AttackerLoadoutID   *int   `json:"attacker_loadout_id,string,omitempty"`
	AttackerVehicleID   *int   `json:"attacker_vehicle_id,string,omitempty"`
	AttackerWeaponID    *int   `json:"attacker_weapon_id,string,omitempty"`
	CharacterID         string `json:"character_id"`
	FacilityID          *int   `json:"facility_id,string,omitempty"`
	FactionID           *int   `json:"faction_id,string,omitempty"`
	VehicleID           int    `json:"vehicle_id,string"`
}
