package resource

import (
	"net/http"

	"github.com/labstack/echo"
)

func (tt WodResource) Get(c echo.Context) error {

	return c.JSON(http.StatusOK, "works")
}

type WodResource struct {
}
