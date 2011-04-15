// Copyright 2009 The GoMatrix Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

import (
	"math"
	"os"
)

/*
Finds the sum of two matrices.
*/
func Sum(A MatrixRO, Bs ...MatrixRO) (C *DenseMatrix) {
	C = MakeDenseCopy(A)
	var err os.Error
	for _, B := range Bs {
		err = C.Add(MakeDenseCopy(B))
		if err != nil {
			break
		}
	}
	if err != nil {
		C = nil
	}
	return
}

/*
Finds the difference between two matrices.
*/
func Difference(A, B MatrixRO) (C *DenseMatrix) {
	C = MakeDenseCopy(A)
	err := C.Subtract(MakeDenseCopy(B))
	if err != nil {
		C = nil
	}
	return
}

/*
Finds the Product of two matrices.
*/
func Product(A MatrixRO, Bs ...MatrixRO) (C *DenseMatrix) {
	C = MakeDenseCopy(A)

	for _, B := range Bs {
		Cm, err := C.Times(B)
		if err != nil {
			return
		}
		C = Cm.(*DenseMatrix)
	}

	return
}

/*
Uses a number of goroutines to do the dot products necessary
for the matrix multiplication in parallel.
*/
func ParallelProduct(A, B MatrixRO) (C *DenseMatrix) {
	if A.Cols() != B.Rows() {
		return nil
	}

	C = Zeros(A.Rows(), B.Cols())

	in := make(chan int)
	quit := make(chan bool)

	dotRowCol := func() {
		for {
			select {
			case i := <-in:
				sums := make([]float64, B.Cols())
				for k := 0; k < A.Cols(); k++ {
					for j := 0; j < B.Cols(); j++ {
						sums[j] += A.Get(i, k) * B.Get(k, j)
					}
				}
				for j := 0; j < B.Cols(); j++ {
					C.Set(i, j, sums[j])
				}
			case <-quit:
				return
			}
		}
	}

	threads := 2

	for i := 0; i < threads; i++ {
		go dotRowCol()
	}

	for i := 0; i < A.Rows(); i++ {
		in <- i
	}

	for i := 0; i < threads; i++ {
		quit <- true
	}

	return
}

/*
Scales a matrix by a scalar.
*/
func Scaled(A MatrixRO, f float64) (B *DenseMatrix) {
	B = MakeDenseCopy(A)
	B.Scale(f)
	return
}

/*
Tests the element-wise equality of the two matrices.
*/
func Equals(A, B MatrixRO) bool {
	if A.Rows() != B.Rows() || A.Cols() != B.Cols() {
		return false
	}
	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			if A.Get(i, j) != B.Get(i, j) {
				return false
			}
		}
	}
	return true
}

/*
Tests to see if the difference between two matrices,
element-wise, exceeds ε.
*/
func ApproxEquals(A, B MatrixRO, ε float64) bool {
	if A.Rows() != B.Rows() || A.Cols() != B.Cols() {
		return false
	}
	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			if math.Fabs(A.Get(i, j)-B.Get(i, j)) > ε {
				return false
			}
		}
	}
	return true
}
