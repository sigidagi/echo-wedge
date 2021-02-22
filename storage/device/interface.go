package device

import(
	"github.com/io-m/echo-wedge/models"
)

// StorageDevice is an interface that defines methods for interacting with storage that
// holds data / DB, file et...
type StorageDevice interface {
	Save(d *models.Device) error
}