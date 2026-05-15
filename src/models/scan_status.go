package models

import "github.com/KristiyanIvanow/go-mqtt-lib/src/types"

// ScanStatus represents the current state of a topic scanning operation.
type ScanStatus struct {
	Status         types.ScanStatusEnum `json:"status"`
	ScanEndAt      string               `json:"scanEndAt"`
	ExploreTopics  []string             `json:"exploreTopics,omitempty"`
}
