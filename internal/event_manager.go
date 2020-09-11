package app

import (
	"encoding/json"
	censusModels "ps2eventlistener/internal/models/census_messages"
	"regexp"
	"strconv"
)

var regServer = regexp.MustCompile(`EventServerEndpoint_(?P<WorldName>.*)_(?P<WorldID>.*)`)

type EventManager struct {
	censusClient   *CensusClient
	eventPublisher *EventPublisher
	exitChan       chan struct{}
}

func NewEventManager(censusKey string, censusNamespace string, ps2censusEndpoint string, natsEndpoint string) (*EventManager, error) {
	eventPublisher, err := NewEventPublisher(ps2censusEndpoint, natsEndpoint)
	if err != nil {
		return nil, err
	}

	return &EventManager{
		censusClient:   NewCensusClient(censusKey, censusNamespace),
		eventPublisher: eventPublisher,
		exitChan:       make(chan struct{}),
	}, nil
}

func (m *EventManager) Connect(eventNames []string, characterIDs []string, worldIDs []string) {
	go m.censusClient.Connect()

	for {
		select {
		case <-m.exitChan:
			return
		case <-m.censusClient.OnConnected:
			go m.censusClient.Subscribe(eventNames, characterIDs, worldIDs)
		case msg := <-m.censusClient.MessageChan:
			go m.processMessage(msg)
		}
	}
}

func (m *EventManager) processMessage(msg []byte) {
	var jsonMap map[string]interface{}
	json.Unmarshal(msg, &jsonMap)

	if jsonMap["type"] != nil && jsonMap["type"] == "serviceStateChanged" {
		m.processServiceStateChanged(msg)
	} else {
		m.processServiceEvent(jsonMap)
	}
}

func (m *EventManager) processServiceStateChanged(msg []byte) {
	var serviceStateMessage censusModels.CensusServiceState
	err := json.Unmarshal(msg, &serviceStateMessage)
	if err != nil {
		return
	}

	matchList := regServer.FindStringSubmatch(serviceStateMessage.Detail)

	if worldID, err := strconv.ParseInt(matchList[2], 10, 64); err == nil {
		m.eventPublisher.PublishServiceStateEvent(int(worldID), serviceStateMessage.IsOnline)
	}
}

func (m *EventManager) processServiceEvent(jsonMap map[string]interface{}) {
	jPayload := jsonMap["payload"]
	if jPayload == nil {
		return
	}

	payloadBytes, _ := json.Marshal(jPayload)

	var payloadMessage censusModels.CensusEventBase
	err := json.Unmarshal(payloadBytes, &payloadMessage)

	if err != nil || payloadMessage.EventName == "" {
		return
	}

	switch payloadMessage.EventName {
	case "AchievementEarned":
		var eventMsg censusModels.AchievementEarned
		json.Unmarshal(payloadBytes, &eventMsg)
		m.eventPublisher.PublishEvent(eventMsg.EventName, eventMsg)
	case "BattleRankUp":
		var eventMsg censusModels.BattleRankUp
		json.Unmarshal(payloadBytes, &eventMsg)
		m.eventPublisher.PublishEvent(eventMsg.EventName, eventMsg)
	case "ContinentLock":
		var eventMsg censusModels.ContinentLock
		json.Unmarshal(payloadBytes, &eventMsg)
		m.eventPublisher.PublishEvent(eventMsg.EventName, eventMsg)
	case "ContinentUnlock":
		var eventMsg censusModels.ContinentUnlock
		json.Unmarshal(payloadBytes, &eventMsg)
		m.eventPublisher.PublishEvent(eventMsg.EventName, eventMsg)
	case "Death":
		var eventMsg censusModels.Death
		json.Unmarshal(payloadBytes, &eventMsg)
		m.eventPublisher.PublishEvent(eventMsg.EventName, eventMsg)
	case "FacilityControl":
		var eventMsg censusModels.FacilityControl
		json.Unmarshal(payloadBytes, &eventMsg)
		m.eventPublisher.PublishEvent(eventMsg.EventName, eventMsg)
	case "GainExperience":
		var eventMsg censusModels.GainExperience
		json.Unmarshal(payloadBytes, &eventMsg)
		m.eventPublisher.PublishEvent(eventMsg.EventName, eventMsg)
	case "MetagameEvent":
		var eventMsg censusModels.MetagameEvent
		json.Unmarshal(payloadBytes, &eventMsg)
		m.eventPublisher.PublishEvent(eventMsg.EventName, eventMsg)
	case "PlayerFacilityCapture":
		var eventMsg censusModels.PlayerFacilityCapture
		json.Unmarshal(payloadBytes, &eventMsg)
		m.eventPublisher.PublishEvent(eventMsg.EventName, eventMsg)
	case "PlayerFacilityDefend":
		var eventMsg censusModels.PlayerFacilityDefend
		json.Unmarshal(payloadBytes, &eventMsg)
		m.eventPublisher.PublishEvent(eventMsg.EventName, eventMsg)
	case "PlayerLogin":
		var eventMsg censusModels.PlayerLogin
		json.Unmarshal(payloadBytes, &eventMsg)
		m.eventPublisher.PublishEvent(eventMsg.EventName, eventMsg)
	case "PlayerLogout":
		var eventMsg censusModels.PlayerLogout
		json.Unmarshal(payloadBytes, &eventMsg)
		m.eventPublisher.PublishEvent(eventMsg.EventName, eventMsg)
	case "VehicleDestroy":
		var eventMsg censusModels.VehicleDestroy
		json.Unmarshal(payloadBytes, &eventMsg)
		m.eventPublisher.PublishEvent(eventMsg.EventName, eventMsg)
	}
}

func (m *EventManager) Dispose() {
	m.censusClient.Dispose()
	m.eventPublisher.Dispose()
}
