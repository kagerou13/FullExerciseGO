package day4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func countArray(nums []int) (result int) {
	for _, v := range nums {
		result += v
	}
	return
}

func TestSimpleLoop(t *testing.T) {
	tests := []struct {
		name string
		req  []int
		exp  int
	}{
		{
			name: "first second",
			req:  []int{1, 4, 2, 5, 7, 4, 2, 8},
			exp:  33,
		},
		{
			name: "second testing",
			req:  []int{},
			exp:  0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := countArray(test.req)
			assert.Equal(t, test.exp, result, "testing done")
		})
	}
}
