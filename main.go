package main

import (
	"go-trading/utils"
	"gotrading/config"
	"log"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	log.Println("test")
}
