# ASCII Art Color (Go)

ASCII Art Color is a Go program that converts text into ASCII art using predefined banner fonts and allows coloring parts of the output using ANSI terminal colors.

The project extends the classic ASCII art renderer by adding color support through a command line flag.

## Features

* Convert text into ASCII art
* Support for multiple banner fonts
* Colorize the entire text or specific substrings
* Handle multiline input
* Unit tests for core functionality
* Clean modular Go code using only the standard library

## Supported Banners

The program supports three ASCII art banners:

* standard
* shadow
* thinkertoy

Example:

```
go run . "Hello" shadow
```

## Supported Colors

The following ANSI colors are supported:

* black
* red
* green
* yellow
* blue
* purple
* cyan
* white

Example:

```
go run . --color=red "Hello"
```

## Usage

```
go run . [OPTION] [STRING]
```

Example:

```
go run . --color=<color> <substring> "something"
```

### Examples

Print ASCII art:

```
go run . "Hello"
```

Use a different banner:

```
go run . "Hello" shadow
```

Color the whole string:

```
go run . --color=green "Hello"
```

Color a substring:

```
go run . --color=red lo "Hello"
```

Color a substring with a banner:

```
go run . --color=blue lo "Hello" shadow
```

Example with substring match:

```
go run . --color=red kit "a king kitten have kit"
```

This colors:

* kit in kitten
* kit at the end of the sentence

## Project Structure

```
ascii-art-color/
│
├── ascii.go
├── banner.go
├── colour.go
├── main.go
├── ascii_test.go
│
├── standard.txt
├── shadow.txt
└── thinkertoy.txt
```

### File Responsibilities

`main.go`

Handles command line argument parsing and program execution.

`banner.go`

Loads ASCII banner fonts from text files and maps characters to their ASCII representations.

`ascii.go`

Responsible for rendering ASCII art and applying color formatting.

`colour.go`

Handles color detection and ANSI color code mapping.

`ascii_test.go`

Contains unit tests for validating color handling and ASCII rendering.

## Running the Program

Make sure you are inside the project directory:

```
cd ascii-art-color
```

Run the program:

```
go run . "Hello"
```

Or with color:

```
go run . --color=red "Hello"
```

## Running Tests

Run all unit tests with:

```
go test ./...
```

Example output:

```
PASS
ok      ascii-art-color   0.003s
```

## Implementation Details

The program works in three main stages:

1. Parse CLI arguments
2. Load the ASCII banner font
3. Render ASCII characters line by line

Each ASCII character consists of **8 rows**.
Characters are loaded from banner files and stored in a map:

```
map[rune][]string
```

Coloring is implemented using **ANSI escape codes**.

Example:

```
\033[31m   → red
\033[0m    → reset
```

These codes wrap the ASCII characters when printed.

## Requirements

* Go 1.22+
* Terminal that supports ANSI colors

## Allowed Packages

This project uses only the Go standard library.

## Learning Objectives

This project helps practice:

* Go file system operations
* Working with maps and slices
* Terminal color formatting
* Command line argument parsing
* Writing unit tests in Go
* Modular project structure

## Author

ASCII Art Color project implemented in Go as part of a Go learning exercise.
