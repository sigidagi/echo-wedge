package models


// Number is struct for dealing with seluxit number model
type Number struct {
	Min            float32       `json:"min,omitempty"`   // required
	Max            float32       `json:"max,omitempty"`   // required
	Step           float32       ` json:"step,omitempty"` // required
	Unit           string        `json:"unit,omitempty"`
}
