package structs

import (
	"fmt"
	"unsafe"
)

// Employee is a struct
type Employee struct {
	ID   int
	Name string
	Age  int
}

// Main is a function
func Main() {
	// e1 := Employee{1, "Bob", 22}
	// e2 := Employee{Name: "Marry", Age: 30}
	// e3 := new(Employee)
	// e3.ID = 3
	// e3.Name = "Alice"
	// e3.Age = 23
	// fmt.Println(e1, e2, e3)
	// fmt.Printf("e is %T", e2)
	// fmt.Printf("e is %T", e3)
	// structOperations()
}

func (e Employee) String() string {
	// 在实例方法被调用时，会对实例成员复制
	fmt.Printf("address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("id:%d/name:%s/age:%d", e.ID, e.Name, e.Age)
}

// func (e *Employee) String() string {
// 避免内存拷贝推荐使用
// 	fmt.Printf("address is %x", unsafe.Pointer(&e.Name))
// 	return fmt.Sprintf("id:%d/name:%s/age:%d", e.ID, e.Name, e.Age)
// }

func structOperations() {
	e := &Employee{1, "Ann", 22}
	fmt.Printf("address is %x\n", unsafe.Pointer(&e.Name))

	fmt.Println(e.String())
}
