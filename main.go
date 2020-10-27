package main

import "github.com/asishshaji/startup/server"

func main() {

	app := server.NewApp()
	app.Run(":9090")
}
