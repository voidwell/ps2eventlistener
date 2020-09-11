package app

type EventPublisher struct {
	eventEnricher *EventEnricher
	natsManager   *NatsManager
}

func NewEventPublisher(ps2censusEndpoint string, natsEndpoint string) (*EventPublisher, error) {
	enricher, err := NewEventEnricher(ps2censusEndpoint)
	if err != nil {
		return nil, err
	}

	natsManager, err := NewNatsManager(natsEndpoint)
	if err != nil {
		return nil, err
	}

	return &EventPublisher{
		eventEnricher: enricher,
		natsManager:   natsManager,
	}, nil
}

func (p *EventPublisher) PublishEvent(eventName string, eventObject interface{}) {
	enrichedEvent := p.eventEnricher.Enrich(eventName, eventObject)

	p.natsManager.SendCensusEventMessage(enrichedEvent)
}

func (p *EventPublisher) PublishServiceStateEvent(worldID int, isOnline bool) {
	// TODO: Update enrich models for service state change
	p.natsManager.SendCensusServiceStateMessage(isOnline)
}

func (p *EventPublisher) Dispose() {
	p.natsManager.Shutdown()
}
