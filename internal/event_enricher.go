package app

import (
	censusModels "ps2eventlistener/internal/models/census_messages"

	enrichedModels "github.com/voidwell/natsclient/transport/ps2eventlistener"
)

type EventEnricher struct {
	censusClient *PS2CensusClient
}

func NewEventEnricher(ps2censusEndpoint string) (*EventEnricher, error) {
	censusClient, err := NewPS2CensusClient(ps2censusEndpoint)
	if err != nil {
		return nil, err
	}

	return &EventEnricher{
		censusClient: censusClient,
	}, nil
}

func (e *EventEnricher) Enrich(eventName string, eventData interface{}) *enrichedModels.CensusEvent {
	censusEvent := (eventData).(*censusModels.CensusEventBase)
	enrichedModel := censusEvent.PrepareEnrichedModel()

	switch event := eventName; event {
	/*
		case "Death":
			eventModel := (eventData).(*censusModels.Death)

			if eventModel.AttackerCharacterID != "" {
				resultData, _ := e.censusClient.GetCharacter(eventModel.AttackerCharacterID)
			}

			if eventModel.CharacterID != "" {
				resultData, _ := e.censusClient.GetCharacter(eventModel.CharacterID)
			}


		case "FacilityControl":
			eventModel := (eventData).(*censusModels.FacilityControl)
	*/
	default:
		return enrichedModel
	}
}
