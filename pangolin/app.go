package main

import (
	"context"
	"fmt"
	"os"
	"flag"
	"encoding/base64"
	"encoding/json"
	"google.golang.org/genai"
	"github.com/joho/godotenv"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) ReadFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(data)
}

func (a *App) GetGeminiResponse() {
	err := godotenv.Load()
  if err != nil {
    return
  }
	ctx := context.Background()
	apiKey := os.Getenv("GEMINI_KEY")
	fmt.Println(apiKey)
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
    APIKey:   apiKey,
    Backend:  genai.BackendGeminiAPI,
	})

	if err != nil {
		return
	}

	var config *genai.GenerateContentConfig = nil
	var model = flag.String("model", "gemini-1.5-pro-002", "the model name, e.g. gemini-1.5-pro-002")
	// Call the GenerateContent method.
	result, err := client.Models.GenerateContent(ctx, *model, genai.Text("What is your name?"), config)
	if err != nil {
		return
	}
	response, err := json.MarshalIndent(*result, "", "  ")
	if err != nil {
		return
	}
	// Log the output.
	fmt.Println(string(response))
}
