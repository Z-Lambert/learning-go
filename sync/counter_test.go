package sync

import (
	"sync"
	"testing"
)



func TestCounter(t *testing.T) {
	t.Run("incrementing the counter three times gives three", func (t *testing.T) {
		counter := Counter{}
		
		counter.Inc()
		counter.Inc()
		counter.Inc()

		expect := 3
		got:= counter.Value()

		if(got != expect) {
			t.Errorf("got %d, expected %d", got, expect)
		}

	})
	t.Run("it runs safely concurrently", func (t *testing.T){
		counter := Counter{}
		iterations := 1000
		var wg sync.WaitGroup
		wg.Add(iterations)

		for i:= 0; i < iterations; i++ {
			go func () {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		got := counter.Value()

		if(iterations != got){
			t.Errorf("got %d, expected %d", got, iterations)
		}
	})
}