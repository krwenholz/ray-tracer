package matrix

import (
	"log"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"happymonday.dev/ray-tracer/src/tuples"
)

func TestConstructingAndInspectingA4x4Matrix(t *testing.T) {
	input := strings.TrimSpace(`
| 1 | 2 | 3 | 4 |
| 5.5 | 6.5 | 7.5 | 8.5 |
| 9 | 10 | 11 | 12 |
| 13.5 | 14.5 | 15.5 | 16.5 |
`)
	m := InitMatrix(input)
	assert.Equal(t, 1.0, m.At(0, 0))
	assert.Equal(t, 4.0, m.At(0, 3))
	assert.Equal(t, 5.5, m.At(1, 0))
	assert.Equal(t, 7.5, m.At(1, 2))
	assert.Equal(t, 11.0, m.At(2, 2))
	assert.Equal(t, 13.5, m.At(3, 0))
	assert.Equal(t, 15.5, m.At(3, 2))
}

func Test2x2Matrix(t *testing.T) {
	input := strings.TrimSpace(`
| -3 | 5 |
| 1 | -2 |
`)
	m := InitMatrix(input)
	assert.Equal(t, -3.0, m.At(0, 0))
	assert.Equal(t, 5.0, m.At(0, 1))
	assert.Equal(t, 1.0, m.At(1, 0))
	assert.Equal(t, -2.0, m.At(1, 1))
}

func Test3x3Matrix(t *testing.T) {
	input := strings.TrimSpace(`
| -3 | 5 | 0 |
| 1 | -2 | -7 |
| 0 | 1 | 1 |
`)
	m := InitMatrix(input)
	assert.Equal(t, -3.0, m.At(0, 0))
	assert.Equal(t, -2.0, m.At(1, 1))
	assert.Equal(t, 1.0, m.At(2, 2))
}

func TestMatrixEqualitySame(t *testing.T) {
	i1 := strings.TrimSpace(`
| 1 | 2 | 3 | 4 |
| 5 | 6 | 7 | 8 |
| 9 | 8 | 7 | 6 |
| 5 | 4 | 3 | 2 |
`)
	i2 := strings.TrimSpace(`
| 1 | 2 | 3 | 4 |
| 5 | 6 | 7 | 8 |
| 9 | 8 | 7 | 6 |
| 5 | 4 | 3 | 2 |
`)
	m1 := InitMatrix(i1)
	m2 := InitMatrix(i2)
	assert.True(t, m1.Equals(m2))
}

func TestMatrixEqualityDifferent(t *testing.T) {
	i1 := strings.TrimSpace(`
| 1 | 2 | 3 | 4 |
| 5 | 6 | 7 | 8 |
| 9 | 8 | 7 | 6 |
| 5 | 4 | 3 | 2 |
`)
	i2 := strings.TrimSpace(`
| 2 | 3 | 4 | 5 |
| 6 | 7 | 8 | 9 |
| 8 | 7 | 6 | 5 |
| 4 | 3 | 2 | 1 |
`)
	m1 := InitMatrix(i1)
	m2 := InitMatrix(i2)
	assert.False(t, m1.Equals(m2))
}

func TestMatrixMultiplication(t *testing.T) {
	i1 := strings.TrimSpace(`
| 1 | 2 | 3 | 4 |
| 5 | 6 | 7 | 8 |
| 9 | 8 | 7 | 6 |
| 5 | 4 | 3 | 2 |
`)
	i2 := strings.TrimSpace(`
| -2 | 1 | 2 | 3 |
| 3 | 2 | 1 | -1 |
| 4 | 3 | 6 | 5 |
| 1 | 2 | 7 | 8 |
`)
	ir := strings.TrimSpace(`
| 20| 22 | 50 | 48 |
| 44| 54 | 114 | 108 |
| 40| 58 | 110 | 102 |
| 16| 26 | 46 | 42 |
`)
	m1 := InitMatrix(i1)
	m2 := InitMatrix(i2)
	mr := InitMatrix(ir)
	assert.True(t, m1.Multiply(m2).Equals(mr))
}

