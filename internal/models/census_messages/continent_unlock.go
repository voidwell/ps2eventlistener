package models

type ContinentUnlock struct {
	CensusEventBase
	TriggeringFaction int `json:"triggering_faction,string"`
	MetagameEventID   int `json:"metagame_event_id,string"`
}
