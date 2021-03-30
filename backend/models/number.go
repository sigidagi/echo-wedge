package models

// Number is struct for dealing with seluxit number model
type Number struct {
	Min  float32 `json:"min"`  // required
	Max  float32 `json:"max"`  // required
	Step float32 `json:"step"` // required
	Unit string  `json:"unit,omitempty"`
}
