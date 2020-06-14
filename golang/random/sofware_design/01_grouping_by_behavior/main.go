// --------------------
// Grouping by behavior
// --------------------

// This is an example of using composition and interfaces.
// This is something we want to do in Go.
// This patter does provide a good design principle in a Go program.
// Think of the monkey / banana problem.
// We want to design by what we do, rather than what we are.

// We will group common types by their behavior and not their state.
// What's brilliant about Go is that it doesn't have to be configured ahead of time.
// The compiler automatically identifies interface and behaviors at compile time.
// It means that we can write code today that's compliant with any interface
// that exists today and tomorrow. It doesn't matter where that is declared
// because the compiler can do this on the fly.

package main

import (
	"fmt"
)

// Speaker provides a common behavior for all concrete types
// to follow if they want to be part of this group.
// This is a contract for these concrete types to follow.
type Speaker interface {
	Speak()
}

// Dog contains everything a dog needs
type Dog struct {
	Name       string
	IsMammal   bool
	PackFactor int
}

// Speak implements the Speaker interface for Dog
// so it knows how to speak like a dog.
func (d Dog) Speak() {
	fmt.Println("Woof!",
		"My name is", d.Name,
		", it is", d.IsMammal,
		"I am a mammal with a pack factor of", d.PackFactor)
}

// Cat contains everything a Cat needs.
// A little copy and paste can go a long way. Decoupling, in many cases,
// is a much better option that reusing code.
type Cat struct {
	Name        string
	IsMammal    bool
	ClimbFactor int
}

// Speak knows how to speak like a cat.
// This makes a Cat now part of a concrete type that knows how to speak.
func (c Cat) Speak() {
	fmt.Println("Meow!",
		"My name is", c.Name,
		", it is", c.IsMammal,
		"I am a mammal with a climb factor of", c.ClimbFactor)
}

func main() {
	// Create a list of Animals that know how to speak.
	speakers := []Speaker{
		// Create a dog by initializing Dog attributes.
		Dog{
			Name:       "Bowser",
			IsMammal:   true,
			PackFactor: 5,
		},

		// Create a Cat by initializing Cat attributes.
		Cat{
			Name:        "Milo",
			IsMammal:    true,
			ClimbFactor: 4,
		},
	}

	for _, spk := range speakers {
		spk.Speak()
	}
}

// ---------------------------------
// Guidelines around declaring types
// ---------------------------------

// - Declare types that represent something new or unique. We don't want to create aliases just for readability.
// - Validate that a value of any type is created or used on its own.
// - Embed types not because we need the state but because we need the behavior. If we are not thinking
// about behavior, we are locking ourselves into the design that we cannot grow in the future.
// - Question types that are aliases or abstraction for an existing type.
// - Question types whose sole purpose is to share common state.
