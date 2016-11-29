package main

import (
	"os"

	"github.com/dimiro1/banner"
	"github.com/jkunii/crossJ/global"
	"github.com/jkunii/crossJ/helper"
	"github.com/jkunii/crossJ/routers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/gommon/log"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/facebookgo/inject"
	mw "github.com/labstack/echo/middleware"
)

var (
	version string
)

func getVersion() string {
	if len(version) == 0 {
		return "devel"
	}
	return version
}

func showBanner() {

	in, err := os.Open("banner.txt")
	defer in.Close()
	helper.PanicErr(err)
	banner.Init(os.Stdout, true, false, in)
}

func main() {
	os.Setenv("APP_VERSION", getVersion())
	// Initialize config
	var config global.Config
	config.Init()

	if len(os.Args) == 2 && os.Args[1] == "-v" {
		showBanner()
		return
	}

	if global.Cfg.ShowBanner {
		showBanner()
	}

	e := echo.New()

	// Middleware
	e.Use(mw.Logger())
	e.Use(mw.Recover())

	// Log
	e.Logger().SetLevel(log.Lvl(global.Cfg.LogLevel))

	// Initialize using DI
	var appRouters routers.ApplicationRouter

	var g inject.Graph
	err := g.Provide(&inject.Object{Value: &appRouters})
	helper.PanicErr(err)

	err = g.Populate()
	helper.PanicErr(err)

	appRouters.Init(e)

	// Start server
	std := standard.New(":" + global.Cfg.Port)
	std.SetHandler(e)

	// Dont stop me if I am doing something
	err = gracehttp.Serve(std.Server)
	helper.PanicErr(err)
}
