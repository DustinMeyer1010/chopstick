package chopstick

import (
	"fmt"
)

// Moves chopstick to the right
//
// Wrap: Chopstick will move down and to start of line
//
// NoWrap: Chopstick will stop at end of line
func (c *chopstick) Right() {
	if c.IsAtEnd() {
		if c.terminal.HasWrap() {
			debug.Printf("wrapping: %d\n", c.x)
			c.rightWithWrap()
		}
		debug.Printf("Right: %d\n", c.x)
		return
	}
	c.x++
	debug.Printf("Right: %d\n", c.x)
	Print(RightArrow)
}

// Handles wrapping of right in terminal
func (c *chopstick) rightWithWrap() {
	c.Down()
	c.StartOfLine()
}

// Move chopstick right n times
func (c *chopstick) RightN(n int) {
	for range n + int(n/c.terminal.width-c.x) {
		c.Right()
	}
}

// Moves Chopstick to the left
//
// Wrap: Chopstick will move to up and end of line
//
// NoWrap: Chopstick will stop at start of line
func (c *chopstick) Left() {
	if c.IsAtStart() {
		if c.terminal.HasWrap() {
			c.leftWithWrap()
		}
		debug.Printf("Left: %d\n", c.x)
		return
	}

	c.x--
	debug.Printf("Left: %d\n", c.x)
	Print(LeftArrow)
}

// Handles wrapping of left in terminal
func (c *chopstick) leftWithWrap() {
	c.Up()
	c.EndOfLine()
	c.x = c.terminal.width
}

func (c *chopstick) LefttN(n int) {
	for range n + int(n/c.terminal.width) {
		c.Left()
	}
}

// Moves Chopstick Up
//
// Wrap: Chopstick will move bottom of page and keep current x
//
// NoWrap: Chopstick will stop at top of terminal
func (c *chopstick) Up() {

	if c.IsAtTop() {
		if c.terminal.HasWrap() {
			c.upWithWrap()
		}
		debug.Printf("Up: %d\n", c.y)
		return
	}

	c.y--
	debug.Printf("Up: %d\n", c.y)
	Print(UpArrow)
}

// Handles wrapping of up in terminal
func (c *chopstick) upWithWrap() {
	c.Bottom()
	c.y = c.terminal.height
}

// Moves Chopstick Down
//
// Wrap: Chopstick will move top of page and keep current x
//
// NoWrap: Chopstick will stop at bottom of terminal
func (c *chopstick) Down() {

	if c.IsAtBottom() {
		if c.terminal.HasWrap() {
			c.downWithWrap()
			debug.Printf("Down: %d\n", c.y)
		}
		return
	}

	c.y++
	debug.Printf("Down: %d\n", c.y)
	Print(DownArrow)
}

// Handles wrapping of down in terminal
func (c *chopstick) downWithWrap() {
	c.y = 0
	c.Top()
}

// If chopstick at top of terminal returns True
func (c chopstick) IsAtTop() bool {
	return c.y <= 0
}

// If chopstick at bottom of terminal returns True
func (c chopstick) IsAtBottom() bool {
	return c.y >= c.terminal.height
}

// If chopstick at end of line returns True
func (c chopstick) IsAtEnd() bool {
	return c.x >= c.terminal.width
}

// If chopstick at start of line returns True
func (c chopstick) IsAtStart() bool {
	return c.x <= 0
}

// Moves Chopstick to end of current line
func (c *chopstick) EndOfLine() {
	Print(fmt.Sprintf("\033[%dG", c.terminal.width-c.x))
	c.x = c.terminal.width
}

// Move Chopstick to start of current line
func (c *chopstick) StartOfLine() {
	Print(Return)
	c.x = 0
}

// Moves Chopstick to Top and Start of line
func (c *chopstick) StartOfPage() {
	Print(Start)
	c.x = 0
	c.y = 0
}

// Move chopstick to Bottom and end of line
func (c *chopstick) EndOfPage() {
	Print(fmt.Sprintf("\033[%d;%dH", c.terminal.height-c.y, c.terminal.width-c.x))
	c.x = c.terminal.width
	c.y = c.terminal.height
}

// Moves chopstick to top of page keeping x
func (c *chopstick) Top() {
	fmt.Printf("\033[%dA", c.terminal.height-c.y)
}

// Moves chopstick to bottom of page keeping x
func (c *chopstick) Bottom() {
	fmt.Printf("\033[%dB", c.terminal.height-c.y)
}

// Move to any cordinate
//
// Eample: Height, Width = 10,  x = 55,  y = 68
//
// x will move to 5 and y will move to 8
func (c *chopstick) MoveTo(x, y int) {
	c.StartOfPage()
	for range x {
		c.Right()
	}
	for range y {
		c.Down()
	}
}
