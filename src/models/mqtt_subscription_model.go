package models

import "github.com/KristiyanIvanow/go-mqtt-lib/src/types"

// MqttSubscriptionModel represents an MQTT topic subscription.
type MqttSubscriptionModel struct {
	Topic string     `json:"topic"`
	QoS   types.EQoS `json:"qos"`
}

// NewMqttSubscriptionModel creates a new MqttSubscriptionModel.
func NewMqttSubscriptionModel(topic string, qos types.EQoS) *MqttSubscriptionModel {
	return &MqttSubscriptionModel{Topic: topic, QoS: qos}
}
