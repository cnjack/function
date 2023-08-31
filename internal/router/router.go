package router

import (
	"github.com/cnjack/cloudfunc/internal/controller"
	"github.com/labstack/echo/v4"
)

func NewRouter(ctl controller.IController) *echo.Echo {
	e := echo.New()

	e.POST("/function/:provider/:function_id/invoke", ctl.FunctionInvoke)

	return e
}
