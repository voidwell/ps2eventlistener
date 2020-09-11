package models

type MetagameEvent struct {
	CensusEventBase
	InstanceID         int     `json:"instance_id,string"`
	MetagameEventID    int     `json:"metagame_event_id,string"`
	MetagameEventState string  `json:"metagame_event_state"`
	FactionVS          float32 `json:"faction_vs,string"`
	FactionNC          float32 `json:"faction_nc,string"`
	FactionTR          float32 `json:"faction_tr,string"`
	ExperienceBonus    float32 `json:"experience_bonus,string"`
}
