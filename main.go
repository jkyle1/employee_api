package main

import (
	"../employee_api/app"
	"../employee_api/config"
)

func main() {
	config := config.GetConfig()
	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
