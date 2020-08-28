package channel

import (
	"fmt"
	"sync"
)

// Main is a function
func Main() {
	// var wg sync.WaitGroup
	// ch := make(chan int)
	// wg.Add(1)
	// dataProducer(ch, &wg)
	// wg.Add(1)
	// dataReceiver(ch, &wg)
	// wg.Add(1)
	// dataReceiver(ch, &wg)
	// wg.Wait()
}

func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		// ch <- 11
		wg.Done()
	}()
}

func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			if data, ok := <-ch; ok {
				fmt.Println(data)
			} else {
				break
			}
		}
		wg.Done()
	}()
}
