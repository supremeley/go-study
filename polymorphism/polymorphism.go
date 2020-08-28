package polymorphism

import (
	"fmt"
)

// Codeing is a struct
type Codeing string

// Programmer is a struct
type Programmer interface {
	Write() Codeing
}

// GoProgrammer is a struct
type GoProgrammer struct{}

// Write is a struct
func (p *GoProgrammer) Write() Codeing {
	return "GoProgrammer"
}

// JavaProgrammer is a struct
type JavaProgrammer struct{}

// Write is a struct
func (p *JavaProgrammer) Write() Codeing {
	return "JavaProgrammer"
}

// writeFirstProgram is a struct
func writeFirstProgram(p Programmer) {
	fmt.Printf("%T %v\n", p, p.Write())
}

// Main is a struct
func Main() {
	// goProm := new(GoProgrammer)
	// javaProm := new(JavaProgrammer)

	// writeFirstProgram(goProm)
	// writeFirstProgram(javaProm)
}
