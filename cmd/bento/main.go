package main

import (
	"time"

	"github.com/DustinMeyer1010/chopstick"
)

func main() {
	stick := chopstick.NewChopstick().Terminal(chopstick.NewTerminal().ALTERNATE()).Shape(chopstick.BlinkingBar)
	chopstick.LogInit()
	redbar := chopstick.Ingredients{
		Position: chopstick.Position{X: 2, Y: 5},
		Value:    "\033[1;32;41m         \033[0m",
	}
	redbar1 := chopstick.Ingredients{
		Position: chopstick.Position{X: 1, Y: 6},
		Value:    "\033[1;32;41m           \033[0m",
	}
	redbar2 := chopstick.Ingredients{
		Position: chopstick.Position{X: 1, Y: 7},
		Value:    "\033[1;32;41m           \033[0m",
	}
	redbar3 := chopstick.Ingredients{
		Position: chopstick.Position{X: 2, Y: 8},
		Value:    "\033[1;32;41m         \033[0m",
	}

	bento := chopstick.Bento{redbar, redbar1, redbar2, redbar3}
	time.Sleep(time.Second * 5)
	bento.Draw(&stick)
	time.Sleep(time.Second * 5)
}
