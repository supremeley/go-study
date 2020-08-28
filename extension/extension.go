package extension

import (
	"fmt"
)

// Main is a function
func Main() {
	// TestDog()
}

// Pet is a struct
type Pet struct {
}

// Speak is a function
func (p *Pet) Speak() {
	fmt.Println("...")
}

// SpeakTo is a function
func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println("", host)
}

// Dog is a function
type Dog struct {
	p *Pet
}

// Speak is a function
func (d *Dog) Speak() {
	fmt.Println("wang")
}

// SpeakTo is a function
func (d *Dog) SpeakTo(host string) {
	d.Speak()
	fmt.Println(" ", host)

}

// TestDog is a function
func TestDog() {
	dog := new(Dog)
	dog.Speak()
}
