package controllers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type networkController struct{}

// NewNetwork is constructor for new instances of Controller interface
func NewNetwork() Controller {
	return &networkController{}
}

// HANDLERS

func (*networkController) GetAll(c echo.Context) error{
	allIds := map[string]string{}
	reply, err := WedgeCallNetwork(allIds)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	
	return c.JSON(http.StatusOK, reply)
}



func (*networkController) GetOne(c echo.Context)error{
	netID := c.Param("netId")
	allIds := map[string]string{"netId": netID}
	reply, err := WedgeCallNetwork(allIds)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	
	log.Println("Response from GET ONE Network")

	return c.JSON(http.StatusOK, reply)
}

func (*networkController) Update(c echo.Context)error{
	return nil
}



