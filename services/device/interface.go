package device

import (
	"github.com/io-m/echo-wedge/models"
)

// ServiceDevice is an interface that defines all methods on service layer
type ServiceDevice interface {
	SaveDevice(d *models.Device) (*models.Device, error)
}