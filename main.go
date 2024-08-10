package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)
import tele "gopkg.in/telebot.v3"

func main() {
	client := &http.Client{
		Timeout: 60 * time.Second,
	}
	bot, err := initBot(botOptions{
		Client:  client,
		Token:   os.Getenv("BOT_TOKEN"),
		BaseURL: os.Getenv("BASE_URL"),
		Port:    os.Getenv("PORT"),
	})

	if err == nil {
		// Handle /start command
		bot.Handle("/start", func(c tele.Context) error {
			message := strings.Replace("Welcome to YourBotUsername!\n\n"+
				"**YourBotUsername** is a Telegram inline bot that transforms Twitter, X (formerly Twitter), and Instagram URLs into optimized versions with enhanced media previews. This bot is particularly useful for quickly sharing media-rich content with large previews directly within Telegram chats.\n"+
				" Here's how to use it:\n"+
				"1. Open any chat (with this bot added) and type @YourBotUsername followed by a URL.\n"+
				"2. For example, try entering:\n"+
				"   - @YourBotUsername https://twitter.com/someuser/status/12345\n"+
				"   - @YourBotUsername https://instagram.com/reels/1234\n\n"+
				"The bot will automatically transform Twitter and Instagram URLs to their equivalent versions for preview.\n"+
				"You can find the source code for this bot on GitHub: [EmbedPreviewBot Repository](https://github.com/fjorgemota/EmbedPreviewBot)", "YourBotUsername", bot.Me.Username, -1)

			return c.Send(message, &tele.SendOptions{ParseMode: tele.ModeMarkdown})
		})

		bot.Handle(tele.OnQuery, func(ctx tele.Context) error {
			query := ctx.Query().Text
			if query == "" {
				return nil
			}

			transformedURL, localErr := TransformURL(query)
			if localErr != nil {
				log.Printf("Error transforming URL: %v, User Query: %s", err, query)
				// Create an error message result
				errorResult := &tele.ArticleResult{
					Text:        "Invalid URL provided. Please check it.",
					Title:       "Error: Invalid URL",
					Description: "The provided URL could not be processed.",
				}

				results := tele.Results{errorResult}
				return ctx.Answer(&tele.QueryResponse{
					Results:   results,
					CacheTime: 0,
				})
			}

			// Create an InlineQueryResultArticle with the transformed URL
			result := &tele.ArticleResult{
				Text:        transformedURL,
				Title:       "URL with preview for embed",
				Description: "Click to send the URL with a proper nice preview.",
			}

			// Enable LargeMedia for the message sent with this result
			result.ParseMode = tele.ModeHTML
			result.Content = &tele.InputTextMessageContent{
				Text: transformedURL,
				PreviewOptions: &tele.PreviewOptions{
					LargeMedia: true,
					URL:        transformedURL,
				},
			}

			results := tele.Results{result}
			return ctx.Answer(&tele.QueryResponse{
				Results:   results,
				CacheTime: 24 * 60 * 60, // One day
			})
		})
	}
	if err != nil {
		log.Fatal(err)
	}

	bot.Start()
}
