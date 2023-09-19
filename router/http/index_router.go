package http

import (
	"github.com/OTeeEnabor/echo_learn/controller/context/pages"
	"github.com/labstack/echo/v4"
)
//  define a function that routes to the index page
func IndexRouter(app *echo.Echo) {
	app.GET("/", pages.IndexContext)
}