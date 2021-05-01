package version

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go-course/demo/app/shared/domain/constants"
	"net/http"
	"os"
)

func init() {
}

type healthHandler struct {
	version string
}

func NewHealthHandler(e *echo.Echo, version string) {
	h := &healthHandler{
		version: version,
	}
	e.GET(fmt.Sprintf("%s/version", os.Getenv(constants.BASE_PATH)), h.HealthCheck)
}

func (h *healthHandler) HealthCheck(c echo.Context) error {

	healthCheck := HealthCheck{
		App:     constants.APP,
		Version: h.version,
		Env:     os.Getenv(constants.ENVIRONMENT_TYPE),
		Author:  constants.AUTHOR,
		Name:    constants.ORIGIN,
	}
	return c.JSON(http.StatusOK, healthCheck)
}
