package main

import (
	"encoding/json"
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Get token from the environment variable
	err := godotenv.Load()
	token := os.Getenv("TOKEN")
	if token == "" {
		panic("TOKEN environment variable is empty")
	}

	// Create bot from environment value.
	b, err := gotgbot.NewBot(token, nil)
	if err != nil {
		panic("failed to create new bot: " + err.Error())
	}
	// Create updater and dispatcher.
	dispatcher := ext.NewDispatcher(&ext.DispatcherOpts{
		// If an error is returned by a handler, log it and continue going.
		Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
			log.Println("an error occurred while handling update:", err.Error())
			return ext.DispatcherActionNoop
		},
		MaxRoutines: ext.DefaultMaxRoutines,
	})
	updater := ext.NewUpdater(dispatcher, nil) //&ext.UpdaterOpts{ErrorLog: logger})

	// /start command to introduce the bot
	dispatcher.AddHandler(handlers.NewCommand("start", start))
	dispatcher.AddHandler(handlers.NewCommand("blitz", blitz))
	dispatcher.AddHandler(handlers.NewCommand("blitzr", blitzr))
	dispatcher.AddHandler(handlers.NewCommand("bullet", bullet))
	dispatcher.AddHandler(handlers.NewCommand("bulletr", bulletr))
	dispatcher.AddHandler(handlers.NewCommand("open", open))

	err = updater.StartPolling(b, &ext.PollingOpts{
		DropPendingUpdates: true,
		GetUpdatesOpts: &gotgbot.GetUpdatesOpts{
			Timeout: 9,
			RequestOpts: &gotgbot.RequestOpts{
				Timeout: time.Second * 10,
			},
		},
	})
	if err != nil {
		panic("failed to start polling: " + err.Error())
	}

	// Idle, to keep updates coming in, and avoid bot stopping.
	updater.Idle()
}

func openChallenge(b *gotgbot.Bot, ctx *ext.Context, clockLimit string, clockIncrement string, duelName string, rated bool) error {

	urlL := "https://lichess.org/api/challenge/open"

	postData := url.Values{}
	postData.Set("rated", strconv.FormatBool(rated))
	postData.Set("clock.limit", clockLimit)
	postData.Set("clock.increment", clockIncrement)
	postData.Set("days", "1")
	postData.Set("variant", "standard")
	postData.Set("fen", "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	postData.Set("name", duelName)

	req, _ := http.NewRequest("POST", urlL, strings.NewReader(postData.Encode()))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

	var LichessChallengeResponse struct {
		Url string `json:"url"`
	}

	if err := json.Unmarshal(body, &LichessChallengeResponse); err != nil {
		return fmt.Errorf("parsing failed: %w", err)
	}

	link := LichessChallengeResponse.Url

	_, err := ctx.EffectiveMessage.Reply(b, link, &gotgbot.SendMessageOpts{
		ParseMode: "HTML",
	},
	)
	if err != nil {
		return fmt.Errorf("failed to send source: %w", err)
	}

	return nil

}

func blitz(b *gotgbot.Bot, ctx *ext.Context) error {

	openChallenge(b, ctx, "180", "2", "Grand Blitz Duel", false)

	return nil

}

func blitzr(b *gotgbot.Bot, ctx *ext.Context) error {

	openChallenge(b, ctx, "180", "2", "Grand Blitz Duel", true)

	return nil

}

func bullet(b *gotgbot.Bot, ctx *ext.Context) error {

	openChallenge(b, ctx, "60", "0", "Grand Blitz Duel", false)

	return nil

}

func bulletr(b *gotgbot.Bot, ctx *ext.Context) error {

	openChallenge(b, ctx, "60", "0", "Grand Blitz Duel", true)

	return nil

}

func open(b *gotgbot.Bot, ctx *ext.Context) error {

	args := ctx.Args()

	if len(args) < 2 {

		_, err := ctx.EffectiveMessage.Reply(b, "Please provide a time limit, e.g., /open 300", nil)

		return err

	}

	clockLimit := args[1]

	increment := "0"

	if len(args) > 2 {

		increment = args[2]

	}

	openChallenge(b, ctx, clockLimit, increment, "Open Challenge Duel", false)

	return nil

}

// start introduces the bot.
func start(b *gotgbot.Bot, ctx *ext.Context) error {

	const startMessage = `
Hey\! I’m *Tewodros* ♟️🤖

I’m the official *ETHCHESS* club bot 🏛️  
Right now I’m still warming up, so I can’t chat naturally yet — but I *can* help you start games and throw down some chess battles 💥

*Game commands:* 🎮

/blitz      \- blitz game ⚡  
/blitzr     \- rated blitz game ⚡  

/bullet     \- bullet game   
/bulletr    \- rated bullet game   


/open x y   \- custom time control ⏱️  
\(x \= seconds, y \= increment\)

Rated games affect rating   
Unrated games are just for fun 😄  

More features are coming soon — stay sharp 
`
	_, err := ctx.EffectiveMessage.Reply(b, startMessage, &gotgbot.SendMessageOpts{
		ParseMode: "MarkdownV2",
	})
	if err != nil {
		return fmt.Errorf("failed to send start message: %w", err)
	}
	return nil
}
