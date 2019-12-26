package types

import (
	"github.com/labstack/echo"
)

// RESTController is a standard REST method mapper to handler functions
type RESTController struct {
	Get          echo.HandlerFunc
	Find         echo.HandlerFunc
	Post         echo.HandlerFunc
	Put          echo.HandlerFunc
	Patch        echo.HandlerFunc
	Delete       echo.HandlerFunc
	SubResources map[string]RESTController
	Custom       map[string]echo.HandlerFunc
}
