package health

import (
	controller "arctic-wolf-risk-manager/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
}

func (h Controller) getHealth(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"health": "OK",
	})
}

func NewHealthController(e *echo.Echo) controller.HealthController {
	hctrlr := Controller{}
	e.GET("/v1/health", hctrlr.getHealth)
	return e
}
