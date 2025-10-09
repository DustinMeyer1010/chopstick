package chopstick

import (
	"log"
	"os"

	"golang.org/x/term"
)

type mode int

const (
	OFFSET = 1 // For the terminal starting at 0 and not one
)

const (
	NORMAL mode = iota
	ALTERNATE
)

const (
	ALTERNATE_START = "\033[?1049h"
	ALTERNATE_EXIT  = "\033[?1049l"
)

type terminal struct {
	verticalWrap   bool
	horizontalWrap bool
	mode           mode
	runeMatrix     Canvas
	height         int
	width          int
	termState      *term.State
}

type Canvas [][]rune

// Creates a new terimal
//
// Default Values - Will automatically put terminal in raw mode
//
//	Height: Max Terminal Height
//	Width: Max Terminal Width
//	Mode: Normal Mode
//	VerticalWrap: False
//	HorizontalWrap: False
func NewTerminal() terminal {
	width, height := getTerminalSize()
	termState, _ := term.MakeRaw(int(os.Stdin.Fd()))
	return terminal{
		height:         width - OFFSET,
		width:          height - OFFSET,
		mode:           NORMAL,
		verticalWrap:   false,
		horizontalWrap: false,
		runeMatrix:     makeMatrix(height, width),
		termState:      termState,
	}
}

func makeMatrix(height, width int) Canvas {
	RuneMatrix := make([][]rune, height) // make the outer slice with 'height' rows
	for y := range RuneMatrix {
		row := make([]rune, width)
		for x := range row {
			row[x] = ' '
		}
		RuneMatrix[y] = row // make each row with 'width' columns
	}

	return RuneMatrix
}

// Set the height of the terminal
//
// If N is greater than terminal height default to terminal height
func (t terminal) Height(n int) terminal {
	t.height = min(n, t.height) - OFFSET
	t.runeMatrix = makeMatrix(t.height+OFFSET, t.width+OFFSET)
	return t
}

// Set the width of the terminal
//
// If N is greater than terminal height default to terminal width
func (t terminal) Width(n int) terminal {
	t.width = min(n, t.width) - OFFSET
	t.runeMatrix = makeMatrix(t.height+OFFSET, t.width+OFFSET)
	return t
}

// Set terminal for nowrapping of lines
func (t terminal) VerticalWrap() terminal {
	t.verticalWrap = true
	return t
}

// Set terminal to wrap lines
func (t terminal) HorizontalWrap() terminal {
	t.horizontalWrap = true
	return t
}

// Set terminal to normal mode
func (t terminal) Normal() terminal {
	print(ALTERNATE_EXIT)
	t.mode = NORMAL
	return t
}

// Set terminal to alternate mode
func (t terminal) ALTERNATE() terminal {
	print(ALTERNATE_START)
	t.mode = ALTERNATE
	return t
}

// Returns True is vertical wrap for terminal is on
func (t terminal) HasVerticalWrap() bool {
	return t.verticalWrap
}

// Returns False if horizontal wrap for terminal is on
func (t terminal) HasHorizontalWrap() bool {
	return t.horizontalWrap
}

// Retrieves the current terminal size
func getTerminalSize() (int, int) {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		log.Fatal("Failed to retrieve dimensions sizes")
	}
	return height, width
}
