package day3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func checkZeroPosNeg(n int) string {
	if n < 0 {
		return "Negative Number"
	} else if n > 0 {
		return "Positive Number"
	} else {
		return "Zero"
	}

}

func TestCheckZeroPosNeg(t *testing.T) {
	tests := []struct {
		name string
		req  int
		exp  string
	}{
		{
			name: "positive testing",
			req:  14124,
			exp:  "Positive Number",
		},
		{
			name: "zero testing",
			req:  0,
			exp:  "Zero",
		},
		{
			name: "negative testing",
			req:  -14,
			exp:  "Negative Number",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := checkZeroPosNeg(test.req)
			assert.Equal(t, test.exp, result, "testing done")
		})
	}
}
