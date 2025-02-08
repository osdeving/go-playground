//go:build ignore

package main

import (
	"context"
	"fmt"
	"io"

	openai "github.com/sashabaranov/go-openai"
)

func main() {
	config := openai.DefaultConfig("xxx") // required, but unused
	config.BaseURL = "http://localhost:11434/v1"
	client := openai.NewClientWithConfig(config)

	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "You are a helpful assistant.",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: "Who won the world series in 2020?",
		},
		{
			Role:    openai.ChatMessageRoleAssistant,
			Content: "The LA Dodgers won in 2020.",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: "Where was it played?",
		},
	}

	stream, err := client.CreateChatCompletionStream(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    "deepseek-r1:7b",
			Messages: messages,
		},
	)
	if err != nil {
		fmt.Printf("Stream error: %v\n", err)
		return
	}
	defer stream.Close()

	// Print the streaming response as it arrives
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			fmt.Println() // add newline at the end
			return
		}
		if err != nil {
			fmt.Printf("Stream error: %v\n", err)
			return
		}
		fmt.Print(response.Choices[0].Delta.Content)
	}
}
