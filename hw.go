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


	
	d := make(chan any)
	dOut := doB(d, 3)
	go func() {
		for i := 1; i <= 10; i++ {
			d <- i
			fmt.Println("vvela: ", i)
			time.Sleep(300 * time.Millisecond)
		}
		close(d)

	}()
	for di := range dOut {
		fmt.Println("dddd", di)
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
// case just for me w select
func doB(d chan any, batchSize int) chan []any {
	resp := make(chan []any)
	var batch []any
	go func() {
		defer close(resp)
		for {
			select {
			case val, ok := <-d:
				if !ok {
					if len(batch) > 0 {
						resp <- batch
					}
					return
				}
				if len(batch) == batchSize {
					resp <- batch
					batch = nil
				}
				batch = append(batch, val)
			}
		}
	}()
	return resp
}
// case w WaitGroup 
/*func main() {
	c := make(chan any)
	d := make(chan any)

	cOut := doBatching(c, 2)
	dOut := doB(d, 2)

	var wg sync.WaitGroup
	wg.Add(4)

	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			c <- i
			fmt.Println("c channel is ", i)
			time.Sleep(1 * time.Second)
		}
		close(c)
	}()
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			d <- i
			fmt.Println("d channel is ", i)
			time.Sleep(1 * time.Second)
		}
		close(d)
	}()

	go func() {
		defer wg.Done()
		for i := range cOut {
			fmt.Println("c batch ", i)
		}
	}()
	go func() {
		defer wg.Done()
		for i := range dOut {
			fmt.Println("d batch ", i)
		}
	}()
	wg.Wait()

}*/
