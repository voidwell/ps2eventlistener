package models

import (
	"time"

	"github.com/golang/protobuf/ptypes"

	enrichedModels "github.com/voidwell/natsclient/transport/ps2eventlistener"
)

type CensusEventBase struct {
	EventName string `json:"event_name"`
	WorldID   int32  `json:"world_id,string"`
	ZoneID    *int32 `json:"zone_id,string,omitempty"`
	Timestamp int64  `json:"timestamp,string"`
}

func (b *CensusEventBase) PrepareEnrichedModel() *enrichedModels.CensusEvent {
	var zoneID int32 = -1
	if b.ZoneID != nil {
		zoneID = *b.ZoneID
	}

	timestamp, _ := ptypes.TimestampProto(time.Unix(b.Timestamp, 0))

	return &enrichedModels.CensusEvent{
		EventName: b.EventName,
		WorldID:   b.WorldID,
		ZoneID:    zoneID,
		Timestamp: timestamp,
	}
}
