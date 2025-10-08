package chopstick

import (
	"fmt"
	"regexp"
	"strings"
)

var ansiRegex = regexp.MustCompile(`\x1b\[[0-9;]*m`)

// Draw Text to screen ansi character are ignored
func (c *chopstick) DrawText(text ...string) {
	printString := strings.Join(text, "")
	realLength := len(ansiRegex.ReplaceAllString(printString, ""))

	debug.Println(realLength)

	inEscape := false

	for i := 0; i < len(printString); i++ {
		b := printString[i]

		if inEscape {
			fmt.Printf("%c", b) // still print escape characters normally
			// check if this is the end of an escape sequence (a letter)
			if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
				inEscape = false
			}
			continue
		}

		if b == 0x1b && i+1 < len(printString) && printString[i+1] == '[' {
			inEscape = true
			fmt.Printf("%c", b) // print the ESC
			continue
		}

		switch b {
		case '\n':
			c.Down()
		default:
			fmt.Printf("%c", b)
			Print(LeftArrow)
			c.Right()
		}

	}

}

func DrawTextWithReturn(text ...string) {
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
