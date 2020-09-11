package models

type CensusSubscription struct {
	Service    string   `json:"service,omitempty"`
	Action     string   `json:"action,omitempty"`
	Characters []string `json:"characters,omitempty"`
	Worlds     []string `json:"worlds,omitempty"`
	EventNames []string `json:"eventNames,omitempty"`
}

func NewCensusSubscription(eventNames []string, characterIDs []string, worldIDs []string) *CensusSubscription {
	return &CensusSubscription{
		Service:    "event",
		Action:     "subscribe",
		Characters: characterIDs,
		Worlds:     worldIDs,
		EventNames: eventNames,
	}
}
