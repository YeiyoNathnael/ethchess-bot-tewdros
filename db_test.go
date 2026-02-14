package main

import (
	"context"
	"database/sql"
	"github.com/YeiyoNathnael/ethchess-bot-tewdros/internal/db"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"testing"
)

func TestTursoConnection(t *testing.T) {
	// 1. Define your Turso connection string
	// Format: libsql://<db-name>-<org>.turso.io?authToken=<your-long-jwt-token>
	dbUrl := "libsql://ethchess-yeiyonathnael.aws-ap-south-1.turso.io?authToken=eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE3NzEwMzEwMTgsImlkIjoiODRmNWQ4MGQtYWZjZi00OGMxLWI0ZTctYjA5NDFiYzNjZDM1IiwicmlkIjoiYTgwNmY1MWQtMTNkMy00MmEzLTlkY2MtYjlkNDI3YzEzMDJjIn0.hJ90AFNoRDSi6EC020l835tX_RMDP0pT77IvCm01-C0G7Qn4zFKJuBHAa2gIKjEaKeESLRpcUPnYvoc90pWNDA"

	// 2. Open the connection
	database, err := sql.Open("libsql", dbUrl)
	if err != nil {
		t.Fatalf("Failed to open connection: %v", err)
	}
	defer database.Close()

	// 3. Verify the connection is actually alive
	if err := database.Ping(); err != nil {
		t.Fatalf("Failed to ping database: %v", err)
	}

	// 4. Test a simple read-only query to ensure sqlc mapping works
	queries := db.New(database)
	authors, err := queries.GetUser(context.Background(), "5353233678")
	if err != nil {
		t.Fatalf("Failed to fetch authors: %v", err)
	}
	author_id := authors.TelegramID
	t.Logf("Success! Connected to Turso and found authors. %s", author_id)
}
