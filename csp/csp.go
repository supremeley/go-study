package csp

import (
	"errors"
	"fmt"
	"time"
)

// Main is a function
func Main() {
	// TestAsyncService()
	// TestSelect()
}

// service is a function
func service() string {
	time.Sleep(time.Millisecond * 10)
	return "Done"
}

// otherTask is a function
func otherTask() {
	fmt.Println("working on someting else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done")
}

// AsyncService is a function
func AsyncService() chan string {
	retCh := make(chan string, 1)

	go func() {
		ret := service()
		fmt.Println("returnd result")
		retCh <- ret
		fmt.Println("service exited")
	}()

	return retCh
}

// TestAsyncService is a function
func TestAsyncService() {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh)
	time.Sleep(time.Second * 1)
}

// TestSelect is a function
func TestSelect() {
	select {
	case ret := <-AsyncService():
		fmt.Println(ret)
	case <-time.After(time.Millisecond * 100):
		fmt.Println(errors.New("time out"))
	}
}
