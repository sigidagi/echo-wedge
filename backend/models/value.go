package models

// Value is struct for dealing with seluxit value model
type Value struct {
	Name       string  `json:"name"`
	Type       string  `json:"type,omitempty"`
	Status     string  `json:"status,omitempty"`
	Permission string  `json:"permission"`
	State      []State `json:"state"`
	Meta       Meta    `json:"meta"`
	Number     Number  `json:"number"`
}
