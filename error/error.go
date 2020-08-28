package error

import (
	"errors"
)

// Main is a function
func Main() {
	// errors.New("n is not a function ")
	// fmt.Println(GetFibonacci(1))

	// if v, err := GetFibonacci(1); err != nil {
	// 	if err == LessThan {
	// 		fmt.Println("less", v)
	// 	}
	// 	if err == MoreThan {
	// 		fmt.Println("more", v)
	// 	}
	// }
}

// ErrLessThan is a function
var ErrLessThan = errors.New("n less")

// ErrMoreThan is a function
var ErrMoreThan = errors.New("n lagger")

// GetFibonacci is a function
func GetFibonacci(n int) ([]int, error) {
	if n < 2 {
		return nil, ErrLessThan
	}
	if n > 100 {
		return nil, ErrMoreThan
	}
	fibList := []int{1, 1}

	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}

	return fibList, nil
}
