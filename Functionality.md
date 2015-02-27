Summary of functionality expected from a matrix type

# Introduction #

X = already implemented

XT = implemented and has a test

- = do not implement

? = maybe implement, maybe not

blank = not yet done

|Functionality | DenseMatrix | PivotMatrix | SparseMatrix|
|:-------------|:------------|:------------|:------------|
|MatrixRO interface | X           | X           | X           |
|Matrix interface | X           | -           | X           |
|String        | X           | X           | X           |
|GetMatrix(reference) | XT          | ?           | XT          |
|Indices       | ?           | X           | ?           |
|L (lower)     | XT          | ?           | X           |
|U (lower)     | XT          | ?           | X           |
|Copy          | X           | X           | X           |
|Augment       | XT          | ?           | XT          |
|Stack         | XT          | ?           | XT          |
|Plus          | X           | ?           | X           |
|Minus         | X           | ?           | X           |
|Add (in place) | XT          | -           | XT          |
|Subtract (in place) | XT          | -           | XT          |
|Times         | XT          | X           | XT          |
|ElementMult   | XT          | -           | XT          |
|Scale (in place) | XT          | -           | X           |
|ScaleMatrix (in place) | XT          | -           | X           |
|Symmetric     | XT          | X           | X           |
|SwapRows      | X           | X           | X           |
|ScaleRow      | X           | -           | X           |
|ScaleAddRow   | X           | -           | X           |
|Inverse       | XT          | X           | ?           |
|Det           | XT          | X           | ?           |
|Trace         | X           | X           | X           |
|OneNorm       | X           | X           | X           |
|TwoNorm       | X           | X           | X           |
|InfinityNorm  | X           | X           | X           |
|Transpose     | X           | X           | X           |
|Solve         | XT          | X           | ?           |
|Cholesky      | XT          | ?           | ?           |
|LU            | XT          | ?           | ?           |
|QR            | XT          | ?           | ?           |
|Eigen         | XT          | ?           | ?           |
|DenseMatrix   | -           | X           | X           |
|PivotMatrix   | -           | -           | -           |
|SparseMatrix  | X           | X           | -           |