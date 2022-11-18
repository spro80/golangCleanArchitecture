package routes

import (
	"fmt"
	controllers "github.com/spro80/golangCleanArchitecture/app/interfaces/input/controllers/templateController"

	"github.com/labstack/echo/v4"
)

type TemplateRouteInterface interface {
	HandlerTemplateRoute() error
}

type TemplateRouteHandler struct {
	templateController controllers.TemplateControllerInterface
}

func NewTemplateRoute(e *echo.Echo, templateController controllers.TemplateControllerInterface) *TemplateRouteHandler {
	h := &TemplateRouteHandler{templateController: templateController}
	e.GET("/template", h.HandlerTemplateRoute)
	return h
}

func (t TemplateRouteHandler) HandlerTemplateRoute(c echo.Context) error {
	fmt.Println("[template_route.go] Init in HandlerTemplateRoute")

	err := t.templateController.HandlerTemplateController()
	if err != nil {
		fmt.Printf("[template_route.go] Error: [%s]", err.Error())
	}
	fmt.Println("[template_route.go] End in TemplateControllerRoute")
	return nil

}
