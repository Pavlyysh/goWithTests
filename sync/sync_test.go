package main

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()
		wg := &sync.WaitGroup{}

		wg.Add(wantedCount)
		for i := 0; i < wantedCount; i++ {
			go func() {
				defer wg.Done()
				counter.Inc()
			}()
		}
		wg.Wait()

		assertCounter(t, counter, wantedCount)
	})

}

// Чтобы не создавать копию мьютекса, мы передаем указатель на Counter
func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}
