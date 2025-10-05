package chopstick

import (
	"regexp"
	"strings"
)

var ansiRegex = regexp.MustCompile(`\x1b\[[0-9;]*m`)

// Draw Text to screen ansi character are ignored
func (c *chopstick) DrawText(text ...string) {
	printString := strings.Join(text, "")
	length := visibleLength(printString)

	if length+c.x > c.terminal.width {
		prevX := c.x
		prevY := c.y
		difference := c.terminal.width - c.x
		Print(printString[:difference])
		c.MoveTo(prevX, prevY)
	}

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

func visibleLength(s string) int {
	return len([]rune(ansiRegex.ReplaceAllString(s, "")))
}
