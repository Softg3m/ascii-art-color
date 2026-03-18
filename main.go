package main

import (
	"fmt"
	"os"
)

// main is the entry point of the program
func main() {

	// Retrieve command line arguments (excluding program name)
	args := os.Args[1:]

	// If no arguments are provided, display usage instructions
	if len(args) == 0 {
		Usage()
		return
	}

	// Initialize variables
	color := ""
	substring := ""
	text := ""
	banner := "standard.txt"

	// Check if the first argument is a color flag
	if IsColorFlag(args[0]) {

		// Extract the ANSI color code
		color = GetColor(args[0])

		// If color is invalid, show usage instructions
		if color == "" {
			Usage2()
			return
		}

		// Determine argument pattern based on argument count
		switch len(args) {

		case 2:
			// Example: go run . --color=red "Hello"
			text = args[1]

		case 3:
			// Could be either:
    		// --color=red "Hello" shadow
    		// --color=red lo "Hello"
			if args[2] == "standard" || args[2] == "shadow" || args[2] == "thinkertoy" {
				text = args[1]
				banner = args[2] + ".txt"
			} else {
				substring = args[1]
				text = args[2]
			}

		case 4:
			// Example: go run . --color=red lo "Hello" shadow
			substring = args[1]
			text = args[2]
			banner = args[3] + ".txt"

		default:
			// Invalid number of arguments
			Usage()
			return
		}

	} else {

		// No color flag provided
		switch len(args) {

		case 1:
			// Example: go run . "Hello"
			text = args[0]

		case 2:
			// Example: go run . "Hello" shadow
			text = args[0]
			banner = args[1] + ".txt"

		default:
			Usage()
			return
		}
	}

	// Load the ASCII banner font
	font, err := LoadBanner(banner)

	// If banner loading fails, print error
	if err != nil {
		fmt.Println(err)
		return
	}

	// Render the ASCII art
	PrintAscii(text, substring, color, font)
}

// Usage prints the correct command format
func Usage() {
	fmt.Println("Usage: go run . [OPTION] [STRING]")
	fmt.Println()
	fmt.Println(`EX: go run . --color=<color> <substring> "something"`)
}