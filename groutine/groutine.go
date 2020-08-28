package groutine

// Main is a function
func Main() {
	// for i := 0; i < 10; i++ {
	// 	// 协程调用顺序跟方法顺序无关
	// 	go func(i int) {
	// 		fmt.Println(i)
	// 	}(i)

	// 	// 协程内会共用一个i，需要用锁
	// 	// go func() {
	// 	// 	fmt.Println(i)
	// 	// }()
	// }
	// time.Sleep(time.Millisecond * 50)

	// 并发的执行 不是安全的协程 会丢失部分数据操作
	// counter := 0
	// for i := 0; i < 5000; i++ {
	// 	go func() {
	// 		counter++
	// 	}()
	// }
	// time.Sleep(1 * time.Second)
	// fmt.Println(counter)

	// 不确定协程量
	// var mut sync.Mutex
	// count := 0
	// for i := 0; i < 5000; i++ {
	// 	go func() {
	// 		defer func() {
	// 			mut.Unlock()
	// 		}()
	// 		mut.Lock()
	// 		count++
	// 	}()
	// }
	// time.Sleep(1 * time.Second)
	// fmt.Println(count)

	// var mut sync.Mutex
	// var wg sync.WaitGroup
	// count := 0
	// for i := 0; i < 5000; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		defer func() {
	// 			mut.Unlock()
	// 		}()
	// 		mut.Lock()
	// 		count++
	// 		wg.Done()
	// 	}()
	// }
	// wg.Wait()
	// fmt.Println(count)
}
