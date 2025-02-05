package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/teilomillet/gollm"
	"github.com/teilomillet/gollm/utils"
)

func main() {

	apikey := os.Getenv("OPENAI_API_KEY")
	if apikey == "" {
		log.Fatal("api key not found")
	}

	llm, err := gollm.NewLLM(
		gollm.SetProvider("openai"),
		gollm.SetModel("gpt-4o-mini"),
		gollm.SetAPIKey(apikey),
		gollm.SetMaxTokens(1000),
		gollm.SetMaxRetries(2),
		gollm.SetRetryDelay(time.Second*2),
		gollm.SetLogLevel(utils.LogLevelInfo),
	)

	if err != nil {
		log.Fatalf("Failed to create llm: %v", err)
	}

	ctx := context.Background()

	prompt := gollm.NewPrompt("what is golang?")

	response, err := llm.Generate(ctx, prompt)
	if err != nil {
		log.Fatalf("Failed to generate response %v", err)
	}

	fmt.Printf("Response:\n%s\n", response)

}
