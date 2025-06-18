package day2

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func delDuplicateSlice(nums []int) []int {
	slices.Sort(nums)
	var result []int
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			result = append(result, nums[i])
		} else {
			if nums[i] == nums[i-1] {
				continue
			} else {
				result = append(result, nums[i])
			}
		}

	}

	return result
}

func TestDelDuplicateSl(t *testing.T) {
	slices := []struct {
		name     string
		req      []int
		expected []int
	}{
		{
			name:     "first testing",
			req:      []int{1, 12, 5, 1, 14, 13, 4, 5, 4},
			expected: []int{1, 4, 5, 12, 13, 14},
		},
		{
			name:     "zero testing",
			req:      []int{},
			expected: []int(nil),
		},
	}

	for _, slice := range slices {
		t.Run(slice.name, func(t *testing.T) {
			result := delDuplicateSlice(slice.req)
			assert.Equal(t, slice.expected, result, "testing done")
		})
	}
}
