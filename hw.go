package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hw#1")
	c := make(chan any)

	cOut := doBatching(c, 2)

	go func() {
		for i := 1; i <= 10; i++ {
			c <- i
			fmt.Println("vvela: ", i)
			time.Sleep(300 * time.Millisecond)
		}
		close(c)

	}()
	for i := range cOut {
		fmt.Println(i) // [1,2] [3,4]
	}
}
func doBatching(c chan any, batchSize int) chan []any {
	resp := make(chan []any)
	go func() {
		defer close(resp)
		batch := make([]any, 0, batchSize)

		for val := range c {
			batch = append(batch, val)
			if len(batch) > batchSize {
				resp <- batch
				batch = make([]any, 0, batchSize)
			}
		}
		if len(batch) > 0 {
			resp <- batch
		}
	}()
	return resp
}
