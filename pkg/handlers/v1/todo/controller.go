package todo

import (
	"github.com/labstack/echo"
	"github.com/manghud/echo-mgo-seed/pkg/types"
)

// Controller maps routes to handlers
var Controller types.RESTController = types.RESTController{
	Find: findHandler,
}

func findHandler(ctx echo.Context) error {
	return ctx.JSON(200, []string{})
}
