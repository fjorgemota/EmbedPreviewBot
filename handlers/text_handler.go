package handlers

import (
	"github.com/fjorgemota/EmbedPreviewBot/transformer"
	tele "gopkg.in/telebot.v3"
	"log"
	"strings"
)

func OnTextHandler(ctx tele.Context) error {
	if ctx.Message() == nil {
		return ctx.Reply("No message found in the context.")
	}

	msg := ctx.Message().Text
	if msg == "" {
		return ctx.Reply("Empty message received. Please provide a URL.")
	}

	transformedURL, localErr := transformer.TransformURL(msg)
	if localErr != nil {
		log.Printf("Error transforming URL: %v, User Query: %s", localErr, msg)
		return ctx.Reply("**Invalid URL provided\\. Please check it\\.**\n The provided URL could not be processed\\.", &tele.SendOptions{
			ParseMode: tele.ModeMarkdownV2,
		})
	}

	transformedURL = strings.ReplaceAll(transformedURL, "_", "\\_")
	transformedURL = strings.ReplaceAll(transformedURL, "-", "\\-")
	transformedURL = strings.ReplaceAll(transformedURL, "*", "\\*")
	transformedURL = strings.ReplaceAll(transformedURL, "[", "\\[")
	transformedURL = strings.ReplaceAll(transformedURL, "`", "\\`")
	transformedURL = strings.ReplaceAll(transformedURL, ".", "\\.")

	return ctx.Reply(transformedURL, &tele.SendOptions{
		ParseMode:             tele.ModeMarkdownV2,
		DisableWebPagePreview: false,
	})
}
