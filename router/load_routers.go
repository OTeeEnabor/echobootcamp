package router

import (
	"github.com/OTeeEnabor/echo_learn/router/http"
	"github.com/labstack/echo/v4"
)

func LoadAllRouters(app *echo.Echo) {
	http.IndexRouter(app)
}