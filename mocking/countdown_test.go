package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

type SpySleeper struct {
	Calls []string
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpySleeper) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpySleeper) Write(p []byte) (n int, err error)  {
	s.Calls = append(s.Calls, write)
	return
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

const sleep = "sleep"
const write = "write"

func TestCountdown (t *testing.T) {
	t.Run("Countdown prints the correct string", func (t *testing.T) {

	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}
	Countdown(buffer, spySleeper)
	got := buffer.String()
	want := "3\n2\n1\nGo!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
})
t.Run("sleep before every print", func(t *testing.T) {
	spySleeper := &SpySleeper{}
	Countdown(spySleeper, spySleeper)
	want := []string{
		write,
		sleep,
		write,
		sleep,
		write,
		sleep,
		write,
	}
	if !reflect.DeepEqual(want, spySleeper.Calls) {
		t.Errorf("wanted calls %v got %v", want, spySleeper.Calls)
	}
})
t.Run("should sleep for the correct amount of time", func(t *testing.T) {
	spyTime := &SpyTime{}
	spySleeper := &SpySleeper{}
	configurableSleeper := &ConfigurableSleeper{1 * time.Second, spyTime.Sleep}
	Countdown(spySleeper, configurableSleeper)
	want := 1 * time.Second
	if spyTime.durationSlept != want {
		t.Errorf("got %v want %v", spyTime.durationSlept, want)
	}

})
}