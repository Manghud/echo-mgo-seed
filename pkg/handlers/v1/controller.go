package v1

import (
	"github.com/labstack/echo"
	types "github.com/manghud/echo-mgo-seed/pkg/types"
	"net/http"
	"time"
)

// Controller maps routes to handlers
var Controller types.RESTController = types.RESTController{
	Find: findHandler,
}

func findHandler(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]time.Time{
		"time": time.Now(),
	})
}
