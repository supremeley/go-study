package firstres

import (
	"fmt"
	"time"
)

// Main is a function
func Main() {
	// fmt.Println("before:", runtime.NumGoroutine())
	// // fmt.Println(FirstResponse())
	// fmt.Println(AllResponse())
	// time.Sleep(time.Second * 1)
	// fmt.Println("after:", runtime.NumGoroutine())
}

// runTask is a function
func runTask(id int) string {
	time.Sleep(10 * time.Millisecond * 10)
	return fmt.Sprintf("The result is from %d", id)

}

// FirstResponse is a function
func FirstResponse() string {
	n := 10
	ch := make(chan string, n)
	for i := 0; i < n; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	return <-ch
}

// AllResponse is a function
func AllResponse() string {
	n := 10
	ch := make(chan string, n)
	for i := 0; i < n; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	finalRet := ""
	for j := 0; j < n; j++ {
		finalRet += <-ch + "\n"
	}
	return finalRet
}
