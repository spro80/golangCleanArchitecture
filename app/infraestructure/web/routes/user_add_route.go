package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/spro80/golangCleanArchitecture/app/infraestructure/web/middlewares"
	"github.com/spro80/golangCleanArchitecture/app/interfaces/input/source"
)

type userRouteInterface interface {
	HandlerUserRoute(serverContext echo.Context) error
}

type userRouteHandler struct {
	userGetInput    source.FromApiInterface
	userAddInput    source.FromApiInterface
	userUpdateInput source.FromApiInterface
	userDeleteInput source.FromApiInterface
}

func NewUserAddRoute(e *echo.Echo, userGetInput source.FromApiInterface, userAddInput source.FromApiInterface, userUpdateInput source.FromApiInterface, userDeleteInput source.FromApiInterface) *userRouteHandler {
	user := &userRouteHandler{userGetInput, userAddInput, userUpdateInput, userDeleteInput}
	e.GET("/api/v1/user/user-get", userGetInput.FromApi, middlewares.ContextMiddleWare)
	e.POST("/api/v1/user/user-add", userAddInput.FromApi, middlewares.ContextMiddleWare)
	e.PUT("/api/v1/user/user-update", userUpdateInput.FromApi, middlewares.ContextMiddleWare)
	e.DELETE("/api/v1/user/user-delete/userId/:userId", userDeleteInput.FromApi, middlewares.ContextMiddleWare)

	return user
}
