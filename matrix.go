package matrix

import (
    "github.com/alexaandru/utils"
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

func (v Vector) Max() int {
    return utils.MaxInt(v...)
}

func (v Vector) MaxIndex() int {
    return utils.MaxIntIndex(v...)
}

func (v Vector) Min() int {
    return utils.MinInt(v...)
}

func (v Vector) MinIndex() int {
    return utils.MinIntIndex(v...)
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
        cellFn = utils.ConstIntFunc(0)
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
        rowFn, colFn = inits[0], utils.IdentIntFunc
    } else {
        rowFn, colFn = utils.IdentIntFunc, utils.IdentIntFunc
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
    out[0] = NewMatrix(m, n, utils.MultIntFunc(-1), utils.MultIntFunc(-1))
    zero := utils.ConstIntFunc(0)
    for i := 1; i < p; i++ {
        out[i] = NewMatrix(m, n, zero, zero)
        out[i][0][0] = -i
    }

    return
}
