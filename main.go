package main

import (
	"github.com/OTeeEnabor/echo_learn/constant"
	"github.com/OTeeEnabor/echo_learn/router"
	"github.com/OTeeEnabor/echo_learn/server"
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