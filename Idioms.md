# Introduction #

Examples of efficient ways to use gomatrix


# Design Goals #

Go is a systems language. Therefore, people doing things in go are likely to be interested in writing extremely efficient code. Because of this, gomatrix puts efficiency over convenience (of notation).

# Basics #

You will need to import the ` matrix ` package.

For the examples below, the ` DenseMatrix ` structure is used. The same ideas work with other matrix types.

To create a new 4x4 matrix:

```
A := Zeros(4, 4);
```

To create a 3x1 matrix from existing data:

```
A := MakeDenseMatrix([]float64{1, 2, 3}, 3, 1);
```

gomatrix uses densely packed arrays in the ` DenseMatrix ` struct: a single one dimensional array holds all the data. If there are R rows and C columns, the first C elements hold the data in the first row, the 2nd C elements hold the data in the 2nd row, etc.

To get and set a matrix's element at the 2nd row, 3rd column:

```
v := A.Get(2, 3);
A.Set(2, 3, v+1);
```

To use the various matrix functions provided for ` DenseMatrix `:
```
C, err := A.Plus(B);
```
` C ` will be the result of the operation, and ` err.OK() ` will be ` true ` if the operation succeeded. If the operation failed, ` C ` will be ` nil ` and ` err.ErrorCode() ` will be an integer for what went wrong, with a message in ` err.String() `. If there is an exceptional, but legal, case, ` C ` will not be ` nil ` but there will be information in ` err `.

For many matrix operations there will be two versions - a general one and an optimized one for the same type.
```
func foo(A *DenseMatrix, B *SparseMatrix) {
  A.Times(B);//Does not take advantage of B's sparseness, but does know about A's internals
  A.TimesDense(A);//optimized for multiplying two dense matrices
  A.TimesDense(B);//will not compile
  B.Times(A);//takes advantage of B's sparseness to avoid looking up all of A's entries
  B.TimesSparse(B);//knows both are sparse, acts accordingly
  B.TimesSparse(A);//will not compile
}
```

The library does not know how to best perform the operation. Perhaps the result should be sparse, perhaps not, depending on the contents of ` A ` and ` B `. The programmer needs to make a conversion explicitly:
```
func foo(A *DenseMatrix, B *SparseMatrix) {
  A.Times(B.DenseMatrix());//returns a dense matrix
  A.TimesDense(B.DenseMatrix());//returns a dense matrix
  A.SparseMatrix().Times(B);//returns a sparse matrix
  A.SparseMatrix().TimesSparse(B);//returns a sparse matrix
}
```

The global arithmetic functions always return ` DenseMatrix `s.
```
func foo(A MatrixRO, B MatrixRO) {
  C := Product(A, B);// C is a DenseMatrix
  D := Sum(A, B);// D is a DenseMatrix
}
```

But generally speaking, if you want efficient code, work directly with the type you need.