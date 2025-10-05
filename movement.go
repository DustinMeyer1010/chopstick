package chopstick

import "fmt"

// Moves chopstick to the right
//
// Wrap: Chopstick will move down and to start of line
//
// NoWrap: Chopstick will stop at end of line
func (c *chopstick) Right() {
	if c.x >= c.terminal.width {
		if c.terminal.wrap && c.y < c.terminal.height {
			c.Down()
			c.StartOfLine()
			c.x = 0
		}
		return
	}

	c.x++
	Print(RightArrow)
}

// Moves Chopstick to the left
//
// Wrap: Chopstick will move to up and end of line
//
// NoWrap: Chopstick will stop at start of line
func (c *chopstick) Left() {
	if c.x <= 0 {
		if c.terminal.wrap && c.y > 0 {
			c.Up()
			c.EndOfLine()
			c.x = c.terminal.width
		}
		return
	}
	c.x--
	Print(LeftArrow)
}

// Moves Chopstick Up
//
// Wrap: Chopstick will move bottom of page and keep current x
//
// NoWrap: Chopstick will stop at top of terminal
func (c *chopstick) Up() {

	if c.y <= 0 {
		if c.terminal.wrap {
			c.Bottom()
			c.y = c.terminal.height
		}
		return
	}
	c.y--
	Print(UpArrow)
}

// Moves Chopstick Down
//
// Wrap: Chopstick will move top of page and keep current x
//
// NoWrap: Chopstick will stop at bottom of terminal
func (c *chopstick) Down() {
	if c.y >= c.terminal.height {
		if c.terminal.wrap {
			c.y = 0
			c.Top()
		}
		return
	}

	c.y++
	Print(DownArrow)
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
	for range x % c.terminal.width {
		c.Right()
	}
	for range y % c.terminal.height {
		c.Down()
	}
}
