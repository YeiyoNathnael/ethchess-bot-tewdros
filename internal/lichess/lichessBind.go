package lichess

import (
	"bytes"
	"encoding/base64"
	"fmt"
	tgmd "github.com/Mad-Pixels/goldmark-tgmd"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/google/uuid"
	"strconv"
)

func LichessBind(b *gotgbot.Bot, ctx *ext.Context) error {

	user := ctx.EffectiveUser

	stateToken := base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(int(user.Id)) + ":" + uuid.New().String()))

	bindLink := fmt.Sprintf("Click the link below to connect your lichess account: https://ethchess-website.vercel.app/telegram-link?state={%v}", stateToken)
	var buf bytes.Buffer
	md := tgmd.TGMD()

	err := md.Convert([]byte(bindLink), &buf)
	if err != nil {
		panic(err)
	}

	_, err = ctx.EffectiveMessage.Reply(b, bindLink, &gotgbot.SendMessageOpts{
		ParseMode: "MarkdownV2",
	},
	)
	if err != nil {
		return fmt.Errorf("failed to send source: %w", err)
	}

	return nil

}
