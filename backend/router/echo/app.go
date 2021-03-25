package echo

import (
	"github.com/labstack/echo/v4"

	"github.com/io-m/echo-wedge/backend/controllers"
)

var (
	dc = controllers.NewDevice()
	nc = controllers.NewNetwork()
	vc = controllers.NewValue()
	sc = controllers.NewState()
	ec = controllers.NewEvent()
) 


// RunApp is an entry point to the app
func RunApp() {
	r := echo.New()

	// Websocket connection
	r.GET("/subscribe", ec.PushUpdates)
	// enpoint for receiveing data from tcp server
	r.POST("/apigtw", ec.Receive)

	// Network endpoints
	r.GET("/network", nc.GetAll)
	r.GET("/network/:netId", nc.GetOne)
	// ============================
	// Device endpoints
	r.GET("/network/:netId/device", dc.GetAll)
	r.GET("/network/:netId/device/:devId", dc.GetOne)
	// ============================
	// Value endpoints
	r.GET("/network/:netId/device/:devId/value", vc.GetAll)
	r.GET("/network/:netId/device/:devId/value/:valId", vc.GetOne)
	// ============================
	// Updating state of the device
	r.GET("/network/:netId/device/:devId/value/:valId/state", sc.GetAll)
	r.GET("/network/:netId/device/:devId/value/:valId/state/:stateId", sc.GetOne)
	r.PUT("/network/:netId/device/:devId/value/:valId/state/:stateId", sc.Update)

	// ============================
	// Server running ...

	r.Logger.Fatal(r.Start("localhost:8000"))

}