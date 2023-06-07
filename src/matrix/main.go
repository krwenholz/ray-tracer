package matrix

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"happymonday.dev/ray-tracer/src/maths"
	"happymonday.dev/ray-tracer/src/tuples"
)

type Matrix struct {
	data   [][]float64
	height int
	width  int
}

func InitEmptyMatrix(height, width int) *Matrix {
	data := make([][]float64, height)
	for i := range data {
		data[i] = make([]float64, width)
	}
	return &Matrix{data, height, width}
}

func InitMatrix(s string) *Matrix {
	s1 := strings.Split(strings.TrimSpace(s), "\n")
	rows := [][]float64{}
	for _, r := range s1 {
		cols := strings.Split(r, "|")
		row := []float64{}
		for _, c := range cols {
			c = strings.TrimSpace(c)
			if c != "" {
				f, err := strconv.ParseFloat(c, 64)
				if err == nil {
					row = append(row, f)
					continue
				}
				i, err := strconv.ParseInt(c, 10, 64)
				if err != nil {
					log.Fatal("Failed to parse matrix input ", c, s)
				}
				row = append(row, float64(i))
			}
		}
		rows = append(rows, row)
	}
	return &Matrix{
		data:   rows,
		height: len(rows),
		width:  len(rows[0]),
	}
}

func InitMatrixIdentity(s int) *Matrix {
	rows := [][]float64{}
	for r := 0; r < s; r++ {
		row := []float64{}
		for c := 0; c < s; c++ {
			v := 0.0
			if c == r {
				v = 1.0
			}
			row = append(row, v)
		}
		rows = append(rows, row)
	}
	return &Matrix{
		data:   rows,
		height: s,
		width:  s,
	}
}

func (m *Matrix) String() string {
	rows := []string{}
	for _, r := range m.data {
		nums := []string{}
		for _, n := range r {
			nums = append(nums, fmt.Sprintf("%5g", n))
		}
		rows = append(rows, strings.Join(nums, " | "))
	}
	return strings.Join(rows, "\n")
}

func (m *Matrix) At(row, col int) float64 {
	return m.data[row][col]
}

func (m *Matrix) Set(row, col int, val float64) {
	m.data[row][col] = val
}

func (m *Matrix) Equals(m2 *Matrix) bool {
	if m.height != m2.height && m.width != m2.width {
		return false
	}
	for i := 0; i < m.height; i++ {
		for j := 0; j < m.width; j++ {
			if !maths.FuzzyEquals(m.At(i, j), m2.At(i, j)) {
				return false
			}
		}
	}
	return true
}

func (m *Matrix) Multiply(m2 *Matrix) *Matrix {
	if m.width != m2.height {
		log.Fatal("Matrix multiplication with invalid dimensions attempted", m.height, m2.width)
	}
	res := InitEmptyMatrix(m.height, m2.width)
	for i := 0; i < m.height; i++ {
		for j := 0; j < m2.width; j++ {
			val := 0.0
			for k := 0; k < m.width; k++ {
				val += m.At(i, k) * m2.At(k, j)
			}
			res.Set(i, j, val)
		}
	}
	return res
}

func (m *Matrix) MultiplyTuple(t *tuples.Tuple) *tuples.Tuple {
	m2 := m.Multiply(&Matrix{[][]float64{{t.X}, {t.Y}, {t.Z}, {t.W}}, 4, 1})
	return &tuples.Tuple{X: m2.At(0, 0), Y: m2.At(1, 0), Z: m2.At(2, 0), W: m2.At(3, 0)}
}

func (m *Matrix) Transpose() *Matrix {
	res := InitEmptyMatrix(m.height, m.width)
	for i := 0; i < m.height; i++ {
		for j := 0; j < m.width; j++ {
			res.Set(i, j, m.At(j, i))
		}
	}
	return res
}

func (m *Matrix) Submatrix(r, c int) *Matrix {
	res := InitEmptyMatrix(m.height-1, m.width-1)
	for i := 0; i < m.height; i++ {
		for j := 0; j < m.width; j++ {
			if i == r || j == c {
				continue
			}
			ni := i
			nj := j
			if ni > r {
				ni -= 1
			}
			if nj > c {
				nj -= 1
			}
			res.Set(ni, nj, m.At(i, j))
		}
	}
	return res
}

func (m *Matrix) Determinant() float64 {
	if m.height == 2 {
		return m.At(0, 0)*m.At(1, 1) - m.At(0, 1)*m.At(1, 0)
	}

	det := 0.0
	for i := 0; i < m.width; i++ {
		det += m.At(0, i) * m.Cofactor(0, i)
	}

	return det
}

func (m *Matrix) Minor(r, c int) float64 {
	return m.Submatrix(r, c).Determinant()
}

func (m *Matrix) Cofactor(r, c int) float64 {
	s := 1.0
	if (c+r)%2 == 1 {
		s = -1
	}
	return s * m.Minor(r, c)
}

func (m *Matrix) IsInvertible() bool {
	return m.Determinant() != 0
}

func (m *Matrix) Inverse() *Matrix {
	if !m.IsInvertible() {
		log.Fatal("Attempted inverting a non-invertible matrix")
	}
	cofactors := InitEmptyMatrix(m.height, m.width)
	for r := 0; r < m.height; r++ {
		for c := 0; c < m.width; c++ {
			cofactors.Set(r, c, m.Cofactor(r, c))
		}
	}
	determinant := m.Determinant()
	inverse := InitEmptyMatrix(m.height, m.width)
	for r := 0; r < m.height; r++ {
		for c := 0; c < m.width; c++ {
			inverse.Set(r, c, cofactors.At(c, r)/determinant)
		}
	}

	return inverse
}
