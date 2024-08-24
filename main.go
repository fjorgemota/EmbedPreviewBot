package main

import (
	"github.com/fjorgemota/EmbedPreviewBot/handlers"
	tele "gopkg.in/telebot.v3"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: 60 * time.Second,
	}
	bot, err := initBot(botOptions{
		Client:   client,
		Token:    os.Getenv("BOT_TOKEN"),
		Username: os.Getenv("BOT_USERNAME"),
		BaseURL:  os.Getenv("BASE_URL"),
		Port:     os.Getenv("PORT"),
	})

	if err == nil {
		// Handle /start command
		bot.Handle("/start", handlers.StartCommand)
		// Handle inline queries
		bot.Handle(tele.OnQuery, handlers.OnQueryHandler)
		// Handle direct messages
		bot.Handle(tele.OnText, handlers.OnTextHandler)
	}
	if err != nil {
		log.Fatal(err)
	}

	bot.Start()
}
