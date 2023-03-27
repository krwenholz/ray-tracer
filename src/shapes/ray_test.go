package shapes

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"happymonday.dev/ray-tracer/src/matrix"
	"happymonday.dev/ray-tracer/src/tuples"
)

func TestInitializeRay(t *testing.T) {
	origin := tuples.InitPoint(1, 2, 3)
	direction := tuples.InitVector(4, 5, 6)
	r := InitRay(origin, direction)
	assert.True(t, r.Origin.Equals(origin))
	assert.True(t, r.Direction.Equals(direction))
}

func TestComputingAPointFromADistance(t *testing.T) {
	r := InitRay(tuples.InitPoint(2, 3, 4), tuples.InitVector(1, 0, 0))
	assert.True(t, r.Position(0).Equals(tuples.InitPoint(2, 3, 4)))
	assert.True(t, r.Position(1).Equals(tuples.InitPoint(3, 3, 4)))
	assert.True(t, r.Position(-1).Equals(tuples.InitPoint(1, 3, 4)))
	assert.True(t, r.Position(2.5).Equals(tuples.InitPoint(4.5, 3, 4)))
}

func TestTranslatingARay(t *testing.T) {
	r := InitRay(tuples.InitPoint(1, 2, 3), tuples.InitVector(0, 1, 0))
	m := matrix.Translation(3, 4, 5)
	r2 := r.Transform(m)
	assert.True(t, r2.Origin.Equals(tuples.InitPoint(4, 6, 8)))
	assert.True(t, r2.Direction.Equals(tuples.InitVector(0, 1, 0)))
}

func TestScalingARay(t *testing.T) {
	r := InitRay(tuples.InitPoint(1, 2, 3), tuples.InitVector(0, 1, 0))
	m := matrix.Scaling(2, 3, 4)
	r2 := r.Transform(m)
	assert.True(t, r2.Origin.Equals(tuples.InitPoint(2, 6, 12)))
	assert.True(t, r2.Direction.Equals(tuples.InitVector(0, 3, 0)))
}
