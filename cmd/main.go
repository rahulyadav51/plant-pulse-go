package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/genai"
)

func main() {
	app := fiber.New() // Initialize Fiber
	ctx := context.Background()

	// 1. Initialize Gemini Client
	apiKey := os.Getenv("GEMINI_API_KEY")
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	// 2. Define the Upload Route
	app.Post("/diagnose", func(c *fiber.Ctx) error {
		// Get the image file from the request
		file, err := c.FormFile("image")
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "No image uploaded"})
		}

		// Open the file and read it into memory
		src, _ := file.Open()
		defer src.Close()
		imgBytes, _ := io.ReadAll(src)

		// 3. Send Image to Gemini 2.0 Flash
		fmt.Println("Analyzing plant image...")
		result, err := client.Models.GenerateContent(ctx, "gemini-2.0-flash", 
			genai.Text("Identify this plant and provide a health diagnosis."),
			genai.Data("image/png", imgBytes), // Multi-modal support
			nil,
		)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "AI failed to process image"})
		}

		// 4. Return JSON response
		return c.JSON(fiber.Map{
			"diagnosis": result.Candidates[0].Content.Parts[0].Text,
		})
	})

	log.Fatal(app.Listen(":3000")) // Start server on port 3000
}