package models

import(
	"time"
)
// Status is struct for dealing with seluxit status model
type Status struct {
	Message   string    `json:"message"`
	Data      string    `json:"data"`
	Timestamp time.Time `json:"timestamp"`
	Level     string    `json:"level"`
	Type      string    `json:"type"`
	Meta      struct {
		ID string `json:"id"`
	} `json:"meta"`
}