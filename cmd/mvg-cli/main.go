package main

import (
	"fmt"
	"mvg-fontifier/internal/transformer"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: mvg-fontifier <text>")
		fmt.Println("Example: mvg-fontifier \"Hello World\"")
		os.Exit(1)
	}

	input := strings.Join(os.Args[1:], " ")
	output := transformer.Transform(input)
	fmt.Println(output)
}
