package models

import "sdk-go/src/types"

// MqttState represents the current MQTT connection state.
type MqttState struct {
	ConnectionState types.ConnectionState `json:"connectionState"`
	Message         string                `json:"message,omitempty"`
}

// NewMqttState creates a new MqttState defaulting to Disconnected.
func NewMqttState() *MqttState {
	return &MqttState{ConnectionState: types.Disconnected}
}
