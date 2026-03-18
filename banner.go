package main

import (
	"os"
	"strings"
)

// LoadBanner loads an ASCII font file and converts it
// into a map where each rune maps to its ASCII art rows.
func LoadBanner(file string) (map[rune][]string, error) {

	// Read the entire banner file
	data, err := os.ReadFile(file)

	// If file reading fails, return the error
	if err != nil {
		return nil, err
	}

	// Split the file content into lines
	lines := strings.Split(string(data), "\n")

	// Create a map to store ASCII art for characters
	font := make(map[rune][]string)

	// ASCII printable characters start from 32 (space)
	ascii := 32

	// Index tracks position in the file
	index := 0

	// Loop through ASCII characters from 32 to 126
	for ascii <= 126 {

		// Each character uses 8 lines in the banner file
		font[rune(ascii)] = lines[index : index+8]

		// Move index to the next character block
		index += 9

		ascii++
	}

	return font, nil
}