package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/config"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/routes"

	_ "github.com/rlanakomara7/koda-b9-eventhub-backend/docs"
)

// @title						KODA EventHub API
// @version						1.0
// @description					Backend for EventHub Application
// @host						localhost:9000
// @BasePath					/
// @securityDefinitions.apikey	BearerToken
// @in							header
// @name						Authorization
// @description					Bearer Token used as identity for accessing backend
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

	//redis

	fmt.Println("REDIS HOST:", os.Getenv("RDB_HOST"))
	fmt.Println("REDIS PORT:", os.Getenv("RDB_PORT"))
	fmt.Println("REDIS USER:", os.Getenv("RDB_USER"))

	rc := config.NewRedisClient(
		os.Getenv("RDB_HOST"),
		os.Getenv("RDB_PORT"),
		os.Getenv("RDB_USER"),
		os.Getenv("RDB_PASS"),
	)
	if err := rc.Connect(); err != nil {
		log.Fatal(err)
	}

	router := routes.InitMainRouter(db, rc.Client)

	router.Run(":9000")
}
