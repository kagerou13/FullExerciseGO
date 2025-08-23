package day5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func simpleDeleteDuplicate(arr []int) []int {
	c := 0

	for i := 0; i < len(arr); i++ {
		if i == 0 {
			arr[c] = arr[i]
		} else if i == len(arr)-1 && arr[i] == arr[i-1] {
			c++
			arr[c] = arr[i]
		} else if i != len(arr)-1 {
			if arr[i] == arr[i+1] {
				continue
			} else {
				c++
				arr[c] = arr[i]
			}
		}
	}
	return arr[:c+1]
}

func TestSimpleDeleteDuplicate(t *testing.T) {
	tests := []struct {
		name string
		req  []int
		exp  []int
	}{
		{
			name: "first testing",
			req:  []int{1, 2, 5, 2, 6, 17, 3, 5, 4},
			exp:  []int{1, 2, 5, 6, 17},
		},
		{
			name: "zero testing",
			req:  []int{},
			exp:  []int(nil),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := simpleDeleteDuplicate(test.req)
			assert.Equal(t, test.exp, result, "testing done")
		})
	}
}
