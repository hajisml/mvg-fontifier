package main

import (
	"fmt"
	"log"
	"mvg-fontifier/internal/transformer"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		// No arguments provided, launch TUI
		if err := runTUI(); err != nil {
			log.Fatalf("Error running TUI: %v", err)
		}
		return
	}

	// Handle special flags
	arg := os.Args[1]
	if arg == "-h" || arg == "--help" {
		fmt.Println("mvg-fontifier - A tool for text stylization")
		fmt.Println("\nUsage:")
		fmt.Println("  mvg-fontifier          Launch interactive TUI")
		fmt.Println("  mvg-fontifier <text>   Transform text and output to stdout")
		fmt.Println("\nOptions:")
		fmt.Println("  -h, --help             Show this help message")
		return
	}

	// Transform arguments
	input := strings.Join(os.Args[1:], " ")
	output := transformer.Transform(input)
	fmt.Println(output)
}
