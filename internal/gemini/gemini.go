package gemini

import (
	"bytes"
	"context"
	"encoding/json"
	tgmd "github.com/Mad-Pixels/goldmark-tgmd"
	"github.com/joho/godotenv"
	"google.golang.org/genai"
	"log"
	"os"
	"strings"
)

// Your Google API key

type GeminiModels int

const (
	Gemma_4_31b GeminiModels = iota
	Gemma_4_26_A4B
)

// A map to store the string representation of each state
var Models = map[GeminiModels]string{
	Gemma_4_31b:    "gemma-4-31b-it",
	Gemma_4_26_A4B: "gemma-4-26b-a4b-it",
}

// String implements the fmt.Stringer interface
func (n GeminiModels) String() string {
	return Models[n]
}

func GeminiResponse(userRequest string, model string, chatt *genai.Chat, systemInstruction *genai.Content) (string, *genai.Chat) {

	GoogleSearch := &genai.GoogleSearch{
		SearchTypes: &genai.SearchTypes{
			WebSearch: &genai.WebSearch{},
			//	ImageSearch: &genai.ImageSearch{},
		},
	}
	Tools := &genai.Tool{
		GoogleSearch: GoogleSearch,
	}
	type Part *genai.Part
	type Content *genai.Content
	type Candidate *genai.Candidate

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

	chat, err := client.Chats.Create(ctx, model, &genai.GenerateContentConfig{
		Tools:             []*genai.Tool{Tools},
		SystemInstruction: systemInstruction,
	}, chatt.History(true))
	if err != nil {
		log.Fatal(err)
	}

	result, err := chat.SendMessage(ctx, genai.Part{Text: userRequest})
	if err != nil {
		return "Sorry, I'm having trouble connecting right now Please try again in a moment", chat
	}
	geminiRes := debugPrint(result)

	json.Unmarshal(geminiRes, &GeminiRes)

	response := extractResponseText(GeminiRes.Candidates[0].Content.Parts)

	var buf bytes.Buffer
	md := tgmd.TGMD()

	err = md.Convert([]byte(response), &buf)
	if err != nil {
		panic(err)
	}

	return buf.String(), chat

}

func debugPrint[T any](r *T) []byte {

	response, err := json.Marshal(*r)
	if err != nil {
		log.Fatal(err)
	}

	return response
}
func extractResponseText(parts []*genai.Part) string {
	var sb strings.Builder
	for _, part := range parts {
		if part.Thought {
			continue // skip reasoning/thinking parts
		}
		if part.Text != "" {
			sb.WriteString(part.Text)
		}
	}
	return sb.String()
}
