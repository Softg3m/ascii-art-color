package main

import (
	"fmt"
	"strings"
)

// colors maps user-provided color values to their corresponding
// ANSI escape codes used by the terminal.
//
// The program supports multiple color formats:
//   1. Named ANSI colors (red, blue, green, etc.)
//   2. HEX color notation (#RRGGBB)
//   3. RGB notation (rgb(r, g, b))
//   4. HSL notation (hsl(h, s%, l%))
//
// All supported formats map to the closest ANSI terminal color.
var colors = map[string]string{

	// Standard ANSI color names
	"black":  "\033[30m",
	"red":    "\033[31m",
	"green":  "\033[32m",
	"yellow": "\033[33m",
	"blue":   "\033[34m",
	"purple": "\033[35m",
	"cyan":   "\033[36m",
	"white":  "\033[37m",
	"orange": "\033[38;5;208m",

	// HEX color equivalents
	"#ff0000": "\033[31m",
	"#00ff00": "\033[32m",
	"#ffff00": "\033[33m",
	"#0000ff": "\033[34m",
	"#800080": "\033[35m",
	"#00ffff": "\033[36m",
	"#ffffff": "\033[37m",
	"#ffa500": "\033[38;5;208m",

	// RGB color notation
	"rgb(0, 0, 0)":       "\033[30m",
	"rgb(255, 0, 0)":     "\033[31m",
	"rgb(0, 255, 0)":     "\033[32m",
	"rgb(255, 255, 0)":   "\033[33m",
	"rgb(0, 0, 255)":     "\033[34m",
	"rgb(255, 0, 255)":   "\033[35m",
	"rgb(0, 255, 255)":   "\033[36m",
	"rgb(255, 255, 255)": "\033[37m",
	"rgb(255, 165, 0)":   "\033[38;5;208m",

	// HSL color notation
	"hsl(0, 0%, 0%)":     "\033[30m",
	"hsl(0, 100%, 50%)":  "\033[31m",
	"hsl(120, 100%, 50%)":"\033[32m",
	"hsl(60, 100%, 50%)": "\033[33m",
	"hsl(240, 100%, 25%)":"\033[34m",
	"hsl(300, 100%, 25%)":"\033[35m",
	"hsl(180, 100%, 50%)":"\033[36m",
	"hsl(0, 0%, 100%)":   "\033[37m",
	"hsl(39, 100%, 50%)": "\033[38;5;208m",
}

// reset is the ANSI escape code that returns the terminal color
// back to the default after colored text has been printed.
const reset = "\033[0m"

// IsColorFlag checks whether the provided CLI argument
// starts with the expected color flag prefix "--color=".
func IsColorFlag(flag string) bool {
	return strings.HasPrefix(flag, "--color=")
}

// GetColor extracts the color value from the flag
// and returns the corresponding ANSI escape code.
//
// Example:
//   --color=red        → "\033[31m"
//   --color=#FF0000    → "\033[31m"
//   --color=rgb(255,0,0) → "\033[31m"
//
// If the color is not supported, the function returns an empty string.
func GetColor(flag string) string {

	// Remove the "--color=" prefix from the flag
	// Both uppercase and lowercase works
	name := strings.ToLower(strings.TrimPrefix(flag, "--color="))

	// Look up the color in the map
	color, ok := colors[name]

	// If color is not found, return empty string
	if !ok {
		return ""
	}

	return color
}

func Usage2(){
	fmt.Println("Color does not exist in the map, Try a different color.")
	return
}