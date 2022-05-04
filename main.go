package main

import (
	"server/config"
	"server/server"
)

func main() {
	server.StartServer(config.LoadDefaults())
}
