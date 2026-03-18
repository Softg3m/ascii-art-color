package main

import (
	"fmt"
	"strings"
)

// PrintAscii renders ASCII art for the given text.
// text       → the input string to convert to ASCII art
// substring  → part of the text that should be colored
// color      → ANSI color code
// font       → map containing ASCII art rows for each character
func PrintAscii(text, substring, color string, font map[rune][]string) {

	// Split text by newline so multiple lines can be printed
	lines := strings.Split(text, "\\n")

	// Process each line separately
	for _, word := range lines {

		// If the line is empty, just print a blank line
		if word == "" {
			fmt.Println()
			continue
		}

		// Each ASCII character has 8 rows
		for row := 0; row < 8; row++ {

			// Iterate through each character in the word
			for i := 0; i < len(word); {

				// Check if the current position matches the substring
				if substring != "" &&
					i+len(substring) <= len(word) &&
					word[i:i+len(substring)] == substring {

					// If substring matches, print it with color
					for _, char := range substring {

						// Apply color before printing the character
						if color != "" {
							fmt.Print(color)
						}

						// Print the ASCII art row for this character
						fmt.Print(font[char][row])

						// Reset color after printing
						if color != "" {
							fmt.Print(reset)
						}
					}

					// Skip forward by the substring length
					i += len(substring)

				} else {

					// Convert the character to rune for map lookup
					char := rune(word[i])

					// If a color is provided but no substring,
					// color the entire word
					if color != "" && substring == "" {
						fmt.Print(color)
					}

					// Print the ASCII art row for the character
					fmt.Print(font[char][row])

					// Reset color after printing
					if color != "" && substring == "" {
						fmt.Print(reset)
					}

					// Move to the next character
					i++
				}
			}

			// After printing all characters in this row,
			// move to the next ASCII row
			fmt.Println()
		}
	}
}