package main

import (
	"stephen/server/pkg/config"
	"stephen/server/pkg/server"
	"log"
)

var cfg *config.Config
var err error

func init() {
	log.Print("Server is starting...")

	// get a config
	cfg, err = config.NewConfig()
	if err != nil {
		log.Fatal("Config init failed", err)
	}

}

func main() {
	server.Start(cfg)
}