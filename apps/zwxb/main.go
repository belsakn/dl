package main

import (
	"fmt"
	"io/ioutil"
	"log"

	G "gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
)

// Calculate z = Wx + b
func main() {
	g := G.NewGraph()

	matB := []float64{0.9, 0.7, 0.4, 0.2}
	matT := tensor.New(tensor.WithBacking(matB), tensor.WithShape(2, 2))
	mat := G.NewMatrix(g,
		tensor.Float64,
		G.WithName("W"),
		G.WithShape(2, 2),
		G.WithValue(matT),
	)

	vecB := []float64{5, 7}
	vecT := tensor.New(tensor.WithBacking(vecB), tensor.WithShape(2))
	vec := G.NewVector(g,
		tensor.Float64,
		G.WithName("x"),
		G.WithShape(2),
		G.WithValue(vecT),
	)

	b := G.NewScalar(g,
		tensor.Float64,
		G.WithName("b"),
		G.WithValue(3.0),
	)

	a, err := G.Mul(mat, vec)
	if err != nil {
		log.Fatal(err)
	}

	z, err := G.Add(a, b)
	if err != nil {
		log.Fatal(err)
	}

	machine := G.NewTapeMachine(g)
	if machine.RunAll() != nil {
		log.Fatal(err)
	}
	fmt.Println(z.Value().Data())
	// Output: [9.4 3.4]

	// Visualize
	ioutil.WriteFile("simple_graph.dot", []byte(g.ToDot()), 0644)
}
