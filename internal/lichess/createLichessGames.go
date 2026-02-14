package lichess

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func OpenChallenge(b *gotgbot.Bot, ctx *ext.Context, clockLimit string, clockIncrement string, duelName string, rated bool) error {

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

func Blitz(b *gotgbot.Bot, ctx *ext.Context) error {

	OpenChallenge(b, ctx, "180", "0", "Grand Blitz Duel", false)

	return nil

}

func Blitzr(b *gotgbot.Bot, ctx *ext.Context) error {

	OpenChallenge(b, ctx, "180", "2", "Grand Blitz Duel", true)

	return nil

}

func Bullet(b *gotgbot.Bot, ctx *ext.Context) error {

	OpenChallenge(b, ctx, "60", "0", "Grand Blitz Duel", false)

	return nil

}

func Bulletr(b *gotgbot.Bot, ctx *ext.Context) error {

	OpenChallenge(b, ctx, "60", "0", "Grand Blitz Duel", true)

	return nil

}

func Open(b *gotgbot.Bot, ctx *ext.Context) error {

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

	OpenChallenge(b, ctx, clockLimit, increment, "Open Challenge Duel", false)

	return nil

}