func TestMatrixTupleMultiplication(t *testing.T) {
	i := strings.TrimSpace(`
| 1 | 2 | 3 | 4 |
| 2 | 4 | 4 | 2 |
| 8 | 6 | 4 | 1 |
| 0 | 0 | 0 | 1 |
	`)
	m := InitMatrix(i)
	tup := tuples.InitPoint(1, 2, 3)
	assert.True(t, m.MultiplyTuple(tup).Equals(tuples.InitPoint(18, 24, 33)))
}

func TestIdentityMatrix(t *testing.T) {
	i := strings.TrimSpace(`
| 0 | 1 | 2 | 4 |
| 1 | 2 | 4 | 8 |
| 2 | 4 | 8 | 16 |
| 4 | 8 | 16 | 32 |
	`)
	m := InitMatrix(i)
	assert.True(t, m.Multiply(InitMatrixIdentity(4)).Equals(m))
}

func TestMatrixTranspose(t *testing.T) {
	i := strings.TrimSpace(`
| 0 | 9 | 3 | 0 |
| 9 | 8 | 0 | 8 |
| 1 | 8 | 5 | 3 |
| 0 | 0 | 5 | 8 |
	`)
	m := InitMatrix(i)
	i2 := strings.TrimSpace(`
| 0 | 9 | 1 | 0 |
| 9 | 8 | 8 | 0 |
| 3 | 0 | 5 | 5 |
| 0 | 8 | 3 | 8 |
	`)
	res := InitMatrix(i2)
	assert.True(t, m.Transpose().Equals(res))
}

func TestTransposeIdentity(t *testing.T) {
	assert.True(t, InitMatrixIdentity(4).Transpose().Equals(InitMatrixIdentity(4)))
}

func TestDeterminant(t *testing.T) {
	i := strings.TrimSpace(`
| 1 | 5 |
| -3 | 2 |
	`)
	m := InitMatrix(i)
	assert.Equal(t, 17.0, m.Determinant())
}

func TestSubmatrix2x2(t *testing.T) {
	m := InitMatrix(strings.TrimSpace(`
| 1 | 5 | 0 |
| -3 | 2 | 7 |
| 0 | 6 | -3 |
	`))
	r := InitMatrix(strings.TrimSpace(`
| -3 | 2 |
| 0 | 6 |
	`))
	assert.True(t, m.Submatrix(0, 2).Equals(r))
}

func TestSubmatrix4x4(t *testing.T) {
	m := InitMatrix(strings.TrimSpace(`
| -6 | 1 | 1 | 6 |
| -8 | 5 | 8 | 6 |
| -1 | 0 | 8 | 2 |
| -7 | 1 | -1 | 1 |
	`))
	r := InitMatrix(strings.TrimSpace(`
| -6 | 1 | 6 |
| -8 | 8 | 6 |
| -7 | -1 | 1 |
	`))
	assert.True(t, m.Submatrix(2, 1).Equals(r))
}

func TestMinor3x3(t *testing.T) {
	m := InitMatrix(strings.TrimSpace(`
| 3 | 5 | 0 |
| 2 | -1 | -7 |
| 6 | -1 | 5 |
	`))
	assert.Equal(t, 25.0, m.Minor(1, 0))
}

func TestCofactor3x3(t *testing.T) {
	m := InitMatrix(strings.TrimSpace(`
| 3 | 5 | 0 |
| 2 | -1 | -7 |
| 6 | -1 | 5 |
	`))
	assert.Equal(t, -12.0, m.Cofactor(0, 0))
	assert.Equal(t, -25.0, m.Cofactor(1, 0))
}

func TestDeterminant3x3(t *testing.T) {
	m := InitMatrix(strings.TrimSpace(`
| 1 | 2 | 6 |
| -5 | 8 | -4 |
| 2 | 6 | 4 |
	`))
	assert.Equal(t, 56.0, m.Cofactor(0, 0))
	assert.Equal(t, 12.0, m.Cofactor(0, 1))
	assert.Equal(t, -46.0, m.Cofactor(0, 2))
	assert.Equal(t, -196.0, m.Determinant())
}

