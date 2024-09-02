package remini

import (
	"context"
	"fmt"
	"github.com/google/generative-ai-go/genai"
	"learnNaturally/commands"
	"log"
)

func ChatAndProofread(message string, proxyUrl string) {
	ctx := context.Background()
	client, err := commands.MyGeminiClient(ctx, proxyUrl)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()
	model := client.GenerativeModel("gemini-1.5-flash")
	session := model.StartChat()
	session.History = []*genai.Content{}
	resp, err := session.SendMessage(ctx, genai.Text(message))
	if err != nil {
		log.Fatalf("Error sending message: %v", err)
	}

	for _, part := range resp.Candidates[0].Content.Parts {
		fmt.Printf("%v\n", part)
	}
}
