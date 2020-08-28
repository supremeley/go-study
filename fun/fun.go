package fun

import (
	"fmt"
	"math/rand"
	"time"
)

// IntConv is a interface
type IntConv func(op int) int

// ran is a function
func ran() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

// timeSpent is a function
func timeSpent(inner IntConv) IntConv {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

// slowFun is a function
func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

// Main is a function
func Main() {
	// a, _ := ran()
	// fmt.Println(a)
	// ts := timeSpent(slowFun)
	// fmt.Println(ts(10))
}
