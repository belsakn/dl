package main

import (
	"fmt"
	"log"

	. "gorgonia.org/gorgonia"
)

func main() {
	g := NewGraph()
	var a, b, c *Node
	var err error
	// define the expression
	a = NewScalar(g, Float64, WithName("a"))
	b = NewScalar(g, Float64, WithName("b"))
	c, err = Add(a, b)
	if err != nil {
		log.Fatal(err)
	}
	// create a VM to run the program on
	machine := NewTapeMachine(g)
	// set initial values then run
	Let(a, 1.0)
	Let(b, 2.0)
	if machine.RunAll() != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v \n", c.Value())
	// Output: 3.0
}
