package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/spro80/golangCleanArchitecture/app/infraestructure/web/middlewares"
	"github.com/spro80/golangCleanArchitecture/app/interfaces/input/source"
)

type userAddRouteInterface interface {
	HandlerRegisterUserRoute(serverContext echo.Context) error
}

type userAddRouteHandler struct {
	userAddInput source.FromApiInterface
}

func NewUserAddRoute(e *echo.Echo, userAddInput source.FromApiInterface) *userAddRouteHandler {
	user := &userAddRouteHandler{userAddInput}
	e.POST("/api/v1/user/add-user", userAddInput.FromApi, middlewares.ContextMiddleWare)
	return user
}
