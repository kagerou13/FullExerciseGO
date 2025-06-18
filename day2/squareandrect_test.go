package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type shape interface {
	Round() float64
}

type square struct {
	radius float64
}

func (s square) Round() float64 {
	return s.radius * s.radius
}

type rect struct {
	height float64
	weight float64
}

func (r rect) Round() float64 {
	return 2 * (r.height + r.weight)
}

func Meas(s shape) float64 {
	if c, ok := s.(square); ok {
		return c.Round()
	}
	return 0
}

func TestRectSquare(t *testing.T) {
	tests := []struct {
		name      string
		requested float64
		expect    float64
	}{
		{
			name:      "first testing",
			requested: 10,
			expect:    100,
		},
		{
			name:      "second testing",
			requested: 2,
			expect:    4,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sq1 := square{radius: test.requested}
			result := sq1.Round()
			assert.Equal(t, test.expect, result, "Testing done")
			// assert.True(t, true)
		})
	}

	tests2 := []struct {
		name   string
		req1   float64
		req2   float64
		expect float64
	}{
		{
			name:   "first testing",
			req1:   3,
			req2:   7,
			expect: 20,
		},
		{
			name:   "second testing",
			req1:   13,
			req2:   28,
			expect: 82,
		},
	}

	for _, test2 := range tests2 {
		t.Run(test2.name, func(t *testing.T) {
			rect1 := rect{height: test2.req1, weight: test2.req2}
			result2 := rect1.Round()
			assert.Equal(t, test2.expect, result2, "testing done")
		})
	}
}
