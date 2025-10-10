package chopstick

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

//var ansiRegex = regexp.MustCompile(`\x1b\[[0-9;]*m`)

// TODO:
//   - Refactor the function to make it short and readable
//   - Fix wrap at the bottom of the terminal when only using horizontalWrap
//   - Does not support UTF-8 Character in terminal (Two columns in length rather than one)??????
//
// Draw Text to screen ansi character are ignored
func (c *chopstick) Draw(text ...string) {
	printString := strings.Join(text, "")

	inEscape := false
	prevX := c.position.X

	for i := 0; i < len(printString); {
		r, size := utf8.DecodeRuneInString(printString[i:])
		i += size // advance by rune length

		if inEscape {
			fmt.Print(string(r)) // still print escape chars
			if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
				inEscape = false
			}
			continue
		}

		// Detect escape start
		if r == '\x1b' && i < len(printString) && printString[i] == '[' {
			inEscape = true
			fmt.Print(string(r))
			continue
		}

		switch r {
		case '\n':
			c.Down()
		case '\t':
			for range 4 {
				c.terminal.canvas.setValue(c.position, r)
				Print(" ")
				c.left()
				c.Right()
			}
		default:
			if unicode.IsPrint(r) {
				c.terminal.canvas.setValue(c.position, r)
				Printf("%s", string(r))
				c.left()
				c.Right()
			}

		}

		if prevX == c.position.X && !c.terminal.HasHorizontalWrap() {
			break
		}
		prevX = c.position.X
	}
}

// Draw the text
func (c *chopstick) DrawWithReturn(text ...string) {
	prevX := c.position.X
	prevY := c.position.Y
	c.Draw(text...)
	c.MoveTo(Position{X: prevX, Y: prevY})
}

// TODO:
// - Clear the canvas as previous elements will still be at those locations

// Clear the Entire terminal
func (c *chopstick) ClearTerminal() {
	Print(EraseEntireTerminal)
}

// Clear from chopstick to end of terminal
func (c *chopstick) ClearToEndOfTermial() {
	Print(EraseToEndOfTerminal)
}

// Clear from chopstick to Start of terminal
func (c *chopstick) ClearToStartOfTerminal() {
	Print(EraseToStartOfTerminal)
}

// Clear from chopstick to Start of line
func (c *chopstick) ClearToStartOfLine() {
	Print(EraseToStartOfLine)
}

// Clear from chopstick to End of line
func (c *chopstick) ClearToEndOfLine() {
	Print(EraseToEndOfLine)
}
