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
	_, reply, err := WedgeCall("network", "get", nil, []string{})
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	
	return c.JSON(http.StatusOK, reply)
}



func (*networkController) GetOne(c echo.Context)error{
	netID := c.Param("netId")
	_, reply, err := WedgeCall("network", "get", nil, []string{netID})
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	// var tcpClient = client.TCPClient{
	// 	Host: "localhost",
	// 	Port: 8051,
	// }
	// tcpClient.Start()

	// in = client.JSONRPC {
	// 	ID:      uuid.New().String(),
	// 	Jsonrpc: "2.0",
	// 	Method:  "GET",
	// 	Params: client.Params{
	// 		URL: fmt.Sprintf("/network/%s", netID),
	// 	},
	// }

	// b, err := json.Marshal(in)
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 		return err
	// 	}
	
	// if err = tcpClient.Write(b); err != nil {
	// 	return err
	// }
	// reply := tcpClient.Read()


	log.Println("Response from GET ONE Network")

	return c.JSON(http.StatusOK, reply)
}

func (*networkController) Update(c echo.Context)error{
	return nil
}
