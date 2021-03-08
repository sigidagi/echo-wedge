package controllers

import (
	"log"
	"net/http"

	"github.com/io-m/echo-wedge/client"
	"github.com/labstack/echo/v4"
)

type stateController struct{}

// NewState is constructor for new instances of Controller interface
func NewState() Controller {
	return &stateController{}
}

// HANDLERS

func (*stateController) GetAll(c echo.Context) error{
	valID := c.Param("valId")
	devID := c.Param("devId")
	netID := c.Param("netId")
	reply, _, err := WedgeCall("state", "get", nil, []string{valID, devID, netID})
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	log.Println("Response from GET ALL State: ", reply)

	return c.JSON(http.StatusOK, reply)
}



func (*stateController) GetOne(c echo.Context)error{
	stateID := c.Param("stateId")
	valID := c.Param("valId")
	devID := c.Param("devId")
	netID := c.Param("netId")
	reply, _, err := WedgeCall("state", "get", nil, []string{stateID, valID, devID, netID})
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	log.Println("Response from GET ONE State")
	return c.JSON(http.StatusOK, reply)
}

func (*stateController) Update(c echo.Context)error{
	// netID := c.Param("id")
	// devID := c.Param("devId")
	// valID := c.Param("valId")
	stateID := c.Param("stateId")

	state := &client.StateWg{}
	if err := c.Bind(state); err != nil {
		return err
	}
	reply, _,err := WedgeCall("state", "put", state, []string{stateID})
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	log.Println("Response from Update State")

	return c.JSON(http.StatusOK, reply)
}
