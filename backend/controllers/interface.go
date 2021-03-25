package controllers

import (
	"github.com/labstack/echo/v4"
)

// Controller is an interface that defines handler methods
type Controller interface {
	GetAll(c echo.Context) error
	GetOne(c echo.Context) error
	Update(c echo.Context) error
}