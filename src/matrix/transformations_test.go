package matrix

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"happymonday.dev/ray-tracer/src/tuples"
)

func TestMultiplyingByATranslationMatrix(t *testing.T) {
	transform := Translation(5, -3, 2)
	point := tuples.InitPoint(-3, 4, 5)
	assert.True(t, tuples.InitPoint(2, 1, 7).Equals(transform.MultiplyTuple(point)))
}

func TestMultiplyingByTheInverseOfATranslationMatrix(t *testing.T) {
	transform := Translation(5, -3, 2)
	inverse := transform.Inverse()
	point := tuples.InitPoint(-3, 4, 5)
	assert.True(t, tuples.InitPoint(-8, 7, 3).Equals(inverse.MultiplyTuple(point)))
}

func TestTranslationDoesNotAffectVectors(t *testing.T) {
	transform := Translation(5, -3, 2)
	vector := tuples.InitVector(-3, 4, 5)
	assert.True(t, tuples.InitVector(-3, 4, 5).Equals(transform.MultiplyTuple(vector)))
}

func TestScalingMatrixAppledToAPoint(t *testing.T) {
	transform := Scaling(2, 3, 4)
	point := tuples.InitPoint(-4, 6, 8)
	assert.True(t, tuples.InitPoint(-8, 18, 32).Equals(transform.MultiplyTuple(point)))
}

func TestScalingMatrixAppledToAVector(t *testing.T) {
	transform := Scaling(2, 3, 4)
	vector := tuples.InitVector(-4, 6, 8)
	assert.True(t, tuples.InitVector(-8, 18, 32).Equals(transform.MultiplyTuple(vector)))
}

func TestMultiplyingByTheInverseOfAScalingMatrix(t *testing.T) {
	transform := Scaling(2, 3, 4)
	inv := transform.Inverse()
	vector := tuples.InitVector(-4, 6, 8)
	assert.True(t, tuples.InitVector(-2, 2, 2).Equals(inv.MultiplyTuple(vector)))
}

func TestReflectionIsScalingByANegativeValue(t *testing.T) {
	transform := Scaling(-1, 1, 1)
	point := tuples.InitPoint(2, 3, 4)
	assert.True(t, tuples.InitPoint(-2, 3, 4).Equals(transform.MultiplyTuple(point)))
	assert.True(t, tuples.InitPoint(-2, 3, 4).Equals(ReflectingX(point)))
}

func TestRotatingAPointAroundTheXAxis(t *testing.T) {
	point := tuples.InitPoint(0, 1, 0)
	halfQuarter := RotationX(0.25)
	fullQuarter := RotationX(0.5)
	assert.True(t, tuples.InitPoint(0, math.Sqrt(2)/2, math.Sqrt(2)/2).Equals(halfQuarter.MultiplyTuple(point)))
	assert.True(t, tuples.InitPoint(0, 0, 1).Equals(fullQuarter.MultiplyTuple(point)))
}

func TestInverseOfAnXRotationRotatesInTheOppositeDirection(t *testing.T) {
	point := tuples.InitPoint(0, 1, 0)
	halfQuarter := RotationX(0.25)
	inverse := halfQuarter.Inverse()
	assert.True(t, tuples.InitPoint(0, math.Sqrt(2)/2, -math.Sqrt(2)/2).Equals(inverse.MultiplyTuple(point)))
}

func TestRotatingAPointAroundTheYAxis(t *testing.T) {
	point := tuples.InitPoint(0, 0, 1)
	halfQuarter := RotationY(0.25)
	fullQuarter := RotationY(0.5)
	assert.True(t, tuples.InitPoint(math.Sqrt(2)/2, 0, math.Sqrt(2)/2).Equals(halfQuarter.MultiplyTuple(point)))
	assert.True(t, tuples.InitPoint(1, 0, 0).Equals(fullQuarter.MultiplyTuple(point)))
}

func TestRotatingAPointAroundTheZAxis(t *testing.T) {
	point := tuples.InitPoint(0, 1, 0)
	halfQuarter := RotationZ(0.25)
	fullQuarter := RotationZ(0.5)
	assert.True(t, tuples.InitPoint(-math.Sqrt(2)/2, math.Sqrt(2)/2, 0).Equals(halfQuarter.MultiplyTuple(point)))
	assert.True(t, tuples.InitPoint(-1, 0, 0).Equals(fullQuarter.MultiplyTuple(point)))
}

func TestShearing(t *testing.T) {
	type opt struct {
		name string
		xy   float64
		xz   float64
		yx   float64
		yz   float64
		zx   float64
		zy   float64
		p    *tuples.Tuple
		res  *tuples.Tuple
	}
	opts := []opt{
		{
			"A shearing transformation moves x in proportion to y",
			1,
			0,
			0,
			0,
			0,
			0,
			tuples.InitPoint(2, 3, 4), tuples.InitPoint(5, 3, 4),
		},
		{
			"A shearing transformation moves x in proportion to z",
			0,
			1,
			0,
			0,
			0,
			0,
			tuples.InitPoint(2, 3, 4), tuples.InitPoint(6, 3, 4),
		},
		{
			"A shearing transformation moves y in proportion to x",
			0,
			0,
			1,
			0,
			0,
			0,
			tuples.InitPoint(2, 3, 4), tuples.InitPoint(2, 5, 4),
		},
		{
			"A shearing transformation moves y in proportion to z",
			0,
			0,
			0,
			1,
			0,
			0,
			tuples.InitPoint(2, 3, 4), tuples.InitPoint(2, 7, 4),
		},
		{
			"A shearing transformation moves z in proportion to x",
			0,
			0,
			0,
			0,
			1,
			0,
			tuples.InitPoint(2, 3, 4), tuples.InitPoint(2, 3, 6),
		},
		{
			"A shearing transformation moves z in proportion to y",
			0,
			0,
			0,
			0,
			0,
			1,
			tuples.InitPoint(2, 3, 4), tuples.InitPoint(2, 3, 7),
		},
	}
	for _, o := range opts {
		transform := Shearing(o.xy, o.xz, o.yx, o.yz, o.zx, o.zy)
		assert.True(t, o.res.Equals(transform.MultiplyTuple(o.p)), o.name)
	}
}

func TestIndividualTransformationsAreAppliedInSequence(t *testing.T) {
	p := tuples.InitPoint(1, 0, 1)
	a := RotationX(0.5)
	b := Scaling(5, 5, 5)
	c := Translation(10, 5, 7)
	p2 := a.MultiplyTuple(p)
	assert.True(t, tuples.InitPoint(1, -1, 0).Equals(p2))
	p3 := b.MultiplyTuple(p2)
	assert.True(t, tuples.InitPoint(5, -5, 0).Equals(p3))
	p4 := c.MultiplyTuple(p3)
	assert.True(t, tuples.InitPoint(15, 0, 7).Equals(p4))
}

func TestChainedTransformationsMustBeAppliedInReverseOrder(t *testing.T) {
	p := tuples.InitPoint(1, 0, 1)
	a := RotationX(0.5)
	b := Scaling(5, 5, 5)
	c := Translation(10, 5, 7)
	transform := Chain(a, b, c)
	assert.True(t, tuples.InitPoint(15, 0, 7).Equals(transform.MultiplyTuple(p)))
}
