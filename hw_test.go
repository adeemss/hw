package main

import "testing"

func BenchmarkDoBatching(b *testing.B) {
	c := make(chan any, b.N)

	go func() {
		for i := 0; i <= b.N; i++ {
			c <- i
		}
		close(c)

	}()
	resp := doBatching(c, 18888)
	for range resp {

	}
}
