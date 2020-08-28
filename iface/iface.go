package iface

import (
	"fmt"
)

// Programmer is a interface
type Programmer interface {
	WriteHelloWorld() string
}

// GoProgrammer is a struct
type GoProgrammer struct {
}

// Main is a function
func Main() {
	// TestClient()
}

// WriteHelloWorld is a function
func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"hello\")"
}

// TestClient is a function
func TestClient() {
	var p Programmer
	p = new(GoProgrammer)
	fmt.Println(p.WriteHelloWorld())
}
