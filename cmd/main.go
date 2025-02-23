package main

import (
	"log"

	"github.com/deposinator/config"
	"github.com/deposinator/db"
	"github.com/deposinator/server"
	_ "github.com/lib/pq"
)

func main() {
	c, err := config.Load()
	if err != nil {
		log.Fatal("error loading configuration: ", err.Error())
	}

	sqlDB, err := db.NewDB(c.POSTGRES_HOST, c.POSTGRES_USER, c.POSTGRES_PASS)
	if err != nil {
		log.Fatal("error connecting to db: ", err)
	}

	if err := server.Run(sqlDB, c.ENVIRONMENT, c.HTTP_PORT, c.STORE_KEY); err != nil {
		log.Fatal("http server returned error: ", err)
	}
}
