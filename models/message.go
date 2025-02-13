package models

import "github.com/google/uuid"

// JSON prefers camelCase
type Message struct {
	StreamId    uuid.UUID   `json:"streamId"`
	MessageType string      `json:"messageType"`
	DataPayload interface{} `json:"dataPayload"`
}
