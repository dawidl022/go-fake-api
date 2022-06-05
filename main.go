package main

import (
	"log"

	"github.com/dawidl022/go-fake-api/config"
	"github.com/dawidl022/go-fake-api/server"
)

func main() {
	conf, err := config.LoadDefaults()
	if err != nil {
		log.Fatal("Unable to load config")
	}
	server.StartServer(conf)
}
