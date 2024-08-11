package handlers

import (
	tele "gopkg.in/telebot.v3"
	"strings"
)

func StartCommand(ctx tele.Context) error {
	message := strings.Replace("Welcome to YourBotUsername!\n\n"+
		"**YourBotUsername** is a Telegram inline bot that transforms Twitter, X (formerly Twitter), and Instagram URLs into optimized versions with enhanced media previews. This bot is particularly useful for quickly sharing media-rich content with large previews directly within Telegram chats.\n"+
		" Here's how to use it:\n"+
		"1. Open any chat (with this bot added) and type @YourBotUsername followed by a URL.\n"+
		"2. For example, try entering:\n"+
		"   - @YourBotUsername https://twitter.com/someuser/status/12345\n"+
		"   - @YourBotUsername https://instagram.com/reels/1234\n\n"+
		"The bot will automatically transform Twitter and Instagram URLs to their equivalent versions for preview.\n"+
		"You can find the source code for this bot on GitHub: [EmbedPreviewBot Repository](https://github.com/fjorgemota/EmbedPreviewBot)", "YourBotUsername", ctx.Bot().Me.Username, -1)

	return ctx.Send(message, &tele.SendOptions{ParseMode: tele.ModeMarkdown})
}
