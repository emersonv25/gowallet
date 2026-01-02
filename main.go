package main

import (
	"github.com/emersonv25/gowallet/config"
	"github.com/emersonv25/gowallet/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("Failed to initialize configuration: %s", err.Error())
		return
	}

	router.Initialize()
}
