package gemini

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"google.golang.org/genai"
	"log"
	"os"
)

// Your Google API key

func GeminiResponse() string {

	err := godotenv.Load()
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		panic("TOKEN environment variable is empty")
	}
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	chat, err := client.Chats.Create(ctx, "gemini-2.5-flash", nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	result, err := chat.SendMessage(ctx, genai.Part{Text: "What's the weather in New York?"})
	if err != nil {
		log.Fatal(err)
	}
	debugPrint(result)

	result, err = chat.SendMessage(ctx, genai.Part{Text: "How about San Francisco?"})
	if err != nil {
		log.Fatal(err)
	}
	response, err := json.MarshalIndent(result, "", "  ")
	return string(response)
}

func debugPrint[T any](r *T) {

	response, err := json.MarshalIndent(*r, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(response))
}
