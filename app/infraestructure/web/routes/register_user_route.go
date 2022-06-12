package routes

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/spro80/golangCleanArchitecture/app/infraestructure/controllers/registerUserController"
)

type RegisterUserRouteInterface interface {
	HandlerRegisterUserRoute(c echo.Context) error
}

type RegisterUserRouteHandler struct {
	controller registerUserController.RegisterUserControllerInterface
}

func NewRegisterUserRoute(e *echo.Echo, controller registerUserController.RegisterUserControllerInterface) *RegisterUserRouteHandler {
	h := &RegisterUserRouteHandler{controller: controller}
	e.POST("/registerUser", h.HandlerRegisterUserRoute)
	return h
}

func (r RegisterUserRouteHandler) HandlerRegisterUserRoute(c echo.Context) error {

	fmt.Println("[register_user_route] Init in HandlerRegisterUserRoute")

	err := r.controller.HandlerRegisterUserController()
	if err != nil {
		fmt.Printf("[register_user_route] Error: [%s]", err.Error())
	}
	fmt.Println("[register_user_route] End in HandlerRegisterUserRoute")
	return nil

}
