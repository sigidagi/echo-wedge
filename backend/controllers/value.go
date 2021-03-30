package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type valueController struct{}

// NewValue is constructor for new instances of Controller interface
func NewValue() Controller {
	return &valueController{}
}

// HANDLERS

func (*valueController) GetAll(c echo.Context) error {
	devID := c.Param("devId")
	netID := c.Param("netId")

	url := fmt.Sprintf("/network/%s/device/%s/value", netID, devID)
	reply, err := WedgeCallItemList(url)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	log.Println("Response from GET ALL Value")

	return c.JSON(http.StatusOK, reply)
}

func (*valueController) GetOne(c echo.Context) error {
	valID := c.Param("valId")
	devID := c.Param("devId")
	netID := c.Param("netId")

	url := fmt.Sprintf("/network/%s/device/%s/value/%s", netID, devID, valID)
	reply, err := WedgeCallValue(url)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	log.Println("Response from GET ONE Value")

	return c.JSON(http.StatusOK, reply)
}

func (*valueController) Update(c echo.Context) error {
	return nil
}
