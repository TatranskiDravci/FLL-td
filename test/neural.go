package main

import (
	"gonum.org/v1/gonum/mat"
)

type ANN struct {
	l   int
	n   int
	ioo int
	W   []*mat.Dense
	b   []*mat.VecDense
	IW  *mat.Dense
	OW  *mat.Dense
	Ib  *mat.VecDense
	Ob  *mat.VecDense
}

func NewANN(l int, n int, ioi int, ioo int) ANN {
	W := make([]*mat.Dense, l-1)
	for i := 0; i < (l - 1); i++ {
		W[i] = mat.NewDense(n, n, SampleData(n, n))
	}

	b := make([]*mat.VecDense, l-1)
	for i := 0; i < (l - 1); i++ {
		b[i] = mat.NewVecDense(n, SampleData(n, 1))
	}

	IW := mat.NewDense(n, ioi, SampleData(n, ioi))
	OW := mat.NewDense(ioo, n, SampleData(ioo, n))

	Ib := mat.NewVecDense(n, SampleData(n, 1))
	Ob := mat.NewVecDense(ioo, SampleData(ioo, 1))

	return ANN{
		n:   n,
		l:   l,
		ioo: ioo,
		W:   W,
		b:   b,
		IW:  IW,
		OW:  OW,
		Ib:  Ib,
		Ob:  Ob,
	}
}

func (ann ANN) Process(I *mat.VecDense) *mat.VecDense {
	N := mat.NewVecDense(ann.n, nil)
	N.MulVec(ann.IW, I)
	N.AddVec(N, ann.Ib)
	N = Sigmoid(N)

	for i := 0; i < (ann.l - 1); i++ {
		N.MulVec(ann.W[i], N)
		N.AddVec(N, ann.b[i])
		N = Sigmoid(N)
	}

	O := mat.NewVecDense(ann.ioo, nil)
	O.MulVec(ann.OW, N)
	O.AddVec(O, ann.Ob)
	return Sigmoid(O)
}
