package models

// Value is struct for dealing with seluxit value model
type Value struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	Status     string `json:"status"`
	Permission string `json:"permission"`
	XML        struct {
		Xsd       string `json:"xsd"`
		Namespace string `json:"namespace"`
	} `json:"xml"`
	State []State`json:"state"`
	Meta struct {
		ID string `json:"id"`
	} `json:"meta"`
	Number Number `json:"number"`
}