package bench

import (
	"github.com/gonum/matrix/mat64"
	"github.com/go-gl/mathgl/mgl32"
	"golang.org/x/mobile/exp/f32"
	"testing"
	"math"
)

var f32a = f32.Mat3{
	{100,200,300},
	{100,200,300},
	{100,200,300},
}

var f32b = f32.Mat3{
	{1, 2, 3},
	{1, 1, 1},
	{1, 1, 1},
}

func BenchmarkIntMul(b *testing.B) {
	ei := 27182818
	for i := 0; i < b.N; i++ {
		_ = ei * i
	}
}


func BenchmarkIntToInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = int64(i)
	}
}

func BenchmarkInt64ToInt(b *testing.B) {
	var i int64
	n := int64(b.N)
	for i = 0; i < n; i++ {
		_ = int(i)
	}
}



func BenchmarkInt64Mul(b *testing.B) {
	var x int64 = 2718281828459
	n := int64(b.N)
	var i int64
	for i = 0; i < n; i++ {
		_ = x * i
	}
}


func BenchmarkFloat32Mul(b *testing.B) {
	e32 := float32(math.E)
	for i := 0; i < b.N; i++ {
		_ = e32 * e32
	}
}

func BenchmarkFloat64Mul(b *testing.B) {

	for i := 0; i < b.N; i++ {
		_ = math.E * math.Pi
	}
}

func BenchmarkFloat64ToInt(b *testing.B) {
	f := 2718.281
	for i := 0; i < b.N; i++ {
		_ = int(f)
	}
}

func BenchmarkIntToFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = float64(i)
	}
}



func BenchmarkGonumMul(b *testing.B) {
	m1 := mat64.NewDense(3,3,[]float64{
		100,200,300,
		100,200,300,
		100,200,300})

	// m2 := cartesianMatrix(100,0,45)
	m2 := mat64.NewDense(3,1,[]float64{1,2,3})
	m3 := mat64.NewDense(3,1,make([]float64,3))
	for i := 0; i < b.N; i++ {
		m3.Mul(m1,m2)
	}
}



func BenchmarkMathglMulMxN(b *testing.B) {
	m1 := mgl32.NewMatrixFromData([]float32{
		100,200,300,
		100,200,300,
		100,200,300},
		3,3)
	m2 := mgl32.NewMatrixFromData([]float32{		100,200,300,
		100,200,300,
		100,200,300},
		3,3)
	for i := 0; i < b.N; i++ {
		if m1.MulMxN(m1,m2) == nil {
			b.Fatalf("incorrectly sized matrix")
		}
	}
}


func BenchmarkMathglMulScalar(b *testing.B) {
	m1 := mgl32.NewMatrixFromData([]float32{
		100,200,300,
		100,200,300,
		100,200,300},
		3,3)
	m2 := mgl32.NewMatrix(1,3)
	for i := 0; i < b.N; i++ {
		m1.Mul(m2,100)
	}
}


func BenchmarkXMobileF32Mul(b *testing.B) {
	m := f32.Mat3{
		{0, 1, 2},
		{4, 5, 6},
		{8, 9, 10}}
	for i := 0; i < b.N; i++ {
		m.Mul(&f32a,&f32b)
	}
}

