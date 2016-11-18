package chassis

import (
	"github.com/gonum/matrix/mat64"
	"fmt"
	"math"
)

func toMatrix(x,y, angular float64) *mat64.Dense {
	return mat64.NewDense(3,1,[]float64{x,y,angular})
}


func logMatrix(s string, m mat64.Matrix) {
	fmt.Println(s)
	fmt.Printf("%0.4v\n", mat64.Formatted(m, mat64.Prefix(" ")))
	fmt.Println("")
}


func matrixAbsCopy(m *mat64.Dense) *mat64.Dense{
	acm := mat64.DenseCopyOf(m)
	acm.Apply(func(r,c int, val float64) float64 {
		return math.Abs(val)
	},m )
	return acm
}

func newMatrix(r,c int) *mat64.Dense {
	return mat64.NewDense(r, c,make([]float64,r*c))
}


func matrixMulEach(m *mat64.Dense,n float64) *mat64.Dense{
	acm := mat64.DenseCopyOf(m)
	acm.Apply(func(r,c int, val float64) float64 {
		return val*n
	},m )
	return acm
}

func matrixMul(m1,m2 *mat64.Dense) *mat64.Dense{
	m3 := mat64.DenseCopyOf(m2)
	m3.Mul(m1,m2)
	return m3
}

func matrixDivElem(m1,m2 *mat64.Dense) *mat64.Dense{
	m3 := mat64.DenseCopyOf(m2)
	m3.DivElem(m1,m2)
	return m3
}

func cartesianMatrix(radius,direction,angular float64) *mat64.Dense {
	m := mat64.NewDense(3,1,[]float64{
		math.Cos(direction) * radius,
		math.Sin(direction) * radius,
		angular,
	})

	return m;
}
