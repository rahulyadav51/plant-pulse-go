package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/genai"
)

func main() {
	ctx := context.Background()

	// 1. Get API Key from your terminal environment
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("Error: GEMINI_API_KEY is not set. Get one at aistudio.google.com")
	}

	// 2. Initialize the Gemini Client
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatalf("Failed to create Gemini client: %v", err)
	}

	// 3. Define the plant identification prompt
	prompt := "Analyze this description: A plant with large, green, violin-shaped leaves. Identify the species and provide 3 quick care tips."

	fmt.Println("PlantPulse Go is identifying your plant...")

	// 4. Generate the diagnosis
	result, err := client.Models.GenerateContent(ctx, "gemini-2.0-flash", genai.Text(prompt), nil)
	if err != nil {
		log.Fatalf("AI Generation Error: %v", err)
	}

	// 5. Output the result
	fmt.Println("\n--- Diagnosis Result ---")
	fmt.Println(result.Candidates[0].Content.Parts[0].Text)
}