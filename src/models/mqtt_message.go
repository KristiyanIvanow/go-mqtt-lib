package models

import "github.com/KristiyanIvanow/go-mqtt-lib/src/types"

// MqttMessage represents an MQTT message received or to be published.
type MqttMessage struct {
	Topic    string     `json:"topic"`
	Payload  []byte     `json:"payload"`
	QoS      types.EQoS `json:"qos"`
	Retained bool       `json:"retained"`
}
