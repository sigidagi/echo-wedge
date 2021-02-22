package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/io-m/echo-wedge/services/device"
	"github.com/io-m/echo-wedge/models"
)

var(
	sd = device.NewDeviceService()
)


type deviceController struct{}

// NewDevice is constructor for new instances of Controller interface
func NewDevice() ControllerDevice {
	return &deviceController{}
}

// HANDLERS

func (*deviceController) GetAll(c echo.Context) error{
	return nil
}
func (*deviceController) GetOne(c echo.Context)error{
	return nil
}
func (*deviceController) Create(c echo.Context)error{
	device := &models.Device{}
	if err := c.Bind(device); err != nil {
		return err
	}
	savedDevice, err := sd.SaveDevice(device)
	if err != nil{
		return err
	}
	return c.JSON(http.StatusCreated, savedDevice)
}
func (*deviceController) Update(c echo.Context)error{
	return nil
}
func (*deviceController) Delete(c echo.Context)error{
	return nil
}