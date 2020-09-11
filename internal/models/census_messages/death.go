package models

type Death struct {
	CensusEventBase
	AttackerCharacterID string `json:"attacker_character_id"`
	AttackerFireModeID  *int   `json:"attacker_fire_mode_id,string,omitempty"`
	AttackerLoadoutID   *int   `json:"attacker_loadout_id,string,omitempty"`
	AttackerVehicleID   *int   `json:"attacker_vehicle_id,string,omitempty"`
	AttackerWeaponID    *int   `json:"attacker_weapon_id,string,omitempty"`
	CharacterID         string `json:"character_id"`
	CharacterLoadoutID  *int   `json:"character_loadout_id,string,omitempty"`
	IsHeadshot          bool   `json:"is_headshot,string"`
}
