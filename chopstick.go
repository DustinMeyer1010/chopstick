package chopstick

import "fmt"

var Print = fmt.Print

// Cursor aka Chopstick
type chopstick struct {
	x        int
	y        int
	terminal terminal
	shape    cshape
}

type code string

type cshape string

/*
ESC[0 q 	changes cursor shape to steady block
ESC[1 q 	changes cursor shape to steady block also
ESC[2 q 	changes cursor shape to blinking block
ESC[3 q 	changes cursor shape to steady underline
ESC[4 q 	changes cursor shape to blinking underline
ESC[5 q 	changes cursor shape to steady bar
ESC[6 q 	changes cursor shape to blinking bar
*/
const (
	Default           cshape = "\033[0 q"
	SteadyBlock       cshape = "\033[1 q"
	BlinkingBlock     cshape = "\033[2 q"
	SteadyUnderline   cshape = "\033[3 q"
	BlinkingUnderline cshape = "\033[4 q"
	SteadyBar         cshape = "\033[5 q"
	BlinkingBar       cshape = "\033[6 q"
)

const (
	Start                  code = "\033[H"    // Moves chopstick to 0, 0 || Start of terminal
	Hide                   code = "\033[?25l" // Hides chopstick
	Show                   code = "\033[?25h" // Shows the chopstick
	RightArrow             code = "\033[C"    // Moves chopstick right
	LeftArrow              code = "\033[D"    // Moves chopstick left
	DownArrow              code = "\033[B"    // Moves chopstick down
	UpArrow                code = "\033[A"    // Moves chopstick up
	Return                 code = "\r"        // Returns to start of line
	EraseToEndOfTerminal   code = "\033[0J"   // Erase from chopstick to end of page
	EraseToStartOfTerminal code = "\033[1J"   // Erase from chopstick to start of page
	EraseEntireTerminal    code = "\033[3J"   // Erase entire screen
	EraseToEndOfLine       code = "\033[0K"   // Erase from chopstick to end of line
	EraseToStartOfLine     code = "\033[1K"   // Erase from chopstick to start of line
	EraseEntireLine        code = "\033[2K"   // Erase entire line chopstick is on
)

// Creates a new chopstick
func NewChopstick() chopstick {
	Print(Start)
	Print(Default)
	return chopstick{x: 0, y: 0, terminal: NewTerminal(), shape: Default}
}

func (c chopstick) Terminal(terminal terminal) chopstick {
	c.terminal = terminal
	return c
}

func (c chopstick) Shape(shape cshape) chopstick {
	Print(shape)
	c.shape = shape
	return c
}

// Hides the chopstick
func (c chopstick) Hide() {
	Print(Hide)
}

// Shows the copstick
func (c chopstick) Show() {
	Print(Show)
}

// Update the terminal
func (c *chopstick) UpdateTerminal(terminal terminal) {
	c.terminal = terminal
	c.StartOfPage()
}
