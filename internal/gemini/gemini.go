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

type GeminiModels int

const (
	Gemini_2_5_flash GeminiModels = iota
	Gemini_2_5_flash_lite
	Gemini_3_flash
	Gemma_3_12b
	Gemma_3_27b
)

// A map to store the string representation of each state
var Models = map[GeminiModels]string{
	Gemini_2_5_flash:      "gemini-2.5-flash",
	Gemini_2_5_flash_lite: "gemini-2.5-flash-lite",
	Gemini_3_flash:        "gemini-3-flash-preview",
	Gemma_3_12b:           "gemma-3-12b-it",
	Gemma_3_27b:           "gemma-3-27b-it",
}

// String implements the fmt.Stringer interface
func (n GeminiModels) String() string {
	return Models[n]
}

func GeminiResponse(userRequest string, model string, chatt *genai.Chat) (string, *genai.Chat) {

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

	chat, err := client.Chats.Create(ctx, model, nil, chatt.History(true))
	if err != nil {
		log.Fatal(err)
	}

	result, err := chat.SendMessage(ctx, genai.Part{Text: userRequest})
	if err != nil {
		log.Fatal(err)
	}

	geminiRes := debugPrint(result)

	json.Unmarshal(geminiRes, &GeminiRes)
	return markdownToMarkdownV2(GeminiRes.Candidates[0].Content.Parts[0].Text), chat

}

func debugPrint[T any](r *T) []byte {

	response, err := json.Marshal(*r)
	if err != nil {
		log.Fatal(err)
	}

	return response
}
