package world

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"happymonday.dev/ray-tracer/src/matrix"
	"happymonday.dev/ray-tracer/src/tuples"
)

func TestTransformationMatrixForDefaultOrientation(t *testing.T) {
	from := tuples.InitPoint(0, 0, 0)
	to := tuples.InitPoint(0, 0, -1)
	up := tuples.InitVector(0, 1, 0)
	assert.True(t, matrix.InitMatrixIdentity(4).Equals(ViewTransformation(from, to, up)))
}

func TestViewTransformationLookingInPositiveZDirection(t *testing.T) {
	from := tuples.InitPoint(0, 0, 0)
	to := tuples.InitPoint(0, 0, 1)
	up := tuples.InitVector(0, 1, 0)
	assert.True(t, matrix.Scaling(-1, 1, -1).Equals(ViewTransformation(from, to, up)))
}

func TestViewTransformationMovesWorld(t *testing.T) {
	from := tuples.InitPoint(0, 0, 8)
	to := tuples.InitPoint(0, 0, 0)
	up := tuples.InitVector(0, 1, 0)
	assert.True(t, matrix.Translation(0, 0, -8).Equals(ViewTransformation(from, to, up)))
}

func TestAnAbritraryViewTransformation(t *testing.T) {
	from := tuples.InitPoint(1, 3, 2)
	to := tuples.InitPoint(4, -2, 8)
	up := tuples.InitVector(1, 1, 0)
	m := matrix.InitMatrix(`
| -0.50709 | 0.50709 | 0.67612 | -2.36643 |
| 0.76772 | 0.60609 | 0.12122 | -2.82843 |
| -0.35857 | 0.59761 | -0.71714 | 0.00000 |
| 0.00000 | 0.00000 | 0.00000 | 1.00000 |
`)
	assert.True(t, m.Equals(ViewTransformation(from, to, up)))
}
