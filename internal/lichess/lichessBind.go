package lichess

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
	"time"

	tgmd "github.com/Mad-Pixels/goldmark-tgmd"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/YeiyoNathnael/ethchess-bot-tewdros/internal/db"
	"github.com/YeiyoNathnael/ethchess-bot-tewdros/internal/gemini"
	"github.com/google/uuid"
)

func LichessBind(b *gotgbot.Bot, ctx *ext.Context) error {

	user := ctx.EffectiveUser

	stateToken := base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(int(user.Id)) + ":" + uuid.New().String()))

	bindLink := fmt.Sprintf("Click the link below to connect your lichess account: https://ethchess-website.vercel.app/telegram-link?state={%v}", stateToken)

	//NOTE: so here there might be several calls for bindLink, like many /links
	// so how to fix it?

	var buf bytes.Buffer
	md := tgmd.TGMD()

	err := md.Convert([]byte(bindLink), &buf)
	if err != nil {
		panic(err)
	}

	_, err = ctx.EffectiveMessage.Reply(b, buf.String(), &gotgbot.SendMessageOpts{
		ParseMode: "MarkdownV2",
	},
	)
	if err != nil {
		return fmt.Errorf("failed to send source: %w", err)
	}

	return nil

}

// NOTE: UNtil the website is made ill not add any check that will make this avoid working manually
func Auth_Success(b *gotgbot.Bot, ctx *ext.Context) error {

	dbUrl := os.Getenv("DBURL")
	contxt := context.Background()
	auth_state := ctx.Args()

	//FIX: Obviously needs better handling
	if len(auth_state) < 3 {
		return nil
	}

	stateToken := auth_state[1]
	lichessUsername := auth_state[2]

	telegramId, err := decodeTelegramId(stateToken)

	if err != nil {
		return err
	}

	authenticatedUser := db.CreateUserParams{
		TelegramID: telegramId,
		LichessUsername: sql.NullString{
			String: lichessUsername,
			Valid:  true,
		},
		CreatedAt: sql.NullString{
			String: time.Now().Format(time.RFC3339),
			Valid:  true,
		},
	}

	database, err := db.Init(dbUrl)

	if err != nil {
		return err
	}

	defer database.Close()
	queries := db.New(database)
	err = queries.CreateUser(contxt, authenticatedUser)

	if err != nil {

		simplify_msg_prompt := fmt.Sprintf(
			"You are a helpful support assistant. Translate the following technical error into a single, plain-English sentence for a non-technical user. "+
				"Focus on the 'what' (the action that failed) rather than the 'why' (the code reason). "+
				"Example: Instead of 'Unique constraint violation', say 'That username is already taken.' "+
				"Error to translate: %v", err.Error())

		simple_err, _ := gemini.GeminiResponse(simplify_msg_prompt, gemini.Gemma_3_27b.String(), nil)

		ctx.EffectiveMessage.Reply(b, simple_err, &gotgbot.SendMessageOpts{
			ParseMode: "MarkdownV2",
		})
		return err
	}

	successMessage := fmt.Sprintf("Successfully linked to Lichess account: %v", lichessUsername)
	_, err = ctx.EffectiveMessage.Reply(b, successMessage, &gotgbot.SendMessageOpts{
		ParseMode: "MarkdownV2",
	})

	return nil
}

func decodeTelegramId(encodedStateToken string) (string, error) {

	stateTokenBytes, err := base64.StdEncoding.DecodeString(encodedStateToken)
	if err != nil {
		return "", err
	}

	stateToken := string(stateTokenBytes)

	for i := 0; i < len(stateToken); i++ {

		if stateToken[i] == ':' {
			return stateToken[0:i], nil
		}
	}

	return "", fmt.Errorf("invalid state token format")

}
