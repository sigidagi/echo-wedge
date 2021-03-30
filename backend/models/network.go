package models

type Network struct {
	Name   string   `json:"name"`
	Device []Device `json:"device"`
	Meta   Meta     `json:"meta"`
}
