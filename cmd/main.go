package main

import (
	"fmt"
	"os"
	"time"

	"github.com/DustinMeyer1010/chopstick"
)

func main() {
	chopstick.LogInit()
	//text := "\033[1;32;41mBold green text on red background\033[0m\n\n"

	ch := chopstick.NewChopstick().
		Terminal(
			chopstick.
				NewTerminal().
				ALTERNATE().
				Height(10).
				Width(10).
				Wrap(),
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
			ch.RightN(5)
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
