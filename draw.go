package chopstick

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

var ansiRegex = regexp.MustCompile(`\x1b\[[0-9;]*m`)

// TODO:
//   - Refactor the function to make it short and readable
//   - Fix wrap at the bottom of the terminal when only using horizontalWrap
//   - Does not support UTF-8 Character in terminal (Two columns in length rather than one)??????
//
// Draw Text to screen ansi character are ignored
func (c *chopstick) DrawText(text ...string) {
	printString := strings.Join(text, "")
	realLength := len([]rune(ansiRegex.ReplaceAllString(printString, ""))) // rune-aware length
	debug.Println(realLength)

	inEscape := false
	prevX := c.position.x

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
				Print(" ")
				c.left()
				c.Right()
			}
		default:
			if unicode.IsPrint(r) {
				Printf("%s", string(r))
				c.left()
				c.Right()
			}

		}

		if prevX == c.position.x && !c.terminal.HasHorizontalWrap() {
			break
		}
		prevX = c.position.x
	}
}

// Samething as DrawText but return cursor to orginal position before drawing text
func (c *chopstick) DrawTextWithReturn(text ...string) {
	prevX := c.position.x
	prevY := c.position.y
	c.DrawText(text...)
	c.MoveTo(prevX, prevY)
}

// Erase the Entire terminal
func (c *chopstick) EraseTerminal() {
	Print(EraseEntireTerminal)
}

// Erase from chopstick to end of terminal
func (c *chopstick) EraseToEndOfTermial() {
	Print(EraseToEndOfTerminal)
}

// Erase from chopstick to Start of terminal
func (c *chopstick) EraseToStartOfTerminal() {
	Print(EraseToStartOfTerminal)
}

// Erase from chopstick to Start of line
func (c *chopstick) EraseToStartOfLine() {
	Print(EraseToStartOfLine)
}

// Erase from chopstick to End of line
func (c *chopstick) EraseToEndOfLine() {
	Print(EraseToEndOfLine)
}
