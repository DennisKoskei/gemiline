package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/denniskoskei/gemiline/internal/gemini"
	"github.com/denniskoskei/gemiline/internal/session"
	"github.com/denniskoskei/gemiline/internal/ui"
)

func main() {
	short := flag.Bool("s", false, "Short answer")
	persistent := flag.Bool("p", false, "Persistent mode")

	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		fmt.Println("Usage: gel [options] <query>")
		return
	}

	query := strings.Join(args, " ")
	apiKey := os.Getenv("GEMINI_API_KEY")

	if apiKey == "" {
		fmt.Println("ERROR: GEMINI_API_KEY not set")
		return
	}

	// Short Results mode
	if *short {
		query = "In a few words " + query
	}

	// Persistent mode
	if *persistent {
		s := session.NewSession(apiKey)
		s.Start(query)
		return
	}

	// Normal mode
	response, err := gemini.AskGemini(apiKey, query)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	ui.PrintResponse(response)
}
