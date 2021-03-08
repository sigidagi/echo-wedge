package controllers

import (
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
	allIds := map[string]string{"netId": netID, "devId": devID, "valId": valID}
	reply, _, err := WedgeCallState(allIds)
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
	allIds := map[string]string{"netId": netID, "devId": devID, "valId": valID, "stateId": stateID}

	_, reply, err := WedgeCallState(allIds)
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
	// stateID := c.Param("stateId")

	// state := &client.StateWg{}
	// if err := c.Bind(state); err != nil {
	// 	return err
	// }
	// reply, err := WedgeCall("state", "put", state, []string{stateID})
	// if err != nil {
	// 	return c.JSON(http.StatusNotFound, err)
	// }
	// log.Println("Response from Update State")

	// return c.JSON(http.StatusOK, reply)
	return nil
}
