package main

import (
	"fmt"
	"os"
	"time"

	"github.com/DustinMeyer1010/chopstick"
)

func main() {
	chopstick.LogInit()

	i1 := chopstick.Ingrident{
		Position: chopstick.Position{X: 0, Y: 5},
		Value:    "\033[1;32;41m          \033[0m",
	}

	i2 := chopstick.Ingrident{
		Position: chopstick.Position{X: 0, Y: 6},
		Value:    "\033[1;32;41m          \033[0m",
	}

	i3 := chopstick.Ingrident{
		Position: chopstick.Position{X: 0, Y: 7},
		Value:    "\033[1;32;41m          \033[0m",
	}

	bento := chopstick.Bento{i1, i2, i3}

	ch := chopstick.NewChopstick().
		Terminal(
			chopstick.
				NewTerminal().
				ALTERNATE().
				Height(10).
				Width(10),
		).
		Shape(chopstick.BlinkingUnderline)

	exit := true
	for exit {
		switch GetKeyPressed() {
		case "\033[A":
			ch.Up()
		case "\033[B":
			ch.Down()
		case "\033[C":
			ch.Right()
		case "\033[D":
			ch.Left()
		case "u":
			bento.Draw(&ch)
		default:
			exit = false
		}
	}
	fmt.Println(ch)
	time.Sleep(time.Second)
}

func GetKeyPressed() string {
	var buf = make([]byte, 3)
	n, err := os.Stdin.Read(buf)

	if err != nil {
		panic(err)
	}

	if n == 0 {
		return string("")
	}

	return string(buf[:n])
}
