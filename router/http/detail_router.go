package http

import (
	"github.com/labstack/echo/v4"
	"github.com/OTeeEnabor/echobootcamp/controller/context/pages"
)
// details router
func DetailsRouter(app *echo.Echo) {
	app.GET("/:productId", pages.DetailsContext)
}