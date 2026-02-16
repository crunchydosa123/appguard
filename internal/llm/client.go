package llm

import (
	"context"
	"os"

	"google.golang.org/genai"
)

func NewClient(ctx context.Context) (*genai.Client, error) {
	return genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: os.Getenv("GEMINI_API_KEY"),
	})
}
