package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv" // Import for .env
	"google.golang.org/genai"
)

func main() {
	// 0. Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	app := fiber.New()
	ctx := context.Background()

	// 1. Initialize Gemini Client
	apiKey := os.Getenv("GEMINI_API_KEY")
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal("Failed to create Gemini client:", err)
	}

	// 2. Define the Upload Route
	app.Post("/diagnose", func(c *fiber.Ctx) error {
		file, err := c.FormFile("image")
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "No image uploaded"})
		}

		src, _ := file.Open()
		defer src.Close()
		imgBytes, _ := io.ReadAll(src)

		// 3. Prepare Multi-modal Content
		parts := []*genai.Part{
			genai.NewPartFromText("Identify this plant and provide a health diagnosis."),
			genai.NewPartFromBytes(imgBytes, "image/png"),
		}

		// Wrap parts into a Content slice with a Role
		contents := []*genai.Content{
			genai.NewContentFromParts(parts, genai.RoleUser),
		}

		// 4. Call Gemini 2.0 Flash
		result, err := client.Models.GenerateContent(ctx, "gemini-2.0-flash", contents, nil)
		if err != nil {
			fmt.Println("❌ AI Error:", err) // This will show you WHY it failed in the terminal
			return c.Status(500).JSON(fiber.Map{"error": "AI failed: " + err.Error()})
		}

		// 5. PRINT THE RESPONSE TO TERMINAL FOR DEBUGGING
		fmt.Println("✅ Diagnosis received!")
		fmt.Println("Response:", result.Candidates[0].Content.Parts[0].Text)

		// 6. Return JSON response to Postman
		return c.JSON(fiber.Map{
			"diagnosis": result.Candidates[0].Content.Parts[0].Text,
		})
	})

	log.Fatal(app.Listen(":3000"))
}
