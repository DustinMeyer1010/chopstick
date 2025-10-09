package main

import (
	"time"

	"github.com/DustinMeyer1010/chopstick"
)

func main() {
	chopstick.LogInit()
	stick := chopstick.NewChopstick().Terminal(chopstick.NewTerminal().ALTERNATE()).Shape(chopstick.BlinkingBar)
	redbar := chopstick.Ingredients{
		Position: chopstick.Position{X: 0, Y: 5},
		Value:    "\033[1;32;41m          \033[0m",
	}
	time.Sleep(time.Second * 10)

	redbar.Draw(&stick)
	time.Sleep(time.Second * 10)
}
