package chopstick

import "fmt"

var Print = fmt.Print
var Printf = fmt.Printf

// Cursor aka Chopstick
type chopstick struct {
	position Position
	terminal terminal
	shape    cshape
}

type code string

type cshape string

// The way the chopstick will apear in the terminal
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
	RightArrowN            code = "\x1b[%dC"  // Move chopstick to right n times
	LeftArrow              code = "\033[D"    // Moves chopstick left
	DownArrow              code = "\033[B"    // Moves chopstick down
	UpArrow                code = "\033[A"    // Moves chopstick up
	Return                 code = "\r"        // Returns to start of line
	ClearToEndOfTerminal   code = "\033[0J"   // Erase from chopstick to end of page
	ClearToStartOfTerminal code = "\033[1J"   // Erase from chopstick to start of page
	ClearEntireTerminal    code = "\033[2J"   // Erase entire screen
	ClearToEndOfLine       code = "\033[0K"   // Erase from chopstick to end of line
	ClearToStartOfLine     code = "\033[1K"   // Erase from chopstick to start of line
	ClearEntireLine        code = "\033[2K"   // Erase entire line chopstick is on
)

// Creates a new chopstick
func NewChopstick() chopstick {
	Print(Start)
	Print(Default)
	return chopstick{position: Position{X: 0, Y: 0}, terminal: NewTerminal(), shape: Default}
}

// Set the terminal for chopstick will use default terminal
func (c chopstick) Terminal(terminal terminal) chopstick {
	c.terminal = terminal
	return c
}

// Sets the shape fo the chopstick
func (c chopstick) Shape(shape cshape) chopstick {
	Print(shape)
	c.shape = shape
	return c
}

// Hides the chopstick
func (c chopstick) Hide() {
	Print(Hide)
}

// Shows the chopstick
func (c chopstick) Show() {
	Print(Show)
}

// Update the terminal
func (c *chopstick) UpdateTerminal(t terminal) {
	c.terminal = t
	c.StartOfPage()
}

// Returns the metadata under the chopstick
func (c *chopstick) GetValueUnderChopstick() string {
	return c.terminal.canvas.getValue(c.position)
}

// Return element under the chopstick
func (c *chopstick) GetElementUnderChopstick() *Element {
	element := c.terminal.canvas.getElement(c.position)
	return element
}

// Return value at specific location
//
// If position outside of terminal bounds will throw error
func (c *chopstick) GetValueAtLocation(p Position) (string, error) {
	if p.OutOfBounds(*c) {
		return "", fmt.Errorf("invalid position")
	}
	return c.terminal.canvas.getValue(p), nil
}

// Return element at specific location
//
// If position outside of terminal bounds will throw error
func (c *chopstick) GetElementAtLocation(p Position) (*Element, error) {
	if p.OutOfBounds(*c) {
		return nil, fmt.Errorf("invalid position")
	}
	return c.terminal.canvas.getElement(p), nil
}

// Set metadata for a specific location
//
// If position outside of terminal bounds will throw error
func (c *chopstick) SetMetaDataAtLocation(p Position, metadata any) error {
	if p.OutOfBounds(*c) {
		return fmt.Errorf("invalid position")
	}
	c.terminal.canvas.setMetaData(p, metadata)
	return nil
}

// Returns metadata for element under the chopstick
func (c *chopstick) GetMetaDataAtChopstick() any {
	return c.terminal.canvas.getMetaData(c.position)
}

// Set an element at a specific location
//
// If position outside of terminal bounds will throw error
func (c *chopstick) SetElementAtLocation(p Position, element Element) error {
	if p.OutOfBounds(*c) {
		return fmt.Errorf("invalid position")
	}
	c.terminal.canvas.setElement(p, element)
	prevPosition := Position{X: c.position.X, Y: c.position.Y}
	c.MoveTo(p)
	c.Draw(string(element.Value))
	c.MoveTo(prevPosition)
	return nil
}

// Sets element under chopstick
func (c *chopstick) SetElementAtChopstick(element Element) {
	c.terminal.canvas.setElement(c.position, element)
	c.DrawWithReturn(string(element.Value))
}

// Returns the current position of the chopstick
func (c *chopstick) GetPosition() Position {
	return c.position
}

// Get the current X positon of the chopstick
func (c *chopstick) GetX() int {
	return c.position.X
}

// Get the current Y position of the chopstick
func (c *chopstick) GetY() int {
	return c.position.Y
}

func (c *chopstick) GetKeyPressed() string {
	return c.terminal.GetKeyPressed()
}

// Position of the chopstick
type Position struct {
	X int
	Y int
}

// Checks to make sure position is within bounds of the terminal width and height
func (p Position) OutOfBounds(c chopstick) bool {
	return p.X > c.terminal.width || p.Y > c.terminal.height || p.Y < 0 || p.X < 0
}
