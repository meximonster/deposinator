package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

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

	err = db.NewDB(c.POSTGRES_HOST, c.POSTGRES_USER, c.POSTGRES_PASS)
	if err != nil {
		log.Fatal("error connecting to db: ", err)
	}

	go func() {
		err := server.Run()
		if err != nil {
			log.Println("http server returned error: ", err)
		}
	}()

	log.Println("server up and running!")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
