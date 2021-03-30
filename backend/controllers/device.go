package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type deviceController struct{}

// NewDevice is constructor for new instances of Controller interface
func NewDevice() Controller {
	return &deviceController{}
}

func (*deviceController) GetAll(c echo.Context) error {
	netID := c.Param("netId")
	url := fmt.Sprintf("/network/%s/device", netID)
	reply, err := WedgeCallItemList(url)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	log.Println("Response from GET ALL Devices", reply)
	return c.JSON(http.StatusOK, reply)
}

func (*deviceController) GetOne(c echo.Context) error {
	devID := c.Param("devId")
	netID := c.Param("netId")

	url := fmt.Sprintf("/network/%s/device/%s", netID, devID)
	reply, err := WedgeCallDevice(url)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	log.Println("Response from GET ONE Device")
	return c.JSON(http.StatusOK, reply)
}

func (*deviceController) Update(c echo.Context) error {
	return nil
}
