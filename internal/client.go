package internal

import (
	"context"
	"fmt"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
	"github.com/Syedowais312/CLI-Girlfriend/model"
)
func QueryChatbotAPI(systemPrompt, userPrompt string) (string, error) {
	ctx := context.Background()

	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("GEMINI_API_KEY not set")
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return "", err
	}
	defer client.Close()

	genModel := client.GenerativeModel("gemini-2.0-flash")

	// SYSTEM PERSONA FIRST
	genModel.SystemInstruction = &genai.Content{
		Role: "system",
		Parts: []genai.Part{
			genai.Text(systemPrompt),
		},
	}

	// LOAD HISTORY
	history, _ := LoadHistory()

	chat := genModel.StartChat()

	// Convert local memory into Gemini roles
	for _, msg := range history.Messages {
		role := "user"
		if msg.Role == "assistant" {
			role = "model"
		}

		chat.History = append(chat.History, &genai.Content{
			Role:  role,
			Parts: []genai.Part{genai.Text(msg.Content)},
		})
	}

	// USER MESSAGE
	resp, err := chat.SendMessage(ctx, genai.Text(userPrompt))
	if err != nil {
		return "", err
	}

	// EXTRACT RESPONSE
	var output string
	for _, cand := range resp.Candidates {
		for _, part := range cand.Content.Parts {
			if text, ok := part.(genai.Text); ok {
				output += string(text)
			}
		}
	}

	// SAVE TO MEMORY
	history.Messages = append(
		history.Messages,
		model.ChatMessage{Role: "user", Content: userPrompt},
		model.ChatMessage{Role: "assistant", Content: output},
	)

	SaveHistory(history)

	return output, nil
}
