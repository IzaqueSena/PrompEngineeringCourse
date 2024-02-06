package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

const CONFIG_PATH = "../deploy/config/dev.json"

func main() {
	// environment config
	envConfig()

	// openai key
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	// basic example

	// create the content
	content := "The sky is"

	response, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: "gpt-3.5-turbo",
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: content,
				},
			},
			MaxTokens:        256,
			Temperature:      0.7,
			TopP:             1,
			FrequencyPenalty: 0,
			PresencePenalty:  0,
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(response)
}

func envConfig() {
	// Open the config file
	file, err := os.Open(CONFIG_PATH)
	if err != nil {
		log.Fatalf("Error opening config file: %v", err)
	}
	defer file.Close()

	// Decode the JSON data into a map[string]string
	var configData map[string]string
	if err := json.NewDecoder(file).Decode(&configData); err != nil {
		log.Fatalf("Error decoding config file: %v", err)
	}

	// Set environment variables based on the config data
	for key, value := range configData {
		if err := os.Setenv(key, value); err != nil {
			log.Fatalf("Error setting environment variable: %v", err)
		}
	}
}
