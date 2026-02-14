package lichess

import (
	"encoding/base64"
	"fmt"
	"strconv"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/google/uuid"
)

func LichessBind(b *gotgbot.Bot, ctx *ext.Context) string {

	user := ctx.EffectiveUser

	stateToken := base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(int(user.Id)) + ":" + uuid.New().String()))

	bindLink := fmt.Sprintf("https://ethchess-website.vercel.app/telegram-link?state={%v}", stateToken)

	return bindLink
}
