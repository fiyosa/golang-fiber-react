package main

import (
	"go-fiber-react/bootstrap"
	"go-fiber-react/config"
)

func main() {
	f := bootstrap.Init()

	f.Listen(":" + config.APP_PORT)
}
