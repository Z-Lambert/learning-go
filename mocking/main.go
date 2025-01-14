package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3
const countdownMultiplier = 1

func main () {
	Countdown(os.Stdout)
}

func Countdown (out io.Writer) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintf(out, "%d\n", i)
		time.Sleep(countdownMultiplier * time.Second)

	}
	fmt.Fprintf(out, "%s\n", finalWord)
} 