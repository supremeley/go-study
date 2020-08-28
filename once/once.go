package once

import (
	"fmt"
	"sync"
	"unsafe"
)

// Main is a function
func Main() {
	// TestGetSingletonObj()
}

// Singleton is a function
type Singleton struct{}

var singleInstance *Singleton
var once sync.Once

// GetSingletonObj is a function
// singleInstance只被创建了一次，适合单列模式
func GetSingletonObj() *Singleton {
	once.Do(func() {
		fmt.Println("Create Obj")
		singleInstance = new(Singleton)
	})
	return singleInstance
}

// TestGetSingletonObj is a function
func TestGetSingletonObj() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletonObj()
			fmt.Printf("%x\n", unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}
