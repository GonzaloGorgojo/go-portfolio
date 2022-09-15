package router

import (
	controller "github.com/GonzaloGorgojo/go-backend-portfolio/src/adapters/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/experience", func(context echo.Context) error { return c.GetExperience(context) })
	e.POST("/experience", func(context echo.Context) error { return c.AddExperience(context) })

	return e
}
