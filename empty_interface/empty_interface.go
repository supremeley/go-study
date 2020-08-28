package emptyinterface

import (
	"fmt"
)

// Main is a struct
func Main() {
	// DoSomething("10")
	// DoSomething(1)
}

// DoSomething is a struct
func DoSomething(p interface{}) {
	switch v := p.(type) {
	case int:
		fmt.Println("int", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("unknow")
	}

}
