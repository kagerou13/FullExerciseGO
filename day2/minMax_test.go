package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func minMax(nums []int) int {
	max := 0
	for _, v := range nums {
		if max < v {
			max = v
		}
	}
	return max
}

func TestMinMax(t *testing.T) {
	tests := []struct {
		name     string
		req      []int
		expected int
	}{
		{
			name:     "first testing",
			req:      []int{1, 3, 51, 154, 15, 1},
			expected: 154,
		},
		{
			name:     "zero testing",
			req:      []int{},
			expected: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := minMax(test.req)
			assert.Equal(t, test.expected, result, "testing done")
		})
	}
}
