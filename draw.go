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
		drawnLen := 0
		for drawnLen < length {
			difference := c.terminal.width - c.x
			if c.y >= c.terminal.height {
				return
			}
			if difference > visibleLength(printString) {
				Print(printString)
				return
			} else {
				Print(printString[:difference])
				printString = printString[difference:]
				c.Down()
				c.StartOfLine()
			}

		}
		return

	}
	Print(printString)

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
