package controller

import "github.com/labstack/echo/v4"

var _ IController = (*Controller)(nil)

type IController interface {
	Function(c echo.Context) error
}

type Controller struct {
}

func (ctl *Controller) Function(c echo.Context) error {
	return c.String(200, "Hello World")
}
