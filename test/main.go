package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {
	ann := NewANN(4, 16, 4, 2)
	I := mat.NewVecDense(4, []float64{0.4, 0.9, 1.0, 0.5})

	fmt.Println(ann.Process(I))
}
