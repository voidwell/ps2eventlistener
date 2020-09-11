package app

import (
	"time"

	natsClient "github.com/voidwell/natsclient/nats"
	enrichedModels "github.com/voidwell/natsclient/transport/ps2eventlistener"
)

type NatsManager struct {
	client       *natsClient.NatsClient
	shutdownChan chan int32
}

var statsStartTime = time.Now()

func NewNatsManager(endpoint string) (*NatsManager, error) {
	client, err := natsClient.NewNatsClient(endpoint)
	if err != nil {
		return nil, err
	}

	return &NatsManager{
		client:       client,
		shutdownChan: make(chan int32),
	}, nil
}

func (m *NatsManager) SendCensusEventMessage(msg *enrichedModels.CensusEvent) {
	_ = m.client.Publish("ps2census_event", msg)
}

func (m *NatsManager) SendCensusServiceStateMessage(msg interface{}) {
	_ = m.client.Publish("ps2census_state_change", msg)
}

func (m *NatsManager) Shutdown() {
	m.shutdownChan <- 1
	m.client.Shutdown()
}
