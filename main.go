package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/deposinator/config"
	"github.com/deposinator/server"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	c, err := config.NewConfig().Load()
	if err != nil {
		log.Fatal("error loading configuration: ", err.Error())
	}

	db, err := sqlx.Connect("postgres", fmt.Sprintf("postgres://%s/postgres?sslmode=disable&user=%s&password=%s&timezone=Europe/Athens", c.POSTGRES_HOST, c.POSTGRES_USER, c.POSTGRES_PASS))
	if err != nil {
		log.Fatal("error connecting to db: ", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	go func() {
		m := map[string]string{"admin": "secret"}
		err := server.Run(m)
		if err != nil {
			log.Println("http server returned error: ", err)
		}
	}()

	log.Println("up and running!")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
