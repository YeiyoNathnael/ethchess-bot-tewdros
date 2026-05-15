package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/YeiyoNathnael/ethchess-bot-tewdros/internal/gemini"
	"github.com/YeiyoNathnael/ethchess-bot-tewdros/internal/lichess"
	"github.com/joho/godotenv"
	"google.golang.org/genai"
)

// var botID int = 7720642643
var botUserName string = "@ETHCHESSSupportbot"
var history *genai.Chat


func main() {

	// Get token from the environment variable
	err := godotenv.Load()
	token := os.Getenv("TOKEN")
	if token == "" {
		panic("TOKEN environment variable is empty")
	}

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
	dispatcher.AddHandler(handlers.NewCommand("blitz", lichess.Blitz))
	dispatcher.AddHandler(handlers.NewCommand("blitzr", lichess.Blitzr))
	dispatcher.AddHandler(handlers.NewCommand("bullet", lichess.Bullet))
	dispatcher.AddHandler(handlers.NewCommand("bulletr", lichess.Bulletr))

	dispatcher.AddHandler(handlers.NewCommand("user", getLichessRating))
	dispatcher.AddHandler(handlers.NewCommand("open", lichess.Open))

	dispatcher.AddHandler(handlers.NewMessage(func(msg *gotgbot.Message) bool {
		for _, e := range msg.Entities {
			if e.Type == "mention" {
				mentioned := msg.Text[e.Offset : e.Offset+e.Length]
				if mentioned == botUserName {
					return true
				}
			}
		}
		return (msg.ReplyToMessage != nil || msg.NewChatMembers != nil)
	}, chat))
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

// NOTE: This is a really random function, not really there to serve a purpose
func getLichessRating(b *gotgbot.Bot, ctx *ext.Context) error {

	username := ctx.Args()

	if len(username) < 2 {

		_, _ = ctx.EffectiveMessage.Reply(b, "PLease PLease provide a proper username", nil)

		return nil

	}

	user := username[1]
	userRating := lichess.GetLichessUser(user)
	_, _ = ctx.EffectiveMessage.Reply(b, "User Bullet rating is: "+strconv.FormatInt(userRating, 10), &gotgbot.SendMessageOpts{
		ParseMode: "HTML",
	})
	return nil
}

func chat(b *gotgbot.Bot, ctx *ext.Context) error {

		systemInstruction := &genai.Content{
    	Role: "user",
	    Parts: []*genai.Part{
        {
				Text: `You are Tewodros (Teddy), the official support bot of EthChess — Ethiopia's fastest-growing chess community, based in Addis Ababa.

Your role is to assist club members, newcomers, and chess enthusiasts with anything related to EthChess: events, membership, tournaments, club info, and general chess questions.

Guidelines:
- Be friendly, warm, and community-oriented
- Keep responses brief and clear
- When you don't know something specific about the club, say so honestly and suggest they reach out to the club directly
- You may use chess analogies or light humor when appropriate
- Always represent EthChess positively and professionally`,
				},
  	  },
		}
	msg := ctx.EffectiveMessage
	if history == nil {
		history = &genai.Chat{}
	}
	for _, e := range msg.NewChatMembers {

		joinedUser := e.Username
		systemInstructionNewJoiningUser := &genai.Content{
    	Role: "user",
	    Parts: []*genai.Part{
        {
					Text: "a brief welcome message for user who just joined our chess club telegram group called ethchess. make it only 2 sentences, very warm and breif as well.ethchess is a chess club found in Ethiopia and the fastest growing chess community in ethiopia.  only send me the welcome message nothing else. the user's name is"+joinedUser,
				},
  	  },
		}

		geminiResponse, chat := gemini.GeminiResponse("",gemini.Gemma_4_31b.String(), history,systemInstructionNewJoiningUser)

		_, err := msg.Reply(b, geminiResponse, &gotgbot.SendMessageOpts{
			ParseMode: "MarkdownV2",
		},
		)
		if err != nil {
			return fmt.Errorf("failed to send source: %w", err)
		}
		history = chat

	}
	// Check if this is a reply and if it's replying to the bot
	if msg.ReplyToMessage != nil && msg.ReplyToMessage.From != nil && msg.ReplyToMessage.From.Id == b.Id {

		//TODO: room for improvement on the hardcoded prompt :)
		reply, chat := gemini.GeminiResponse(msg.Text, gemini.Gemma_4_31b.String(), history,systemInstruction)
		_, err := msg.Reply(b, reply, &gotgbot.SendMessageOpts{
			ParseMode: "MarkdownV2",
		},
		)
		if err != nil {
			return fmt.Errorf("failed to send source: %w", err)
		}
		history = chat
	}
	for _, e := range msg.Entities {
		if e.Type == "mention" {
			mentioned := msg.Text[e.Offset : e.Offset+e.Length]
			if mentioned == botUserName {
				reply, chat := gemini.GeminiResponse(msg.Text, gemini.Gemma_4_31b.String(), history,systemInstruction)
				_, err := msg.Reply(b, reply, &gotgbot.SendMessageOpts{
					ParseMode: "MarkdownV2",
				},
				)
				if err != nil {
					return fmt.Errorf("failed to send source: %w", err)
				}
				history = chat

			}

		}

	}

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
