package main

import (
	"github.com/DustinMeyer1010/chopstick"
)

func main() {
	stick := chopstick.NewChopstick().Terminal(chopstick.NewTerminal().HorizontalWrap()).Shape(chopstick.BlinkingBar)
	orginal := chopstick.Position{X: 0, Y: 0}
	chopstick.LogInit()
	chopstick.LogInit()
	redbar := chopstick.Ingredients{
		Position: chopstick.Position{X: 175, Y: orginal.Y},
		Value:    "\033[1;32;41m                                    \033[0m",
	}
	redbar1 := chopstick.Ingredients{
		Position: chopstick.Position{X: 175, Y: orginal.Y + 1},
		Value:    "\033[1;32;41m         \033[0m",
	}
	redbar2 := chopstick.Ingredients{
		Position: chopstick.Position{X: 175, Y: orginal.Y + 2},
		Value:    "\033[1;32;41m         \033[0m",
	}
	redbar3 := chopstick.Ingredients{
		Position: chopstick.Position{X: 175, Y: orginal.Y + 2},
		Value:    "\033[1;32;41m         \033[0m",
	}
	bento := chopstick.Bento{redbar, redbar1, redbar2, redbar3}
	bento.Draw(&stick)
}
