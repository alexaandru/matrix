package matrix

import (
    u "github.com/alexaandru/utils"
    "strconv"
    "strings"
)

// Vector holds a list of ints. Used for rows (may as well be for colums).
type Vector []int

// Matrix the good old 2D matrix with rows and columns.
type Matrix []Vector

// Matrix3d holds a 3d matrix, an extension upon Matrix by adding one more dimension.
type Matrix3d []Matrix

// CellInitFunc holds a callback used for initializing matrix cells.
type CellInitFunc func(int) int

const (
    _   = iota
    // Down ... guess
    Down
    // Right ... guess
    Right
    // Diag ... guess
    Diag
)

func (v Vector) String() (str string) {
    str = ""
    for _, v := range v {
        str = str + strconv.Itoa(v) + " "
    }

    return strings.Trim(str, " ")
}

// FIXME: Should use the corresponding funcs from util, BUT ONLY AFTER WE ADD TESTS HERE FIRST
func (v Vector) Max() (m int) {
    m = v[0]
    for _, v := range v[1:] {
        if v >= m {
            m = v
        }
    }

    return
}

// FIXME: Should use the corresponding funcs from util, BUT ONLY AFTER WE ADD TESTS HERE FIRST
func (v Vector) MaxIndex() (i int) {
    m := v[0]
    for k, v := range v[1:] {
        if v >= m {
            m, i = v, k+1
        }
    }

    return
}

// FIXME: Should use the corresponding funcs from util, BUT ONLY AFTER WE ADD TESTS HERE FIRST
func (v Vector) Min() (m int) {
    m = v[0]
    for _, v := range v[1:] {
        if v <= m {
            m = v
        }
    }

    return
}

// FIXME: Should use the corresponding funcs from util, BUT ONLY AFTER WE ADD TESTS HERE FIRST
func (v Vector) MinIndex() (i int) {
    m := v[0]
    for k, v := range v[1:] {
        if v <= m {
            m, i = v, k+1
        }
    }

    return
}

func (m Matrix) String() (str string) {
    str = "(" + strconv.Itoa(len(m)) + "x" + strconv.Itoa(len(m[0])) + ")\n"
    for _, v := range m {
        str = str + v.String() + "\n"
    }

    return strings.Trim(str, " ")
}

func NewVector(n int, inits ...CellInitFunc) (out Vector) {
    var cellFn CellInitFunc
    if len(inits) > 0 {
        cellFn = inits[0]
    } else {
        cellFn = u.ConstIntFunc(0)
    }

    out = make(Vector, n)
    for i := 0; i < n; i++ {
        out[i] = cellFn(i)
    }

    return
}

func NewMatrix(m, n int, inits ...CellInitFunc) (out Matrix) {
    var rowFn, colFn CellInitFunc

    if len(inits) == 2 {
        rowFn, colFn = inits[0], inits[1]
    } else if len(inits) == 1 {
        rowFn, colFn = inits[0], u.IdentIntFunc
    } else {
        rowFn, colFn = u.IdentIntFunc, u.IdentIntFunc
    }

    row := make(Vector, n)
    for i := 0; i < n; i++ {
        row[i] = colFn(i)
    }

    out = make(Matrix, m)
    out[0] = row

    for i := 1; i < m; i++ {
        row = make(Vector, n)
        row[0] = rowFn(i)
        out[i] = row
    }

    return
}

func NewMatrix3d(m, n, p int, inits ...CellInitFunc) (out Matrix3d) {
    out = make(Matrix3d, p)
    out[0] = NewMatrix(m, n, u.MultIntFunc(-1), u.MultIntFunc(-1))
    zero := u.ConstIntFunc(0)
    for i := 1; i < p; i++ {
        out[i] = NewMatrix(m, n, zero, zero)
        out[i][0][0] = -i
    }

    return
}
