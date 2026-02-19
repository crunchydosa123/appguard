package main

import (
	"appguard/cmd"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found (using system env)")
	}

	print(os.Getenv("GEMINI_API_KEY"))
	// Optional: sanity check
	if os.Getenv("GEMINI_API_KEY") == "" {
		log.Println("Warning: GOOGLE_API_KEY not set")
	}
	cmd.Execute()
}
