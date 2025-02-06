package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
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

	fmt.Println("a ...any")
	scanner := bufio.NewScanner(os.Stdin)
	for {

		fmt.Print("🧑: ")
		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		if strings.ToLower(input) == "" {
			break
		}
		fmt.Print("\n🤖: ")

		if run(llm, ctx, input) != nil {
			log.Fatalln(err)
		}
		fmt.Print("\n\n")
	}

}

func run(llm gollm.LLM, ctx context.Context, input string) error {

	prompt := gollm.NewPrompt(input)

	response, err := llm.Generate(ctx, prompt)
	if err != nil {
		log.Fatalf("Failed to generate response %v", err)
		return err
	}

	fmt.Printf("Response:\n%s\n", response)
	return nil
}
