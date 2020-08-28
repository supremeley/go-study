package pool

import (
	"errors"
	"fmt"
	"time"
)

// Main is a funciton
func Main() {
	// TestObjPool()
}

// Reusableobj is a funciton
type Reusableobj struct {
}

// ObjPool is a funciton
type ObjPool struct {
	bufChan chan *Reusableobj
}

// NewObjPool is a funciton
func NewObjPool(num int) *ObjPool {
	objPool := ObjPool{}
	objPool.bufChan = make(chan *Reusableobj, num)
	for i := 0; i < num; i++ {
		objPool.bufChan <- &Reusableobj{}
	}
	return &objPool
}

// GetObj is a funciton
func (p *ObjPool) GetObj(timeout time.Duration) (*Reusableobj, error) {
	// func (p *ObjPool) GetObj(timeout time.Duration) (interface{}, error) {
	select {
	case ret := <-p.bufChan:
		return ret, nil
	case <-time.After(timeout):
		return nil, errors.New("time out")
	}
}

// ReleaseObj is a funciton
func (p *ObjPool) ReleaseObj(obj *Reusableobj) error {
	select {
	case p.bufChan <- obj:
		return nil
	default:
		return errors.New("overflow")
	}
}

// TestObjPool is a funciton
func TestObjPool() {
	pool := NewObjPool(10)
	if err := pool.ReleaseObj(&Reusableobj{}); err != nil {
		fmt.Println(err)
	}
	for i := 0; i < 11; i++ {
		if v, err := pool.GetObj(time.Second * 1); err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%T\n", v)
			// if err := pool.ReleaseObj(v); err != nil {
			// 	fmt.Println(err)
			// }
		}
	}
	fmt.Println("Done")
}
