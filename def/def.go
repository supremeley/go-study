package def

import (
	"fmt"
)

// Sum is a function
func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

// Main is a function
func Main() {
	// fmt.Println(Sum(1, 2, 3, 4, 5))
	// fmt.Println(Sum(1, 2, 3, 4))
	// TestDefer()
}

// Clear is a function
func Clear() {
	fmt.Println("clear")
}

// TestDefer is a function
func TestDefer() {
	defer Clear()
	fmt.Println("test")
	// panic("err")
}
