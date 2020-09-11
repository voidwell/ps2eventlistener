package models

type ContinentLock struct {
	CensusEventBase
	TriggeringFaction int     `json:"triggering_faction,string"`
	MetagameEventID   int     `json:"metagame_event_id,string"`
	VsPopulation      float32 `json:"vs_population,string"`
	NcPopulation      float32 `json:"nc_population,string"`
	TrPopulation      float32 `json:"tr_population,string"`
}
