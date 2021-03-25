package controllers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type valueController struct{}

// NewValue is constructor for new instances of Controller interface
func NewValue() Controller {
	return &valueController{}
}

// HANDLERS

func (*valueController) GetAll(c echo.Context) error{
	devID := c.Param("devId")
	netID := c.Param("netId")
	allIds := map[string]string{"netId": netID, "devId": devID}
	reply, _, err := WedgeCallValue(allIds)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	log.Println("Response from GET ALL Value")

	return c.JSON(http.StatusOK, reply)
}



func (*valueController) GetOne(c echo.Context)error{
	valID := c.Param("valId")
	devID := c.Param("devId")
	netID := c.Param("netId")
	allIds := map[string]string{"netId": netID, "devId": devID, "valId": valID}

	_, reply, err := WedgeCallValue(allIds)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	log.Println("Response from GET ONE Value")

	return c.JSON(http.StatusOK, reply)
}

func (*valueController) Update(c echo.Context)error{
	return nil
}


