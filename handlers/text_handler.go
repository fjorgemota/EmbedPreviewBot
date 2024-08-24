package handlers

import (
	"github.com/fjorgemota/EmbedPreviewBot/transformer"
	tele "gopkg.in/telebot.v3"
	"log"
)

func OnTextHandler(ctx tele.Context) error {
	msg := ctx.Message().Text
	if msg == "" {
		return nil
	}

	transformedURL, localErr := transformer.TransformURL(msg)
	if localErr != nil {
		log.Printf("Error transforming URL: %v, User Query: %s", localErr, msg)
		return ctx.Reply(&tele.Message{
			Text: "Invalid URL provided. Please check it. The provided URL could not be processed.",
		})
	}

	result := tele.Message{
		Text: transformedURL,
	}
	result.PreviewOptions = &tele.PreviewOptions{
		LargeMedia: true,
		URL:        transformedURL,
	}

	return ctx.Reply(&result)
}
