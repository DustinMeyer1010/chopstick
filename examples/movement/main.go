package main

import (
	"fmt"

	"github.com/DustinMeyer1010/chopstick"
)

type Key string

const (
	UpArrow    Key = "\033[A"
	DownArrow  Key = "\033[B"
	RightArrow Key = "\033[C"
	LeftArrow  Key = "\033[D"
	ControlC   Key = "\x03"
)

func main() {
	// Define chopstick

	stick := chopstick.NewChopstick().
		Terminal(
			chopstick.
				NewTerminal().    // Creates a new terminal
				ALTERNATE().      // Put the terminal in Alternate Mode
				HorizontalWrap(). // Turns on HorizontalWrap
				LineWrap().       // Turns on LineWrap
				Width(20),        // Set the max width of terminal to 20
		).
		Shape(chopstick.SteadyBlock) // Changes the shape of the chopstick
	fmt.Print(RightArrow)
	fmt.Print(RightArrow)
	fmt.Print(RightArrow)
	fmt.Print(DownArrow)
	exit := true
	for exit {
		switch key := stick.GetKeyPressed(); Key(key) {
		case UpArrow:
			stick.Up()
		case DownArrow:
			stick.Down()
		case RightArrow:
			stick.Right()
		case LeftArrow:
			stick.Left()
		case ControlC:
			exit = false
		default:
			// Draw all other characters
			stick.Draw(key)
		}
	}
}
