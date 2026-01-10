package gemini

import (
	"context"
	"encoding/json"
	"github.com/joho/godotenv"
	"google.golang.org/genai"
	"log"
	"os"
)

// Your Google API key

func GeminiResponse(userRequest string) string {

	//	var GeminiRes struct {
	//		Url string `json:"url"`
	//	}

	//if err := json.Unmarshal(body, &LichessChallengeResponse); err != nil {
	//		return fmt.Errorf("parsing failed: %w", err)
	//	}

	type Part struct {
		Text string `json:"text"`
		Role string `json:"role"`
	}

	type Content struct {
		Parts []Part `json:"parts"`
		Role  string `json:"role"`
	}

	type Candidate struct {
		Content      Content `json:"content"`
		FinishReason string  `json:"finishReason"`
	}
	var GeminiRes struct {
		Url string `json:"url"`

		Candidates []Candidate `json:"candidates"`
	}

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

	chat, err := client.Chats.Create(ctx, "gemma-3-27b", nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	result, err := chat.SendMessage(ctx, genai.Part{Text: userRequest})
	if err != nil {
		log.Fatal(err)
	}

	geminiRes := debugPrint(result)

	json.Unmarshal(geminiRes, &GeminiRes)
	return GeminiRes.Candidates[0].Content.Parts[0].Text

}

func debugPrint[T any](r *T) []byte {

	response, err := json.Marshal(*r)
	if err != nil {
		log.Fatal(err)
	}

	return response
}
