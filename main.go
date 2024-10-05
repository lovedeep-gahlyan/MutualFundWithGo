package main

import (
	"log"
	"mutualfund/config"
	"mutualfund/server"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
	log.Println("Starting Mutual Fund App")

	// taking runners.toml file
	log.Println("Initializig configuration")
	config := config.InitConfig("mutualfund")

	log.Println("Initializig database")
	dbHandler := server.InitDatabase(config)

	log.Println("Initializig HTTP sever")
	httpServer := server.InitHttpServer(config, dbHandler)

	httpServer.Start()
}