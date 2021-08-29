package main

import (
	"math"
	"math/rand"

	"gonum.org/v1/gonum/mat"
)

func SampleData(x int, y int) []float64 {
	data := make([]float64, x*y)
	for i := range data {
		data[i] = rand.NormFloat64()
	}

	return data
}

func Sigmoid(u *mat.VecDense) *mat.VecDense {
	data := make([]float64, u.Len())
	for i := range data {
		data[i] = 2/(1+math.Exp(-u.AtVec(i))) - 1
	}
	return mat.NewVecDense(u.Len(), data)
}
