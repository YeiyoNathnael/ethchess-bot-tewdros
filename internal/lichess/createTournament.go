package lichess

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func CreateArenaTournamentLink(tournamentDraft CreateArenaTournament) string {
	var link ArenaTournamentResponse
	_ = godotenv.Load()
	lichessAPI := os.Getenv("lichessAPI")
	if lichessAPI == "" {
		panic("lichessAPI is empty")
	}

	urL := "https://lichess.org/api/tournament"

	postData := url.Values{}
	postData.Set("name", tournamentDraft.Name)

	postData.Set("clockTime", strconv.FormatInt(tournamentDraft.ClockTime, 10))

	postData.Set("clockIncrement", strconv.FormatInt(tournamentDraft.ClockIncrement, 10))

	postData.Set("minutes", strconv.FormatInt(tournamentDraft.Minutes, 10))

	postData.Set("startDate", strconv.FormatInt(tournamentDraft.StartDate, 10))

	postData.Set("variant", tournamentDraft.Variant)

	postData.Set("rated", strconv.FormatBool(tournamentDraft.Rated))

	postData.Set("position", tournamentDraft.Position)

	postData.Set("berserkable", strconv.FormatBool(tournamentDraft.Berserkable))

	postData.Set("streakable", strconv.FormatBool(tournamentDraft.Streakable))

	postData.Set("hasChat", strconv.FormatBool(tournamentDraft.HasChat))

	postData.Set("description", tournamentDraft.Description)

	postData.Set("password", "")

	postData.Set("teamBattleByTeam", "")

	postData.Set("conditions.teamMember.teamId", tournamentDraft.Conditions.teamMember.teamId)

	postData.Set("conditions.minRating.rating", strconv.FormatInt(tournamentDraft.Conditions.minRating.toInt(), 10))

	postData.Set("conditions.maxRating.rating", strconv.FormatInt(tournamentDraft.Conditions.maxRating.toInt(), 10))

	postData.Set("conditions.nbRatedGame.nb", strconv.FormatInt(tournamentDraft.Conditions.nbRatedGame.toInt(), 10))

	postData.Set("conditions.allowList", tournamentDraft.Conditions.allowList)

	postData.Set("conditions.bots", strconv.FormatBool(tournamentDraft.Conditions.bots))

	postData.Set("conditions.accountAge", strconv.FormatInt(tournamentDraft.Conditions.accountAge, 10))

	req, _ := http.NewRequest("POST", urL, strings.NewReader(postData.Encode()))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", lichessAPI))

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	json.Unmarshal(body, &link)
	arenaTournamentLink := fmt.Sprintf("lichess.org/tournament/%v", link.Id)

	return arenaTournamentLink

}
