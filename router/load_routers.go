package router

import (
	"github.com/OTeeEnabor/echobootcamp/router/http"
	"github.com/labstack/echo/v4"
)

func LoadAllRouters(app *echo.Echo) {
	http.IndexRouter(app)
	//form router
	http.FormRouter(app)

	//details details
	http.DetailsRouter(app)
}