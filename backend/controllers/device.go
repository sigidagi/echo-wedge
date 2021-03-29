package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type deviceController struct{}

// NewDevice is constructor for new instances of Controller interface
func NewDevice() Controller {
	return &deviceController{}
}

// HANDLERS

func (*deviceController) GetAll(c echo.Context) error{
	netID := c.Param("netId")
	reply, err := WedgeCallAllDevices(fmt.Sprintf("/network/%s/device", netID))
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	
	log.Println("Response from GET ALL Device", reply)
	return c.JSON(http.StatusOK, reply)
}



func (*deviceController) GetOne(c echo.Context)error{
	devID := c.Param("devId")
	netID := c.Param("netId")

	reply, err := WedgeCallOneDevice(fmt.Sprintf("/network/%s/device/%s", netID, devID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	log.Println("Response from GET ONE Device")
	return c.JSON(http.StatusOK, reply)
}

func (*deviceController) Update(c echo.Context)error{
	return nil
}