func TestDeterminant4x4(t *testing.T) {
	m := InitMatrix(strings.TrimSpace(`
| -2 | -8 | 3 | 5 |
| -3 | 1 | 7 | 3 |
| 1 | 2 | -9 | 6 |
| -6 | 7 | 7 | -9 |
	`))
	assert.Equal(t, 690.0, m.Cofactor(0, 0))
	assert.Equal(t, 447.0, m.Cofactor(0, 1))
	assert.Equal(t, 210.0, m.Cofactor(0, 2))
	assert.Equal(t, 51.0, m.Cofactor(0, 3))
	assert.Equal(t, -4071.0, m.Determinant())
}

func TestInvertibility(t *testing.T) {
	type opt struct {
		m           *Matrix
		determinant float64
		invertible  bool
	}
	opts := []opt{
		{
			InitMatrix(strings.TrimSpace(`
| 6 | 4 | 4 | 4 |
| 5 | 5 | 7 | 6 |
| 4 | -9 | 3 | -7 |
| 9 | 1 | 7 | -6 |
	`)),
			-2120,
			true,
		},
		{
			InitMatrix(strings.TrimSpace(`
| -4 | 2 | -2 | -3 |
| 9 | 6 | 2 | 6 |
| 0 | -5 | 1 | -5 |
| 0 | 0 | 0 | 0 |
	`)),
			0,
			false,
		},
	}
	for _, o := range opts {
		assert.Equal(t, o.determinant, o.m.Determinant())
		assert.Equal(t, o.invertible, o.m.IsInvertible())
	}
}

func TestInverse(t *testing.T) {
	type opt struct {
		m       *Matrix
		inverse *Matrix
	}
	opts := []opt{
		{
			InitMatrix(strings.TrimSpace(`
| -5 | 2 | 6 | -8 |
| 1 | -5 | 1 | 8 |
| 7 | 7 | -6 | -7 |
| 1 | -3 | 7 | 4 |
	`)),
			InitMatrix(strings.TrimSpace(`
| 0.21805 | 0.45113 | 0.24060 | -0.04511 |
| -0.80827 | -1.45677 | -0.44361 | 0.52068 |
| -0.07895 | -0.22368 | -0.05263 | 0.19737 |
| -0.52256 | -0.81391 | -0.30075 | 0.30639 |
`)),
		},
		{
			InitMatrix(strings.TrimSpace(`
| 8 | -5 | 9 | 2 |
| 7 | 5 | 6 | 1 |
| -6 | 0 | 9 | 6 |
| -3 | 0 | -9 | -4 |
	`)),
			InitMatrix(strings.TrimSpace(`
| -0.15385 | -0.15385 | -0.28205 | -0.53846 |
| -0.07692 | 0.12308 | 0.02564 | 0.03077 |
| 0.35897 | 0.35897 | 0.43590 | 0.92308 |
| -0.69231 | -0.69231 | -0.76923 | -1.92308 |
`)),
		},
		{
			InitMatrix(strings.TrimSpace(`
| 9 | 3 | 0 | 9 |
| -5 | -2 | -6 | -3 |
| -4 | 9 | 6 | 4 |
| -7 | 6 | 6 | 2 |
	`)),
			InitMatrix(strings.TrimSpace(`
| -0.04074 | -0.07778 | 0.14444 | -0.22222 |
| -0.07778 | 0.03333 | 0.36667 | -0.33333 |
| -0.02901 | -0.14630 | -0.10926 | 0.12963 |
| 0.17778 | 0.06667 | -0.26667 | 0.33333 |
`)),
		},
	}
	for _, o := range opts {
		log.Println(o.inverse)
		log.Println(o.m.Inverse())
		assert.True(t, o.inverse.Equals(o.m.Inverse()))
	}
}

func TestMatrixMultiplicationInverted(t *testing.T) {
	m1 := InitMatrix(strings.TrimSpace(`
| 3 | -9 | 7 | 3 |
| 3 | -8 | 2 | -9 |
| -4 | 4 | 4 | 1 |
| -6 | 5 | -1 | 1 |
`))
	m2 := InitMatrix(strings.TrimSpace(`
| 8 | 2 | 2 | 2 |
| 3 | -1 | 7 | 0 |
| 7 | 0 | 5 | 4 |
| 6 | -2 | 0 | 5 |
`))
	assert.True(t, m1.Multiply(m2).Multiply(m2.Inverse()).Equals(m1))
}
