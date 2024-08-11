package handlers

import (
	"github.com/fjorgemota/EmbedPreviewBot/transformer"
	tele "gopkg.in/telebot.v3"
	"log"
)

func OnQueryHandler(ctx tele.Context) error {
	query := ctx.Query().Text
	if query == "" {
		return nil
	}

	transformedURL, localErr := transformer.TransformURL(query)
	if localErr != nil {
		log.Printf("Error transforming URL: %v, User Query: %s", localErr, query)
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
}
