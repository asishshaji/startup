package main

import (
	"github.com/asishshaji/startup/apps/auth/delivery"
	"github.com/asishshaji/startup/server"
)

var (
	httpRouter delivery.Router = delivery.NewMuxRouter()
)

func main() {

	app := server.NewApp(&httpRouter, ":9090")

	app.Run()
}
