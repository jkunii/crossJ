package routers

import (
	"github.com/jkunii/crossJ/global"
	"github.com/jkunii/crossJ/resource"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

type ApplicationRouter struct {
	Wod *resource.WodResource `inject:""`
}

func (r ApplicationRouter) Init(e *echo.Echo) {

	// Swagger
	e.Static("/", "public/swagger/")

	g := e.Group("/crossj")
	g.Use(mw.BasicAuth(func(username, password string) bool {
		if username == global.Cfg.UserName && password == global.Cfg.Secret {
			return true
		}
		return false
	}))

	g.GET("/wod", r.Wod.Get)
}
