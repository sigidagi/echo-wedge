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
	reply, _, err := WedgeCall("device", "GET", nil, []string{netID})
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	
	log.Println("Response from GET ALL Device", reply)
	return c.JSON(http.StatusOK, reply)
}



func (*deviceController) GetOne(c echo.Context)error{
	devID := c.Param("devId")
	netID := c.Param("netId")
	allIDs := []string{netID, devID}
	fmt.Println(allIDs)
	reply, _, err := WedgeCall("device", "GET", nil, allIDs)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	log.Println("Response from GET ONE Device")
	return c.JSON(http.StatusOK, reply)
}

func (*deviceController) Update(c echo.Context)error{
	return nil
}