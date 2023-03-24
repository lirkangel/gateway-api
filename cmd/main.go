package main

import (
	"fmt"
	"log"

	"github.com/lirkangel/gateway-api/config"
	"github.com/lirkangel/gateway-api/server"
	"github.com/spf13/viper"
)

func main() {
	fmt.Println("start loading")
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	app := server.RunServer()
	if err := app.Run(viper.GetString("port")); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
