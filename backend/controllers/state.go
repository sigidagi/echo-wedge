package controllers

import (
	"fmt"
	"log"
	"net/http"

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
	reply, err := WedgeCallAllStates(fmt.Sprintf("/network/%s/device/%s/value/%s/state", netID, devID, valID))
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
	reply, err := WedgeCallOneState(fmt.Sprintf("/network/%s/device/%s/value/%s/state/%s", netID, devID, valID, stateID))
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	log.Println("Response from GET ONE State")
	return c.JSON(http.StatusOK, reply)
	
}

func (*stateController) Update(c echo.Context)error{

	return nil
}




