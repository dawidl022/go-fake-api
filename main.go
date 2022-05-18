package main

import (
	"log"
	"server/config"
	"server/server"
)

func main() {
	conf, err := config.LoadDefaults()
	if err != nil {
		log.Fatal("Unable to load config")
	}
	server.StartServer(conf)
}
