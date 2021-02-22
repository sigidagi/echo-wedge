package device

import (
	"fmt"
	
	"github.com/io-m/echo-wedge/models"
)

var(
	devices = []*models.Device{}
	allDevices = models.NewDeviceStorage()
)

type deviceStorage struct{}

// NewDeviceStorage is contructor function for making new instances
// of StorageDevice interface -> used in service layer as dependency injection
func NewDeviceStorage() StorageDevice{
	return &deviceStorage{}
}

func (*deviceStorage) Save(d *models.Device) error {
	id := d.Meta.ID
	allDevices[id] = d
	var err error
	if len(allDevices) == 0 {
		return err
	}
	fmt.Println("ID of devices after saving into:", id)

	return nil
}

