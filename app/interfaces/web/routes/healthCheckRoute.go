package routes

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spro80/golangCleanArchitecture/app/shared/config"
)

var startTime time.Time

func initTime() {
	startTime = time.Now()
}

type HandlerCheckRouterInterface interface {
	HandlerHealtCheck() error
}

type HealthCheckRouterHandler struct {
	config config.ConfigInterface
}

type healthCheck struct {
	Status      string `json:"status"`
	Version     string `json:"version"`
	Uptime      string `json:"uptime"`
	Environment string `json:"environment"`
	Region      string `json:"region"`
}

func NewHealthCheckRoute(e *echo.Echo, config config.ConfigInterface) {
	h := &HealthCheckRouterHandler{config: config}
	e.GET("/health", h.HandlerHealtCheck)
}

func (h HealthCheckRouterHandler) HandlerHealtCheck(c echo.Context) error {

	versionApp, _ := h.config.GetVersionApp() //os.Getenv("VERSION_APP")
	environment, _ := h.config.GetEnvironment()
	region, _ := h.config.GetRegion()

	healthCheck := healthCheck{
		Status:      "UP",
		Version:     versionApp,
		Uptime:      time.Since(startTime).String(),
		Environment: environment,
		Region:      region,
	}

	return c.JSON(http.StatusOK, healthCheck)
}
