package main

import (
	"fmt"

	"github.com/diogosevilha/goLangProjects/config"
	"github.com/diogosevilha/goLangProjects/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// Initialize Configs
	err := config.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Initialize Router
	router.Initialize()
}
