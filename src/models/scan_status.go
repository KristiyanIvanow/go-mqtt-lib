package models

import "sdk-go/src/types"

// ScanStatus represents the current state of a topic scanning operation.
type ScanStatus struct {
	Status         types.ScanStatusEnum `json:"status"`
	ScanEndAt      string               `json:"scanEndAt"`
	ExploreTopics  []string             `json:"exploreTopics,omitempty"`
}
