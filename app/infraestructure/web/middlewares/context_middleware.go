package middlewares

import (
	"github.com/labstack/echo/v4"
	utils_context "github.com/spro80/golangCleanArchitecture/app/shared/utils/context"
)

func ContextMiddleWare(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		headerTraceParent := c.Request().Header.Get(utils_context.TraceParent)
		//log.Info("header traceparent found %s", headerTraceParent)
		ctx := utils_context.CreateTraceContext(headerTraceParent)
		//log.Info("TraceID %s", ctx.Value("TraceID").(string))
		c.Set("traceContext", ctx)

		return next(c)
	}
}
