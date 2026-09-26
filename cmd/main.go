package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/config"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/routes"
)

func main() {
	//load env
	if err := godotenv.Load(); err != nil {
		log.Println(err.Error())
		return
	}

	db, err := config.InitDB()
	if err != nil {
		log.Println("FAILED TO CONNECT DB")
		return
	}
	defer db.Close()

	err = config.PingDB(db)
	if err != nil {
		log.Println("PING TO DB FAILED", err.Error())
		return
	}
	log.Println("DB CONNECTED")

	router := routes.InitMainRouter(db)

	router.Run(":9000")
}
