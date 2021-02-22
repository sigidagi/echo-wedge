package models

import(
	"time"
)

// State is struct for dealing with seluxit state model
type State struct {
	Timestamp time.Time `json:"timestamp"`
	Data      string    `json:"data"`
	Type      string    `json:"type"`
	Meta      struct {
		ID string `json:"id"`
	} `json:"meta"`
}