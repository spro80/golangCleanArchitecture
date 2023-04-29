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
	userAddInput    source.FromApiInterface
	userDeleteInput source.FromApiInterface
}

func NewUserAddRoute(e *echo.Echo, userAddInput source.FromApiInterface, userUpdateInput source.FromApiInterface, userDeleteInput source.FromApiInterface) *userAddRouteHandler {
	user := &userAddRouteHandler{userAddInput, userDeleteInput}
	e.POST("/api/v1/user/user-add", userAddInput.FromApi, middlewares.ContextMiddleWare)
	e.PUT("/api/v1/user/user-update", userUpdateInput.FromApi, middlewares.ContextMiddleWare)
	e.DELETE("/api/v1/user/delete-user/userId/:userId", userDeleteInput.FromApi, middlewares.ContextMiddleWare)

	return user
}
