package tournament

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/YeiyoNathnael/ethchess-bot-tewdros/internal/db"
	"github.com/google/uuid"
	"os"
	"strconv"
	"time"
)

func CreateTournamentDraft(b *gotgbot.Bot, ctx *ext.Context) error {

	args := ctx.Args()
	msg := ctx.EffectiveMessage
	if len(args) < 2 {

		return fmt.Errorf("Please thats not enough amount")
	}

	Name := args[1]
	Format := args[2]
	TimeControl := args[3]

	StartDatetime := args[4]
	EndDatetime := args[5]
	size := args[6]
	Capacity, _ := strconv.Atoi(size)
	Description := args[7]
	Prizes := args[8]

	tournamentDraftId := uuid.New().String()
	tournamentDraft := db.TournamentDraft{
		DraftUuid: tournamentDraftId,
		TournamentClass: sql.NullString{
			String: " Test for now",
			Valid:  true,
		},
		RequestedBy: sql.NullString{
			String: msg.SenderChat.FirstName,
			Valid:  true,
		},
		Name: sql.NullString{
			String: Name,
			Valid:  true,
		},

		Format: sql.NullString{
			String: Format,
			Valid:  true,
		},

		TimeControl: sql.NullString{
			String: TimeControl,
			Valid:  true,
		},
		StartDatetime: sql.NullString{
			String: StartDatetime,
			Valid:  true,
		},

		EndDatetime: sql.NullString{
			String: EndDatetime,
			Valid:  true,
		},

		Timezone: sql.NullString{
			String: "GMT+3",
			Valid:  true,
		},
		Capacity: sql.NullInt64{
			Int64: int64(Capacity),
			Valid: true,
		},
		Visibility: sql.NullString{
			String: "Public",
			Valid:  true,
		},
		Prizes: sql.NullString{
			String: Prizes,
			Valid:  true,
		},
		OrganizerTelegramID: sql.NullString{
			String: strconv.FormatInt(msg.SenderChat.Id, 10),
			Valid:  true,
		},
		Rationale: sql.NullString{
			String: Description,
			Valid:  true,
		},
		CreatedAt: sql.NullString{
			String: time.Now().String(),
			Valid:  true,
		},
	}

	err := draft(tournamentDraft)

	if err != nil {
		return err
	}
	return nil

}

func draft(tournamentDraft db.TournamentDraft) error {

	contx := context.Background()
	dbUrl := os.Getenv("DBURL")

	//////////////////////////////////////
	database, err := db.Init(dbUrl)
	if err != nil {
		return err
	}
	defer database.Close()
	///////////////////////////////////////

	newtournamentDraft := db.CreateDraftParams{

		DraftUuid:           tournamentDraft.DraftUuid,
		TournamentClass:     tournamentDraft.TournamentClass,
		RequestedBy:         tournamentDraft.RequestedBy, // should autopopulate using name of org.tg.id
		Name:                tournamentDraft.Name,
		Format:              tournamentDraft.Format,
		TimeControl:         tournamentDraft.TimeControl,
		StartDatetime:       tournamentDraft.StartDatetime,
		EndDatetime:         tournamentDraft.EndDatetime,
		Timezone:            tournamentDraft.Timezone,
		Capacity:            tournamentDraft.Capacity,
		Visibility:          tournamentDraft.Visibility,
		Prizes:              tournamentDraft.Prizes,
		OrganizerTelegramID: tournamentDraft.OrganizerTelegramID, // should autopopulate
		Rationale:           tournamentDraft.Rationale,
		CreatedAt:           tournamentDraft.CreatedAt, // should autopopulate

	}

	queries := db.New(database)

	err = queries.CreateDraft(contx, newtournamentDraft)
	if err != nil {
		return err
	}

	return nil
}
