package day3

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func chBigNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	for _, val := range nums {
		if max < val {
			max = val
		}
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic occur: ", r)
		}
	}()
	return max
}

func TestChBigNumber(t *testing.T) {
	tests := []struct {
		name string
		req  []int
		exp  int
	}{
		{
			name: "first testing",
			req:  []int{1, 5, 4},
			exp:  5,
		},
		{
			name: "zero testing",
			req:  []int{},
			exp:  0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := chBigNumber(test.req)
			assert.Equal(t, test.exp, result, "testing done")
		})
	}
}
