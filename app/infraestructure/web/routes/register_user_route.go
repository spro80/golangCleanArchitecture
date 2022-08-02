package routes

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/spro80/golangCleanArchitecture/app/domain/entities"
	"github.com/spro80/golangCleanArchitecture/app/infraestructure/controllers/registerUserController"
)

type RegisterUserRouteInterface interface {
	HandlerRegisterUserRoute(c echo.Context) error
}

type RegisterUserRouteHandler struct {
	controller registerUserController.ControllerRegisterUserInterface
}

func NewRegisterUserRoute(e *echo.Echo, controller registerUserController.ControllerRegisterUserInterface) *RegisterUserRouteHandler {
	h := &RegisterUserRouteHandler{controller: controller}
	e.POST("/registerUser", h.HandlerRegisterUserRoute)
	return h
}

func (r RegisterUserRouteHandler) HandlerRegisterUserRoute(c echo.Context) error {

	fmt.Println("[register_user_route] Init in HandlerRegisterUserRoute")

	var err0 error
	u := new(entities.User)
	if err0 = c.Bind(u); err0 != nil {
		return nil
	}

	user, statusCode, err := r.controller.HandlerRegisterUserController(u)
	fmt.Printf("[register_user_route] user: [%v]", user)
	fmt.Printf("[register_user_route] statusCode: [%d]", statusCode)

	if err != nil {
		fmt.Printf("[register_user_route] Error: [%s]", err.Error())
	}
	fmt.Println("[register_user_route] End in HandlerRegisterUserRoute")
	return nil

}
