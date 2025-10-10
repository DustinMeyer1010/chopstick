package chopstick

import (
	"fmt"
)

// Move Chopstick to right in terminal
func (c *chopstick) Right() {

	if !c.IsAtEnd() {
		c.position.X++
		Print(RightArrow)
		return
	}

	switch {
	case c.terminal.HasLineWrap():
		c.rightWithLineWrap()
	case c.terminal.HasHorizontalWrap():
		c.rightWithHorizontalWrap()
	}
}

// Wraps the chopstick down & to start of line
func (c *chopstick) rightWithLineWrap() {
	c.Down()
	c.StartOfLine()
}

// Wraps the copstick back to start of line
func (c *chopstick) rightWithHorizontalWrap() {
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

// Moves Chopstick left in terminal
func (c *chopstick) Left() {
	if !c.IsAtStart() {
		c.position.X--
		Print(LeftArrow)
		return
	}

	switch {
	case c.terminal.HasLineWrap():
		c.leftWithLineWrap()
	case c.terminal.HasHorizontalWrap():
		c.leftWithHorizontalWrap()
	}
}

// Wraps chopstick up & to end of line
func (c *chopstick) leftWithLineWrap() {
	c.Up()
	c.EndOfLine()
}

// Wraps chopstick to end of line
func (c *chopstick) leftWithHorizontalWrap() {
	c.EndOfLine()
}

// TODO:
//   - Comment
func (c *chopstick) LeftN(n int) {
	for range n {
		c.Left()
	}
}

// Move chopstick without changing position of x (Used for printing character which move the chopstick on it own)
func (c *chopstick) left() {
	Print(LeftArrow)
}

// Moves the chopstick up in terminal
func (c *chopstick) Up() {

	if !c.IsAtTop() {
		c.position.Y--
		Print(UpArrow)
		return
	}

	if c.terminal.HasVerticalWrap() {
		c.upWithVerticalWrap()
	}
}

// Wraps chopstick back to bottom
func (c *chopstick) upWithVerticalWrap() {
	c.Bottom()
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

	if !c.IsAtBottom() {

		c.position.Y++
		Print(DownArrow)
		return
	}

	if c.terminal.HasVerticalWrap() {
		c.downWithVerticalWrap()
	}

}

// Handles wrapping of down in terminal
func (c *chopstick) downWithVerticalWrap() {
	c.Top()
}

// TODO:
//   - Comment
func (c *chopstick) DownN(n int) {
	for range n {
		c.Down()
	}
}

// check for chopstick if at top of terminal
func (c chopstick) IsAtTop() bool {
	return c.position.Y <= 0
}

// check for chopstick if at bottom of terminal
func (c chopstick) IsAtBottom() bool {
	return c.position.Y >= c.terminal.height
}

// check for chopstick if at end of terminal
func (c chopstick) IsAtEnd() bool {
	return c.position.X >= c.terminal.width
}

// check for chopstick if at start of terminal
func (c chopstick) IsAtStart() bool {
	return c.position.X <= 0
}

// Moves Chopstick to end of current line
func (c *chopstick) EndOfLine() {
	Print(fmt.Sprintf("\033[%dG", c.terminal.width-c.position.X))
	c.position.X = c.terminal.width
}

// Move Chopstick to start of current line
func (c *chopstick) StartOfLine() {
	Print(Return)
	c.position.X = 0
}

// Moves Chopstick to Top and Start of line
func (c *chopstick) StartOfPage() {
	Print(Start)
	c.position.X = 0
	c.position.Y = 0
}

// Move chopstick to Bottom and end of line
func (c *chopstick) EndOfPage() {
	Print(fmt.Sprintf("\033[%d;%dH", c.terminal.height-c.position.Y, c.terminal.width-c.position.Y))
	c.position.X = c.terminal.width
	c.position.Y = c.terminal.height
}

// Moves chopstick to top of page keeping x
func (c *chopstick) Top() {
	fmt.Printf("\033[%dA", c.terminal.height-c.position.Y)
	c.position.Y = 0
}

// Moves chopstick to bottom of page keeping x
func (c *chopstick) Bottom() {
	fmt.Printf("\033[%dB", c.terminal.height-c.position.Y)
	c.position.Y = c.terminal.height
}

// Move to any cordinate
//
// Eample: Height, Width = 10,  x = 55,  y = 68
//
// x will move to 5 and y will move to 8
func (c *chopstick) MoveTo(p Position) {
	c.StartOfPage()
	for range p.X % (c.terminal.width + OFFSET) {
		c.Right()
	}
	for range p.Y % (c.terminal.height + OFFSET) {
		c.Down()
	}
}
