package main

import (
	"github.com/OTeeEnabor/echobootcamp/constant"
	"github.com/OTeeEnabor/echobootcamp/router"
	"github.com/OTeeEnabor/echobootcamp/server"
	"github.com/labstack/echo/v4"
)

func main() {
	//  initialize echo app
	app := echo.New()

	// render the templates
	app.Renderer = constant.LoadTemplate()

	//  load static  files
	constant.LoadStatic(app)
	
	// define URL
	router.LoadAllRouters(app)
	// setup the server 
	server.RunServer(app)
	//  start the server
	app.Start(":3000")
}