# ETHCHESS Bot — Tewodros ♟️🤖

A Telegram bot for the **ETHCHESS** chess club that facilitates Lichess chess games, retrieves player ratings, and provides AI-powered chat using Google Gemini.

## Features

- **Chess Challenges** — Create open Lichess challenges directly from Telegram
  - Blitz (rated & unrated)
  - Bullet (rated & unrated)
  - Custom time controls
- **Player Lookup** — Retrieve a Lichess user's bullet rating
- **AI Chat** — Responds to mentions and replies using Google Gemini models
- **Welcome Messages** — AI-generated greetings for new group members

## Bot Commands

| Command | Description |
|---|---|
| `/start` | Introduce the bot and list available commands |
| `/blitz` | Create an unrated blitz game (3+0) |
| `/blitzr` | Create a rated blitz game (3+2) |
| `/bullet` | Create an unrated bullet game (1+0) |
| `/bulletr` | Create a rated bullet game (1+0) |
| `/open x y` | Create a custom game (x = seconds, y = increment) |
| `/user <username>` | Look up a Lichess user's bullet rating |

## Project Structure

```
.
├── main.go                  # Bot entry point, command handlers, Lichess challenge logic
├── go.mod                   # Go module definition
├── go.sum                   # Dependency checksums
├── Procfile                 # Heroku deployment config
└── internal/
    ├── gemini/
    │   └── gemini.go        # Google Gemini AI integration
    └── lichess/
        ├── getLichessUser.go # Lichess API user lookup
        └── userTypes.go     # Lichess data structures
```

## Prerequisites

- [Go](https://go.dev/) 1.25.5 or later
- A [Telegram Bot Token](https://core.telegram.org/bots#botfather)
- A [Google Gemini API Key](https://ai.google.dev/)

## Setup

1. **Clone the repository**

   ```bash
   git clone https://github.com/YeiyoNathnael/ethchess-bot-tewdros.git
   cd ethchess-bot-tewdros
   ```

2. **Install dependencies**

   ```bash
   go mod download
   ```

3. **Configure environment variables**

   Create a `.env` file in the project root:

   ```env
   TOKEN=your-telegram-bot-token
   GEMINI_API_KEY=your-google-gemini-api-key
   ```

4. **Build and run**

   ```bash
   go build -o bin/ethchess-bot-tewdros
   ./bin/ethchess-bot-tewdros
   ```

## Deployment

The project includes a `Procfile` for deployment on [Heroku](https://www.heroku.com/):

```
worker: bin/ethchess-bot-tewdros
```

Set the `TOKEN` and `GEMINI_API_KEY` environment variables in your Heroku app settings.

## Tech Stack

- **Language** — [Go](https://go.dev/)
- **Telegram Framework** — [gotgbot](https://github.com/PaulSonOfLars/gotgbot)
- **AI** — [Google Gemini API](https://ai.google.dev/) via `google.golang.org/genai`
- **Chess Platform** — [Lichess API](https://lichess.org/api)
- **Markdown** — [goldmark-tgmd](https://github.com/Mad-Pixels/goldmark-tgmd) for Telegram-compatible Markdown