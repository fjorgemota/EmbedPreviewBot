# EmbedPreviewBot

**EmbedPreviewBot** is a Telegram inline bot that transforms Twitter, X (formerly Twitter), and Instagram URLs into optimized versions with enhanced media previews. This bot is particularly useful for quickly sharing media-rich content with large previews directly within Telegram chats.

## Features

- **URL Transformation**: Converts Twitter, X, and Instagram URLs to `fxtwitter.com` and `ddinstagram.com` formats, removing unnecessary query strings.
- **Inline Queries**: Works seamlessly within any chat by using the bot's inline query feature.
- **User-friendly Error Handling**: Provides clear feedback for invalid or malformed URLs.

## How to Use

1. **Add the Bot**: You can use this bot in any Telegram chat where it is added.
2. **Inline Query**: 
   - In any chat, type `@EmbedPreviewBot <URL>` to transform and share a URL.
   - For example:
     - `@EmbedPreviewBot https://twitter.com/someuser/status/12345`
     - `@EmbedPreviewBot https://instagram.com/reels/1234`
3. **View and Share**: The bot will return a link with an optimized preview. You can then share this directly in the chat.

## Commands

- **/start**: Displays a welcome message and usage examples.

## Installation

To run this bot yourself, follow these steps:

### Prerequisites

- [Go](https://golang.org/) installed on your machine.
- A Telegram bot token obtained from [BotFather](https://t.me/BotFather).

### Setup

1. Clone this repository:

   ```bash
   git clone https://github.com/fjorgemota/EmbedPreviewBot.git
   cd EmbedPreviewBot
   ```

2. Install the required dependencies:

   ```bash
   go get
   ```

3. Create a `.env` file and add your Telegram bot token:

   ```
   BOT_TOKEN=your-telegram-bot-token
   ```

4. Run the bot:

   ```bash
   go run main.go
   ```

## Development

### Testing

Tests are written using Go's `testing` package. To run tests:

```bash
go test -v
```

### Contributing

Contributions are welcome! If you find a bug or have a feature request, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

For any inquiries, feel free to reach out:

- **GitHub**: [fjorgemota](https://github.com/fjorgemota)
```