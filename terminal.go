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

// Different modes that the terminal will be in
const (
	NORMAL mode = iota
	ALTERNATE
)

const (
	ALTERNATE_START = "\033[?1049h" // Enter Alternate Mode
	ALTERNATE_EXIT  = "\033[?1049l" // Exits Alternate Mode
)

type terminal struct {
	verticalWrap   bool        // Chopstick should wrap around top and bottom of terminal
	horizontalWrap bool        // Chopstick will wrap around the same line
	lineWrap       bool        // Chopstill wrap around to next line
	mode           mode        // What mode the termianl is going to be in (ALTERNATE OR ORGINAL)
	canvas         canvas      // Track location of thing printed in terminal
	height         int         // Height of the terminal
	width          int         // Width of the terminal
	termState      *term.State // Terminal state (Only used for return from rawMode)
}

// Creates a new terimal
//
// Default Values - Will automatically put terminal in raw mode
//
//	Height: Max Terminal Height
//	Width: Max Terminal Width
//	Mode: Normal Mode
//	VerticalWrap: False
//	HorizontalWrap: False
//	LineWrap: False
func NewTerminal() terminal {
	height, width := getTerminalSize()
	termState, _ := term.MakeRaw(int(os.Stdin.Fd()))
	return terminal{
		height:         height - OFFSET,
		width:          width - OFFSET,
		mode:           NORMAL,
		verticalWrap:   false,
		horizontalWrap: false,
		lineWrap:       false,
		canvas:         makeCanvas(height, width),
		termState:      termState,
	}
}

// Set the height of the terminal
//
// Sets height to Min(n, terminal_height)
func (t terminal) Height(n int) terminal {
	t.height = min(n, t.height) - OFFSET
	t.canvas = makeCanvas(t.height+OFFSET, t.width+OFFSET)
	return t
}

// Set the width of the terminal
//
// Sets width to min(n, terminal_width)
func (t terminal) Width(n int) terminal {
	t.width = min(n, t.width) - OFFSET
	t.canvas = makeCanvas(t.height+OFFSET, t.width+OFFSET)
	return t
}

// Turns on vertical wrapping for terminal
func (t terminal) VerticalWrap() terminal {
	t.verticalWrap = true
	return t
}

// Turns on Horizontal wrapping for terminal
func (t terminal) HorizontalWrap() terminal {
	t.horizontalWrap = true
	return t
}

// Set terminal to Normal Mode
func (t terminal) Normal() terminal {
	print(ALTERNATE_EXIT)
	t.mode = NORMAL
	return t
}

// Set terminal to Alternate Mode
func (t terminal) ALTERNATE() terminal {
	print(ALTERNATE_START)
	t.mode = ALTERNATE
	return t
}

func (t terminal) LineWrap() terminal {
	t.lineWrap = true
	return t
}

// Check for vertical wrap is on
func (t terminal) HasVerticalWrap() bool {
	return t.verticalWrap
}

// Check for Horizontal wrap is on
func (t terminal) HasHorizontalWrap() bool {
	return t.horizontalWrap
}

func (t terminal) HasLineWrap() bool {
	return t.lineWrap
}

func (t terminal) GetKeyPressed() string {
	var buf = make([]byte, 3)
	n, err := os.Stdin.Read(buf)

	if err != nil {
		panic(err)
	}

	if n == 0 {
		return string("")
	}

	return string(buf[:n])
}

// For getting the current terminal size
func getTerminalSize() (int, int) {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		log.Fatal("Failed to retrieve dimensions sizes")
	}
	return height, width
}
