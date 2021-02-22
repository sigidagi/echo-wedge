package echo

import (
	"github.com/labstack/echo/v4"

	"github.com/io-m/echo-wedge/controllers"
)

var dc = controllers.NewDevice()

// RunApp is an entry point to the app
func RunApp() {
	r := echo.New()
	dev := r.Group("/device")
	dev.GET("/", dc.GetAll)
	dev.GET("/:id", dc.GetOne)
	dev.POST("/", dc.Create)


	// ================
	// Server running ...
	r.Logger.Fatal(r.Start("localhost:8080"))
}