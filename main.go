package main

import (
	"flag"
	"miaospeed-gateway/config"
	"miaospeed-gateway/log"
	"miaospeed-gateway/service"
)

func main() {
	log.Infof("Miaospeed Gateway %s build %s", config.Version, config.Build)
	log.Debugf("Commit: %s", config.Commit)

	var confpath string
	flag.StringVar(&confpath, "c", "config.yaml", "config file path")
	err := config.Load(confpath, &config.GConf)
	if err != nil {
		log.Fatalf("Error when parsing config, %s", err.Error())
	}

	service.LaunchServer()
}
