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
	log.Println(m.MultiplyTuple(tup))
	assert.True(t, m.MultiplyTuple(tup).Eq(tuples.InitPoint(18, 24, 33)))
}
