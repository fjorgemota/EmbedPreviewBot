package main

import (
	"errors"
	tele "gopkg.in/telebot.v3"
	"log"
	"net/http"
	"net/url"
	"time"
)

type botOptions struct {
	Token    string
	Username string
	Client   *http.Client
	BaseURL  string
	Port     string
}

func initBot(options botOptions) (*tele.Bot, error) {
	var err error
	botSettings := tele.Settings{
		Token:  options.Token,
		Client: options.Client,
	}
	if options.BaseURL != "" || options.Port != "" {
		botSettings.Poller, err = initWebhookPoller(options)
	}
	if botSettings.Poller == nil && err == nil {
		log.Printf("Polling Telegram API")
		botSettings.Poller = &tele.LongPoller{
			Timeout: 60 * time.Second,
		}
	}
	var bot *tele.Bot
	if err == nil {
		bot, err = tele.NewBot(botSettings)
	}
	if err == nil && options.Username != "" && bot.Me.Username != options.Username {
		err = errors.New("bot username does not match the one provided")
	}
	return bot, err
}

func initWebhookPoller(options botOptions) (tele.Poller, error) {
	if options.Port == "" {
		return nil, nil
	}
	log.Printf("Starting server")
	var handler *tele.Webhook
	var urlPath *url.URL
	var err error
	if options.BaseURL != "" {
		log.Printf("Listening to Webhook")
		var parsedBaseUrl *url.URL
		parsedBaseUrl, err = url.Parse(options.BaseURL)
		if err == nil {
			urlPath, err = url.Parse("/" + options.Token)
		}
		if err == nil && options.BaseURL != "" {
			fullUrl := parsedBaseUrl.ResolveReference(urlPath)
			handler = &tele.Webhook{
				Endpoint: &tele.WebhookEndpoint{
					PublicURL: fullUrl.String(),
				},
			}
		}
	}
	if err == nil {
		mux := http.NewServeMux()
		if options.Username != "" {
			mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.Redirect(w, r, "https://t.me/"+options.Username, http.StatusMovedPermanently)
			}))
		}
		if handler != nil {
			mux.Handle(urlPath.Path, handler)
		}
		mux.Handle("/status", http.HandlerFunc(statusEndpoint))
		go func() {
			serveErr := http.ListenAndServe(":"+options.Port, mux)
			if serveErr != nil {
				log.Fatal("Error when listening: ", err)
			}
		}()
	}
	return handler, err
}

func statusEndpoint(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("OK"))
	if err != nil {
		log.Println("Error while processing status endpoint: ", err)
	}
}
