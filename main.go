package main

import (
	"log"
	"os"
	"time"
	"tsheet/app/route"

	"github.com/joho/godotenv"
	tele "gopkg.in/telebot.v3"
	teleMid "gopkg.in/telebot.v3/middleware"
)

func main() {
	godotenv.Load(".env")
	pref := tele.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	route.MainRoute(b)

	b.Use(teleMid.Logger())
	b.Start()	
}
