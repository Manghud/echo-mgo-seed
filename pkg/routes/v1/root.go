package v1

import (
	rootHandlers "github.com/manghud/echo-mgo-seed/pkg/handlers/v1"
	todoHandlers "github.com/manghud/echo-mgo-seed/pkg/handlers/v1/todo"

	"github.com/labstack/echo"
)

// RegisterRootRouter adds a root router to the given echo server param and root route name param
func RegisterRootRouter(server *echo.Echo, route string) {
	server.GET("/", rootHandlers.Controller.Find)
	rootRouter := server.Group(route)
	registerRootRoutes(rootRouter)
	registerTodoRoutes(rootRouter.Group("/todo"))
}

func registerRootRoutes(router *echo.Group) {
	router.GET("/", rootHandlers.Controller.Find)
}

func registerTodoRoutes(router *echo.Group) {
	router.GET("/", todoHandlers.Controller.Find)
}
