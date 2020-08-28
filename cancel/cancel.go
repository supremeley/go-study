package cancel

import (
	"context"
	"fmt"
	"time"
)

// Main is a function
func Main() {
	// TestCancel()
}

// cancel_1 is a function
// 只传递了一个信号，有多少chan就需要取消多少次
func cancelA(cancelChan chan struct{}) {
	cancelChan <- struct{}{}
}

// cancel_2 is a function
// close(cancelChan)会使所有chan处于阻塞等待状态的
func cancelB(cancelChan chan struct{}) {
	close(cancelChan)
}

// isCancelled is a function
// func isCancelled(cancelChan chan struct{}) bool {
func isCancelled(ctx context.Context) bool {
	select {
	// case <-cancelChan:
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

// TestCancel is a function
func TestCancel() {
	// cancelChan := make(chan struct{}, 0)
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		// go func(i int, cancelCh chan struct{}) {
		go func(i int, ctx context.Context) {
			for {
				if isCancelled(ctx) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Cancelled")
			// }(i, cancelCh)
		}(i, ctx)
	}
	// cancelB(cancelChan)
	cancel()
	time.Sleep(time.Second * 1)
}
