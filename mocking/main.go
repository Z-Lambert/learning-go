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


type Sleeper interface {
	Sleep() 
}

type DefaultSleeper struct {}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(countdownMultiplier * time.Second)
}

func main () {
	Countdown(os.Stdout, &DefaultSleeper{})
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}


	fmt.Fprint(out, finalWord)
}