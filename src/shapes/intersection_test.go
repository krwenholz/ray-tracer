package shapes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnIntersectionEncapsulatesTAndObject(t *testing.T) {
	s := InitSphere()
	i := InitIntersection(3.5, *s)
	assert.Equal(t, 3.5, i.T)
	assert.True(t, s.Equals(i.Object))
}

func TestAggregatingIntersections(t *testing.T) {
	s := InitSphere()
	i1 := InitIntersection(1, s)
	i2 := InitIntersection(2, s)
	xs := InitIntersections(i1, i2)
	assert.Equal(t, 2, len(xs.Intersections))
	assert.Equal(t, 1.0, xs.Intersections[0].T)
	assert.Equal(t, 2.0, xs.Intersections[1].T)
}

func TestHitWhenAllIntersectionsHavePositiveT(t *testing.T) {
	s := InitSphere()
	i1 := InitIntersection(1, s)
	i2 := InitIntersection(2, s)
	xs := InitIntersections(i2, i1)
	i := xs.Hit()
	assert.True(t, i1.Equals(i))
}

func TestHitWhenSomeIntersectionsHaveNegativeT(t *testing.T) {
	s := InitSphere()
	i1 := InitIntersection(-1, s)
	i2 := InitIntersection(1, s)
	xs := InitIntersections(i2, i1)
	i := xs.Hit()
	assert.True(t, i2.Equals(i))
}

func TestHitWhenAllIntersectionsHaveNegativeT(t *testing.T) {
	s := InitSphere()
	i1 := InitIntersection(-2, s)
	i2 := InitIntersection(-1, s)
	xs := InitIntersections(i2, i1)
	i := xs.Hit()
	assert.Nil(t, i)
}

func TestHitIsAlwaysTheLowestNonegativeIntersection(t *testing.T) {
	s := InitSphere()
	i1 := InitIntersection(5, s)
	i2 := InitIntersection(7, s)
	i3 := InitIntersection(-3, s)
	i4 := InitIntersection(2, s)
	xs := InitIntersections(i1, i2, i3, i4)
	i := xs.Hit()
	assert.True(t, i4.Equals(i))
}
