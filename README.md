# Tewodros

Tewodros is the Telegram bot for the ETHCHESS chess club in Addis Ababa, Ethiopia. Members use it to start Lichess games from a group chat, look up player ratings, and talk to an assistant that answers mentions and replies.

The bot is written in Go and runs as a long-polling worker.

## Features

- Creates open Lichess challenges in blitz, bullet and custom time controls
- Looks up a player's Lichess bullet rating
- Answers mentions and replies with Gemma models through the Google GenAI SDK, with Google Search available to the model
- Greets new group members with a generated welcome message
- Formats model output as Telegram-compatible Markdown

## Commands

| Command | Description |
|---|---|
| `/start` | Introduces the bot and lists its commands |
| `/blitz` | Creates an unrated 3+0 game |
| `/blitzr` | Creates a rated 3+2 game |
| `/bullet` | Creates an unrated 1+0 game |
| `/bulletr` | Creates a rated 1+0 game |
| `/open x y` | Creates a custom game with `x` seconds and `y` seconds of increment |
| `/user <username>` | Returns a player's Lichess bullet rating |

The bot responds in chat when it is mentioned, when a member replies to a message, and when a new member joins.

## Getting started

### Requirements

- Go 1.25.5 or later
- A Telegram bot token from [@BotFather](https://t.me/BotFather)
- A [Google AI API key](https://ai.google.dev/)

### Configuration

Set these environment variables:

| Variable | Description |
|---|---|
| `TOKEN` | Telegram bot token |
| `GEMINI_API_KEY` | Google AI API key |
| `DBURL` | libSQL database URL, used by the tournament draft module |
| `lichessAPI` | Lichess API token, used to create tournaments |

### Run

```bash
git clone https://github.com/YeiyoNathnael/ethchess-bot-tewdros.git
cd ethchess-bot-tewdros
go build -o bin/ethchess-bot-tewdros .
./bin/ethchess-bot-tewdros
```

### Docker

```bash
docker build -t tewodros .
docker run --env TOKEN --env GEMINI_API_KEY tewodros
```

## Deployment

The repository includes a multi-stage `Dockerfile` and a `heroku.yml` that runs the bot as a Heroku worker dyno. Set the environment variables in the app's config before deploying.

## Project layout

```text
main.go                    Entry point, dispatcher and chat handler
internal/
  gemini/                  Google GenAI client and response handling
  lichess/                 Lichess challenges, user lookups and tournaments
  tournament/              Tournament draft handling
  db/                      sqlc-generated database access
  schema.sql               Tables for users, organizers, drafts and the audit log
  query.sql                SQL queries used by sqlc
  sqlc.yaml                sqlc configuration
```

## Built with

- [gotgbot](https://github.com/PaulSonOfLars/gotgbot) for the Telegram Bot API
- [Google GenAI SDK for Go](https://pkg.go.dev/google.golang.org/genai)
- [Lichess API](https://lichess.org/api)
- [libsql-client-go](https://github.com/tursodatabase/libsql-client-go) and [sqlc](https://sqlc.dev/)
- [goldmark-tgmd](https://github.com/Mad-Pixels/goldmark-tgmd)
