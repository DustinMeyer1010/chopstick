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
		if c.terminal.HasHorizontalWrap() {
			debug.Printf("wrapping: %d\n", c.position.x)
			c.rightWithWrap()
		}
		debug.Printf("Right: %d\n", c.position.x)
		return
	}
	c.position.x++
	debug.Printf("Right: %d\n", c.position.x)
	Print(RightArrow)
}

// Handles wrapping of right in terminal
func (c *chopstick) rightWithWrap() {
	c.Down()
	c.StartOfLine()
}

// RightN moves the chopstick cursor to the right n times.
//
// Note: Wrapping is included in the count. This means if the terminal width
// is 10, calling RightN(10) will move the cursor to the end of the current line,
// while RightN(11) will move it to the start of the next line (after wrapping).
//
// If no wrap then n will just stop at end of line or width of the terminal
func (c *chopstick) RightN(n int) {
	for range n {
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
		if c.terminal.HasHorizontalWrap() {
			c.leftWithWrap()
		}
		debug.Printf("Left: %d\n", c.position.x)
		return
	}

	c.position.x--
	debug.Printf("Left: %d\n", c.position.x)
	Print(LeftArrow)
}

// Handles wrapping of left in terminal
func (c *chopstick) leftWithWrap() {
	c.Up()
	c.EndOfLine()
	c.position.x = c.terminal.width
}

// TODO:
//   - Comment
func (c *chopstick) LeftN(n int) {
	for range n {
		c.Left()
	}
}

// Move chopstick without changing position of x
func (c *chopstick) left() {
	Print(LeftArrow)
}

// Moves Chopstick Up
//
// Wrap: Chopstick will move bottom of page and keep current x
//
// NoWrap: Chopstick will stop at top of terminal
func (c *chopstick) Up() {

	if c.IsAtTop() {
		if c.terminal.HasVerticalWrap() {
			c.upWithWrap()
		}
		debug.Printf("Up: %d\n", c.position.y)
		return
	}

	c.position.y--
	debug.Printf("Up: %d\n", c.position.y)
	Print(UpArrow)
}

// Handles wrapping of up in terminal
func (c *chopstick) upWithWrap() {
	c.Bottom()
	c.position.y = c.terminal.height
}

// TODO:
//   - Comment
func (c *chopstick) UpN(n int) {
	for range n {
		c.Up()
	}
}

// Moves Chopstick Down
//
// Wrap: Chopstick will move top of page and keep current x
//
// NoWrap: Chopstick will stop at bottom of terminal
func (c *chopstick) Down() {

	if c.IsAtBottom() {
		if c.terminal.HasVerticalWrap() {
			c.downWithWrap()
			debug.Printf("Down: %d\n", c.position.y)
		}
		return
	}

	c.position.y++
	debug.Printf("Down: %d\n", c.position.y)
	Print(DownArrow)
}

// Handles wrapping of down in terminal
func (c *chopstick) downWithWrap() {
	c.position.y = 0
	c.Top()
}

// TODO:
//   - Comment
func (c *chopstick) DownN(n int) {
	for range n {
		c.Down()
	}
}

// If chopstick at top of terminal returns True
func (c chopstick) IsAtTop() bool {
	return c.position.y <= 0
}

// If chopstick at bottom of terminal returns True
func (c chopstick) IsAtBottom() bool {
	return c.position.y >= c.terminal.height
}

// If chopstick at end of line returns True
func (c chopstick) IsAtEnd() bool {
	return c.position.x >= c.terminal.width
}

// If chopstick at start of line returns True
func (c chopstick) IsAtStart() bool {
	return c.position.x <= 0
}

// Moves Chopstick to end of current line
func (c *chopstick) EndOfLine() {
	Print(fmt.Sprintf("\033[%dG", c.terminal.width-c.position.x))
	c.position.x = c.terminal.width
}

// Move Chopstick to start of current line
func (c *chopstick) StartOfLine() {
	Print(Return)
	c.position.x = 0
}

// Moves Chopstick to Top and Start of line
func (c *chopstick) StartOfPage() {
	Print(Start)
	c.position.x = 0
	c.position.y = 0
}

// Move chopstick to Bottom and end of line
func (c *chopstick) EndOfPage() {
	Print(fmt.Sprintf("\033[%d;%dH", c.terminal.height-c.position.y, c.terminal.width-c.position.y))
	c.position.x = c.terminal.width
	c.position.y = c.terminal.height
}

// Moves chopstick to top of page keeping x
func (c *chopstick) Top() {
	fmt.Printf("\033[%dA", c.terminal.height-c.position.y)
}

// Moves chopstick to bottom of page keeping x
func (c *chopstick) Bottom() {
	fmt.Printf("\033[%dB", c.terminal.height-c.position.y)
}

// Move to any cordinate
//
// Eample: Height, Width = 10,  x = 55,  y = 68
//
// x will move to 5 and y will move to 8
func (c *chopstick) MoveTo(x, y int) {
	c.StartOfPage()
	for range x % (c.terminal.width + OFFSET) {
		c.Right()
	}
	for range y % (c.terminal.height + OFFSET) {
		c.Down()
	}
}
