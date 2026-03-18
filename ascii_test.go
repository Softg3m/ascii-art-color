package main

import (
	"bytes"
	"os"
	"testing"
)

// captureOutput temporarily redirects stdout so we can capture
// anything printed to the terminal during a test.
func captureOutput(f func()) string {
	// Save the current stdout
	old := os.Stdout

	// Create a pipe to capture output
	r, w, _ := os.Pipe()

	// Redirect stdout to the pipe writer
	os.Stdout = w

	// Execute the function whose output we want to capture
	f()

	// Close the writer and restore stdout
	w.Close()
	os.Stdout = old

	// Read the captured output
	var buf bytes.Buffer
	buf.ReadFrom(r)

	return buf.String()
}

// TestColorFlag verifies that a valid color flag returns a color code
func TestColorFlag(t *testing.T) {
	color := GetColor("--color=red")

	// If no color code is returned, the test should fail
	if color == "" {
		t.Error("Expected valid color")
	}
}

// TestBlankStringNoColor checks behavior when the input is a space character
func TestBlankStringNoColor(t *testing.T) {
	font, _ := LoadBanner("standard.txt")

	output := captureOutput(func() {
		PrintAscii(" ", "", "", font)
	})

	// A space should still produce ASCII rows
	if output == "" {
		t.Error("Expected ASCII art for space character")
	}
}

// TestBlankStringWithColor ensures a space still prints ASCII rows
// when color is applied
func TestBlankStringWithColor(t *testing.T) {
	font, _ := LoadBanner("standard.txt")

	output := captureOutput(func() {
		PrintAscii(" ", "", "\033[31m", font)
	})

	if output == "" {
		t.Error("Expected ASCII art output for blank string with color")
	}
}