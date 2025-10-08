package chopstick

import (
	"log"
	"os"

	"golang.org/x/term"
)

type mode int

const (
	NORMAL mode = iota
	ALTERNATE
)

const (
	ALTERNATE_START = "\033[?1049h"
	ALTERNATE_EXIT  = "\033[?1049l"
)

type terminal struct {
	wrap      bool
	mode      mode
	height    int
	width     int
	termState *term.State
}

// Creates a new terimal
//
// Default Values
//
//	Height: Max Terminal Height
//	Width: Max Terminal Width
//	Mode: Normal Mode
//	wrap: False
func NewTerminal() terminal {
	width, height := getTerminalSize()
	termState, _ := term.MakeRaw(int(os.Stdin.Fd()))
	return terminal{height: width - 1, width: height - 1, mode: NORMAL, wrap: false, termState: termState}
}

// Set the height of the terminal
func (t terminal) Height(n int) terminal {
	t.height = n
	return t
}

// Set the Width of the terminal
func (t terminal) Width(n int) terminal {
	t.width = n
	return t
}

// Set terminal for nowrapping of lines
func (t terminal) NoWrap() terminal {
	t.wrap = false
	return t
}

// Set terminal to wrap lines
func (t terminal) Wrap() terminal {
	t.wrap = true
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

func (t terminal) HasWrap() bool {
	return t.wrap
}

// Retrieves the current terminal size
func getTerminalSize() (int, int) {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		log.Fatal("Failed to retrieve dimensions sizes")
	}
	return height, width
}
