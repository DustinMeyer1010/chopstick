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
	RAW
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

func NewTerminal() terminal {
	width, height := getTerminalSize()
	state, _ := term.GetState(int(os.Stdin.Fd()))
	return terminal{height: width, width: height, mode: NORMAL, wrap: false, termState: state}
}

// Set the height of the terminal
func (t terminal) Height(n int) terminal {
	t.height = n
	return t
}

func (t terminal) Width(n int) terminal {
	t.width = n
	return t
}

func (t terminal) NoWrap() terminal {
	t.wrap = false
	return t
}

// Chopstick will wrap and text will wrap
func (t terminal) Wrap() terminal {
	t.wrap = true
	return t
}

func (t terminal) Normal() terminal {
	print(ALTERNATE_EXIT)
	t.mode = NORMAL
	term.Restore(int(os.Stdin.Fd()), t.termState)
	state, _ := term.GetState(int(os.Stdin.Fd()))
	t.termState = state
	return t
}

func (t terminal) RawMode() terminal {
	print(ALTERNATE_EXIT)
	t.mode = RAW
	t.termState, _ = term.MakeRaw(int(os.Stdin.Fd()))
	return t
}

func (t terminal) ALTERNATE() terminal {
	print(ALTERNATE_START)
	t.mode = ALTERNATE
	return t
}

func getTerminalSize() (int, int) {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		log.Fatal("Failed to retrieve dimensions sizes")
	}
	return height, width
}
