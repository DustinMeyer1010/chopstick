package main

import (
	"fmt"
	"os"
	"time"

	"github.com/DustinMeyer1010/chopstick"
)

func main() {
	chopstick.LogInit()

	ch := chopstick.NewChopstick().
		Terminal(
			chopstick.
				NewTerminal(),
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
			//bento.Draw(&ch)
			ch.DrawText("hello world")
		case "c":
			element := ch.GetElementUnderChopstick()
			element.MetaData = "Test"
			chopstick.Debug.Println(ch.GetElementUnderChopstick().MetaData)
		case "a":
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
