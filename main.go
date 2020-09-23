package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"tickethub_service/apimodel"
	"tickethub_service/config"
	"tickethub_service/util"
	"tickethub_service/util/log"
	"tickethub_service/midderware"
)

func main() {
	util.Parse()
	err := config.ReadConfigFromFile()
	if err != nil {
		log.Errorf("Set by config file failed use default config: %v", err)
	}
	err = config.InitDB()
	if err != nil {
		log.Fatalf("Failed to init DB")
		os.Exit(1)
	}

	server := gin.New()
	midderwares := make([]gin.HandlerFunc, 0)

	midderwares = append(midderwares, midderware.LogRequest())
	midderwares = append(midderwares, midderware.Cors())
	server.Use(midderwares...)
	apimodel.RegisterRoutes(server)

	hostAddr := fmt.Sprintf("%s:%d", config.SystemConfig.ServerHost, config.SystemConfig.ServerPort)
	if err := server.Run(hostAddr); err != nil {
		log.Errorf("Server failed to run, err: %v", err)
		os.Exit(1)
	}

	os.Exit(1)
}
