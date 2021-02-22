package models


// Device is struct for dealing with seluxit device model
type Device struct {
	Name          string `json:"name"`
	Manufacturer  string `json:"manufacturer"`
	Product       string `json:"product"`
	Serial        string `json:"serial"`
	Description   string `json:"description"`
	Protocol      string `json:"protocol"`
	Communication string `json:"communication"`
	Value         []Value `json:"value"`
	Status []Status `json:"status"`
	Meta struct {
		ID string `json:"id"`
	} `json:"meta"`
}

// AllDevices is Fake DB storage -> slice of Device models
type AllDevices map[string]*Device

// NewDeviceStorage is constructor func for new instances of devices storage
func NewDeviceStorage() map[string]*Device {
	devices := make(map[string]*Device)
	return devices
}