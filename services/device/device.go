package device

import(
	"github.com/io-m/echo-wedge/models"
	"github.com/io-m/echo-wedge/storage/device"
)

var(
	ds = device.NewDeviceStorage()
)

type deviceService struct{}

// NewDeviceService is constructor function for making new instances of
// ServiceDevice interface -> used in controllers layer as dependency injection
func NewDeviceService() ServiceDevice {
	return &deviceService{}
}

func (*deviceService) SaveDevice(d *models.Device) (*models.Device, error) {
	err := ds.Save(d)
	if err != nil {
		return nil, err
	}
	savedDevice := d
	return savedDevice, nil
}

