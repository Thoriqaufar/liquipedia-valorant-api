package main

import (
	"github.com/thoriqaufar/liquipedia-valorant-api/config"
	"github.com/thoriqaufar/liquipedia-valorant-api/route"
)

func main() {
	config.ConnectDB()

	e := route.Init()

	e.Logger.Fatal(e.Start(":8080"))
}
