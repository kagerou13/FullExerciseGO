package day4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func reverseArray(nums []int) []int {
	var result []int
	for i := len(nums) - 1; i >= 0; i-- {
		result = append(result, nums[i])
	}
	return result
}

func TestReverseArray(t *testing.T) {
	tests := []struct {
		name string
		req  []int
		exp  []int
	}{
		{
			name: "first testing",
			req:  []int{1, 5, 2, 6, 4, 3},
			exp:  []int{3, 4, 6, 2, 5, 1},
		},
		{
			name: "second testing",
			req:  []int{},
			exp:  []int(nil),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := reverseArray(test.req)
			assert.Equal(t, test.exp, result, "testing done")
		})
	}
}
