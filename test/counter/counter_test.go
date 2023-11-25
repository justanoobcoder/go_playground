package counter

import (
	"go_playground/src/counter"
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		cnter := counter.Counter{}
		cnter.Inc()
		cnter.Inc()
		cnter.Inc()

		expected := 3
		actual := cnter.Value()

		assertCounter(t, actual, expected)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		expected := 1000
		cnter := counter.Counter{}

		var wg sync.WaitGroup
		wg.Add(expected)
		for i := 0; i < expected; i++ {
			go func() {
				cnter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()
		assertCounter(t, cnter.Value(), expected)
	})
}

func assertCounter(t testing.TB, actual int, expected int) {
	t.Helper()
	if actual != expected {
		t.Errorf("got %d, want %d", actual, expected)
	}
}
