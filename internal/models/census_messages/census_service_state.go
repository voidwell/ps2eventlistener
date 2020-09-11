package models

type CensusServiceState struct {
	Detail      string `json:"detail"`
	IsOnline    bool   `json:"online,string"`
	Service     string `json:"service"`
	MessageType string `json:"type"`
}
