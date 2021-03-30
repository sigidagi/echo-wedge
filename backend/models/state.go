package models

import (
	"time"
)

// StateWg is struct for data state
type StateWg struct {
	Data interface{} `json:"data"`
	Meta Meta        `json:"meta"`
}

// State is struct for dealing with seluxit state model
type State struct {
	Timestamp time.Time `json:"timestamp"`
	Data      string    `json:"data"`
	Type      string    `json:"type"`
	Meta      Meta      `json:"meta"`
}
